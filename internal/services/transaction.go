package services

import (
	"context"
	"fmt"
	"golang-e-wallet-rest-api/internal/apperrors"
	"golang-e-wallet-rest-api/internal/apperrors/errormsg"
	"golang-e-wallet-rest-api/internal/constants"
	"golang-e-wallet-rest-api/internal/dtos"
	"golang-e-wallet-rest-api/internal/repositories"
	"math"

	"github.com/shopspring/decimal"
)

const transactionServ = "transaction service"

type TransactionService interface {
	TransferBalance(ctx context.Context, userId int64, transferReq *dtos.TransferRequest) (*dtos.TransactionResponse, error)
	TopUpBalance(ctx context.Context, userId int64, topUpReq *dtos.TopUpRequest) (*dtos.TransactionResponse, error)
	GetAllTransactions(ctx context.Context, userId int64, page int) ([]dtos.TransactionsRowResponse, error)
}

type transactionService struct {
	transactionRepository repositories.TransactionRepository
	transactor            repositories.Transactor
}

func NewTransactionService(tr repositories.TransactionRepository, t repositories.Transactor) *transactionService {
	return &transactionService{
		transactionRepository: tr,
		transactor:            t,
	}
}

func (ts *transactionService) TransferBalance(ctx context.Context, userId int64, transferReq *dtos.TransferRequest) (*dtos.TransactionResponse, error) {
	res, err := ts.transactor.WithTransaction(ctx, func(ts repositories.TxStore) (any, error) {
		txTransactionRepo := ts.TxTransactionRepository()
		txWalletrepo := ts.TxWalletRepository()

		transaction := dtos.TransferReqToTransactionModel(transferReq)

		if transaction.Amount.LessThan(constants.MinTransferAmount) {
			return nil, apperrors.NewCustomError(nil, errormsg.ErrMsgTransferAmountLessThanMinThreshold, transactionServ)
		}

		walletRecipient, err := txWalletrepo.GetByWalletNumber(ctx, transferReq.ToWalletNumber)
		if err != nil {
			return nil, err
		}

		walletSender, err := txWalletrepo.GetByUserId(ctx, userId)
		if err != nil {
			return nil, err
		}

		if walletRecipient.WalletNumber == walletSender.WalletNumber {
			return nil, apperrors.NewCustomError(nil, errormsg.ErrMsgTransferToSelf, transactionServ)
		}
		if (walletSender.Balance.Sub(transaction.Amount)).LessThan(constants.MinTotalBalance) {
			return nil, apperrors.NewCustomError(nil, errormsg.ErrMsgBalanceInsufficient, transactionServ)
		}

		err = txWalletrepo.DecreaseBalance(ctx, transaction.Amount, walletSender.WalletNumber)
		if err != nil {
			return nil, err
		}
		err = txWalletrepo.IncreaseBalance(ctx, transaction.Amount, walletRecipient.WalletNumber)
		if err != nil {
			return nil, err
		}

		transaction.ToWalletId = walletRecipient.Id
		transaction.FromWalletId = &walletSender.Id
		transaction, err = txTransactionRepo.Capture(ctx, transaction)
		if err != nil {
			return nil, err
		}

		transactionRes := dtos.ModelsToTransactionResponse(transaction, walletRecipient.WalletNumber, &walletSender.WalletNumber)
		return transactionRes, nil
	})
	if err != nil {
		return nil, err
	}

	transactionRes := res.(*dtos.TransactionResponse)
	return transactionRes, nil
}

func (ts *transactionService) TopUpBalance(ctx context.Context, userId int64, topUpReq *dtos.TopUpRequest) (*dtos.TransactionResponse, error) {
	res, err := ts.transactor.WithTransaction(ctx, func(ts repositories.TxStore) (any, error) {
		txTransactionRepo := ts.TxTransactionRepository()
		txWalletrepo := ts.TxWalletRepository()
		txSourceOfFund := ts.TxSourceOfFundRepository()
		txUserRepo := ts.TxUserRepsitory()

		transaction := dtos.TopUpReqToTransactionModel(topUpReq)

		sourceOfFund, err := txSourceOfFund.GetByName(ctx, &topUpReq.SourceOfFund)
		if err != nil {
			return nil, err
		}

		transaction.SourceOfFundId = &sourceOfFund.Id

		if transaction.Amount.LessThan(constants.MinTopUpAmount) {
			return nil, apperrors.NewCustomError(nil, errormsg.ErrMsgTopUpAmountLessThanMinThreshold, transactionServ)
		} else if transaction.Amount.GreaterThan(constants.MaxTopUpAmount) {
			return nil, apperrors.NewCustomError(nil, errormsg.ErrMsgTopUpAmountMoreThanMaxThreshold, transactionServ)
		}

		walletRecipient, err := txWalletrepo.GetByUserId(ctx, userId)
		if err != nil {
			return nil, err
		}

		err = txWalletrepo.IncreaseBalance(ctx, transaction.Amount, walletRecipient.WalletNumber)
		if err != nil {
			return nil, err
		}

		transaction.ToWalletId = walletRecipient.Id
		transaction.Description = fmt.Sprintf("Top Up from %s", sourceOfFund.Name)
		transaction, err = txTransactionRepo.Capture(ctx, transaction)
		if err != nil {
			return nil, err
		}

		attempts := transaction.Amount.Div(decimal.NewFromInt(constants.TopUpThresholdForGameAttempt)).IntPart()

		err = txUserRepo.IncreaseGameAttempt(ctx, attempts, userId)
		if err != nil {
			return nil, err
		}

		transactionRes := dtos.ModelsToTransactionResponse(transaction, walletRecipient.WalletNumber, nil)
		return transactionRes, nil
	})
	if err != nil {
		return nil, err
	}

	transactionRes := res.(*dtos.TransactionResponse)
	return transactionRes, nil
}

func (ts *transactionService) GetAllTransactions(ctx context.Context, userId int64, page int) ([]dtos.TransactionsRowResponse, error) {
	res, err := ts.transactor.WithTransaction(ctx, func(tss repositories.TxStore) (any, error) {
		txTransactionRepo := tss.TxTransactionRepository()
		txWalletRepo := tss.TxWalletRepository()

		offset := (page - 1) * constants.PaginationLimit

		wallet, err := txWalletRepo.GetByUserId(ctx, userId)
		if err != nil {
			return nil, err
		}

		allTransactionsRow, err := txTransactionRepo.GetAll(ctx, wallet.WalletNumber, offset, constants.PaginationLimit)
		if err != nil {
			return nil, err
		}

		var totalRecords int32
		totalRecords = 0
		if len(allTransactionsRow) != 0 {
			totalRecords = allTransactionsRow[0].TotalData
		}
		totalPage := math.Ceil(float64(totalRecords) / float64(10))

		resTransactions := dtos.TransRowModelToResponse(allTransactionsRow, int(totalPage), page)

		return resTransactions, nil
	})
	if err != nil {
		return nil, err
	}

	resTransactions := res.([]dtos.TransactionsRowResponse)
	return resTransactions, nil
}
