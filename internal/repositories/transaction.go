package repositories

import (
	"context"
	"golang-e-wallet-rest-api/internal/apperrors"
	"golang-e-wallet-rest-api/internal/apperrors/errormsg"
	"golang-e-wallet-rest-api/internal/models"
)

const transactionRepo = "transaction repository"

type TransactionRepository interface {
	Capture(ctx context.Context, transaction *models.Transaction) (*models.Transaction, error)
	GetAll(ctx context.Context, walletNumber int64, offset, limit int) ([]models.TransactionsRow, error)
}

type transactionRepository struct {
	dbtx dbtx
}

func NewTransactionRepository(dbtx dbtx) *transactionRepository {
	return &transactionRepository{
		dbtx: dbtx,
	}
}

func (tr *transactionRepository) Capture(ctx context.Context, transaction *models.Transaction) (*models.Transaction, error) {
	query := `
	INSERT INTO transactions
		(source_of_fund_id, from_wallet_id, to_wallet_id, amount, description)
	VALUES
		($1,$2,$3,$4,$5)
	RETURNING id,created_at
	;
`
	row := tr.dbtx.QueryRowContext(ctx, query,
		transaction.SourceOfFundId,
		transaction.FromWalletId,
		transaction.ToWalletId,
		transaction.Amount,
		transaction.Description,
	)
	if err := row.Err(); err != nil {
		return nil, apperrors.NewCustomError(err, errormsg.ErrMsgInvalidQuery, transactionRepo)
	}

	err := row.Scan(&transaction.Id, &transaction.CreatedAt)
	if err != nil {
		return nil, apperrors.NewCustomError(err, errormsg.ErrMsgFailedToScanData, transactionRepo)
	}

	return transaction, nil
}

func (tr *transactionRepository) GetAll(ctx context.Context, walletNumber int64, offset, limit int) ([]models.TransactionsRow, error) {
	transactions := []models.TransactionsRow{}

	query := `
	SELECT
		t.id,
		s.name as source_of_funds,
		ws.wallet_number as sender_wallet_number,
		wr.wallet_number as recipient_wallet_number,
		t.amount,
		t.description,
		t.created_at,
		count(t.id)
		OVER() as total_rows
	FROM
		transactions t
	LEFT JOIN wallets wr ON wr.id = t.to_wallet_id
	LEFT JOIN wallets ws ON ws.id = t.from_wallet_id
	LEFT JOIN source_of_funds s ON s.id = t.source_of_fund_id
	WHERE
		ws.wallet_number = $1
	OR
		wr.wallet_number = $1
	ORDER BY 
		created_at DESC
	OFFSET $2
	LIMIT $3
	;
	`

	rows, err := tr.dbtx.QueryContext(ctx, query, walletNumber, offset, limit)
	if err != nil {
		return nil, apperrors.NewCustomError(err, errormsg.ErrMsgInvalidQuery, transactionRepo)
	}
	defer rows.Close()

	for rows.Next() {
		var transaction models.TransactionsRow

		err := rows.Scan(
			&transaction.Id,
			&transaction.SourceOfFundName,
			&transaction.FromWalletNumber,
			&transaction.ToWalletNumber,
			&transaction.Amount,
			&transaction.Description,
			&transaction.CreatedAt,
			&transaction.TotalData,
		)
		if err != nil {
			return nil, apperrors.NewCustomError(err, errormsg.ErrMsgFailedToScanData, transactionRepo)
		}

		transactions = append(transactions, transaction)
	}
	err = rows.Err()
	if err != nil {
		return nil, apperrors.NewCustomError(err, errormsg.ErrMsgFailedToGetData, transactionRepo)
	}

	return transactions, nil
}
