package repositories

import (
	"context"
	"database/sql"
	"golang-e-wallet-rest-api/internal/apperrors"
	"golang-e-wallet-rest-api/internal/apperrors/errormsg"
	"golang-e-wallet-rest-api/internal/models"

	"github.com/shopspring/decimal"
)

const walletRepo = "wallet repository"

type WalletRepsitory interface {
	SetupByUserId(ctx context.Context, userId *int64) error
	GetByWalletNumber(ctx context.Context, walletNumber int64) (*models.Wallet, error)
	GetByUserId(ctx context.Context, userId int64) (*models.Wallet, error)
	IncreaseBalance(ctx context.Context, amount decimal.Decimal, walletNumber int64) error
	DecreaseBalance(ctx context.Context, amount decimal.Decimal, walletNumber int64) error
}

type walletRepository struct {
	dbtx dbtx
}

func NewWalletRepository(dbtx dbtx) *walletRepository {
	return &walletRepository{
		dbtx: dbtx,
	}
}

func (wr *walletRepository) SetupByUserId(ctx context.Context, userId *int64) error {
	query := `
		INSERT INTO wallets
			(user_id)
		VALUES
			($1)
	`
	row := wr.dbtx.QueryRowContext(ctx, query, userId)
	if err := row.Err(); err != nil {
		return apperrors.NewCustomError(err, errormsg.ErrMsgInvalidQuery, walletRepo)
	}

	return nil
}

func (wr *walletRepository) GetByWalletNumber(ctx context.Context, walletNumber int64) (*models.Wallet, error) {
	wallet := new(models.Wallet)

	query := `
		SELECT
			id,
			user_id,
			balance
		FROM
			wallets
		WHERE
			wallet_number = $1
		;
	`

	row := wr.dbtx.QueryRowContext(ctx, query, walletNumber)
	if err := row.Err(); err != nil {
		if err == sql.ErrNoRows {
			return nil, apperrors.NewCustomError(err, errormsg.ErrMsgWalletNumberNotExist, walletRepo)
		}
		return nil, apperrors.NewCustomError(err, errormsg.ErrMsgInvalidQuery, walletRepo)
	}

	err := row.Scan(
		&wallet.Id,
		&wallet.UserId,
		&wallet.Balance,
	)
	if err != nil {
		return nil, err
	}

	wallet.WalletNumber = walletNumber

	return wallet, nil
}

func (wr *walletRepository) GetByUserId(ctx context.Context, userId int64) (*models.Wallet, error) {
	wallet := new(models.Wallet)

	query := `
		SELECT
			id,
			wallet_number,
			balance
		FROM
			wallets
		WHERE
			user_id = $1
		;
	`

	row := wr.dbtx.QueryRowContext(ctx, query, userId)
	if err := row.Err(); err != nil {
		if err == sql.ErrNoRows {
			return nil, apperrors.NewCustomError(err, errormsg.ErrMsgWalletNumberNotExist, walletRepo)
		}
		return nil, apperrors.NewCustomError(err, errormsg.ErrMsgInvalidQuery, walletRepo)
	}

	err := row.Scan(
		&wallet.Id,
		&wallet.WalletNumber,
		&wallet.Balance,
	)
	if err != nil {
		return nil, err
	}

	wallet.UserId = userId

	return wallet, nil
}
func (wr *walletRepository) IncreaseBalance(ctx context.Context, amount decimal.Decimal, walletNumber int64) error {
	query := `
		UPDATE
			wallets
		SET
			balance = balance + $1,
			updated_at = NOW()
		WHERE
			wallet_number = $2
		;
	`

	newAmount, _ := amount.Float64()

	row := wr.dbtx.QueryRowContext(ctx, query, newAmount, walletNumber)
	if err := row.Err(); err != nil {
		if err == sql.ErrNoRows {
			return apperrors.NewCustomError(err, errormsg.ErrMsgWalletNumberNotExist, walletRepo)
		}
		return apperrors.NewCustomError(err, errormsg.ErrMsgInvalidQuery, walletRepo)
	}

	return nil
}
func (wr *walletRepository) DecreaseBalance(ctx context.Context, amount decimal.Decimal, walletNumber int64) error {
	query := `
		UPDATE
			wallets
		SET
			balance = balance - $1,
			updated_at = NOW()
		WHERE
			wallet_number = $2
		;
	`

	newAmount, _ := amount.Float64()

	row := wr.dbtx.QueryRowContext(ctx, query, newAmount, walletNumber)
	if err := row.Err(); err != nil {
		if err == sql.ErrNoRows {
			return apperrors.NewCustomError(err, errormsg.ErrMsgWalletNumberNotExist, walletRepo)
		}
		return apperrors.NewCustomError(err, errormsg.ErrMsgInvalidQuery, walletRepo)
	}

	return nil
}
