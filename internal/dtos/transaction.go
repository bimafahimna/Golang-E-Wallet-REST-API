package dtos

import (
	"golang-e-wallet-rest-api/internal/models"
	"time"

	"github.com/shopspring/decimal"
)

type TransactionResponse struct {
	Id               int64           `json:"transaction_id"`
	SourceOfFundId   *int64          `json:"source_of_fund_id,omitempty"`
	FromWalletNumber *int64          `json:"from_wallet_number,omitempty"`
	ToWalletNumber   int64           `json:"to_wallet_number"`
	Amount           decimal.Decimal `json:"amount"`
	Description      string          `json:"description"`
	CreatedAt        time.Time       `json:"transaction_time"`
}

type TransactionsRowResponse struct {
	Id               int64      `json:"transaction_id"`
	SourceOfFundName *string    `json:"source_of_funds,omitempty"`
	FromWalletNumber *int64     `json:"from_wallet_number,omitempty"`
	ToWalletNumber   int64      `json:"to_wallet_number"`
	Amount           float64    `json:"amount"`
	Description      string     `json:"description"`
	CreatedAt        time.Time  `json:"transaction_time"`
	Pagination       Pagination `json:"-"`
}

func ModelsToTransactionResponse(transaction *models.Transaction, walletRecipientNumber int64, walletSenderNumber *int64) *TransactionResponse {
	return &TransactionResponse{
		Id:               transaction.Id,
		SourceOfFundId:   transaction.SourceOfFundId,
		FromWalletNumber: walletSenderNumber,
		ToWalletNumber:   walletRecipientNumber,
		Amount:           transaction.Amount,
		Description:      transaction.Description,
		CreatedAt:        transaction.CreatedAt,
	}
}

func TransRowModelToResponse(transactions []models.TransactionsRow, totalPage int, currentPage int) []TransactionsRowResponse {
	newTransResponses := []TransactionsRowResponse{}

	for _, transaction := range transactions {
		amount, _ := transaction.Amount.Float64()

		newTransResponse := TransactionsRowResponse{
			Id:               transaction.Id,
			SourceOfFundName: transaction.SourceOfFundName,
			FromWalletNumber: transaction.FromWalletNumber,
			ToWalletNumber:   transaction.ToWalletNumber,
			Amount:           amount,
			Description:      transaction.Description,
			CreatedAt:        transaction.CreatedAt,
			Pagination: Pagination{
				TotalRecords: int(transaction.TotalData),
				CurrentPage:  currentPage,
				TotalPages:   totalPage,
			},
		}

		newTransResponses = append(newTransResponses, newTransResponse)
	}

	return newTransResponses
}
