package repositories

import (
	"context"
	"database/sql"
)

type dbtx interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

type Transactor interface {
	WithTransaction(ctx context.Context, fn TxFn) (any, error)
}

type transactor struct {
	db *sql.DB
}

func InitTransactor(db *sql.DB) *transactor {
	return &transactor{
		db: db,
	}
}

type TxFn func(TxStore) (any, error)

func (t *transactor) WithTransaction(ctx context.Context, fn TxFn) (any, error) {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	txStore := InitTxStore(tx)
	result, err := fn(txStore)
	if err != nil {
		return nil, err
	}
	return result, nil
}
