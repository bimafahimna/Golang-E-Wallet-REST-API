package repositories

import "database/sql"

type TxStore interface {
	TxUserRepsitory() UserRepository
	TxResetPasswordTokenRepository() ResetPasswordTokenRepository
	TxWalletRepository() WalletRepsitory
	TxTransactionRepository() TransactionRepository
	TxSourceOfFundRepository() SourceOfFundRepository
}

type txStore struct {
	tx *sql.Tx
}

func InitTxStore(tx *sql.Tx) *txStore {
	return &txStore{
		tx: tx,
	}
}

func (ts *txStore) TxUserRepsitory() UserRepository {
	return &userRepository{dbtx: ts.tx}
}

func (ts *txStore) TxResetPasswordTokenRepository() ResetPasswordTokenRepository {
	return &resetPasswordTokenRepository{dbtx: ts.tx}
}

func (ts *txStore) TxWalletRepository() WalletRepsitory {
	return &walletRepository{dbtx: ts.tx}
}

func (ts *txStore) TxTransactionRepository() TransactionRepository {
	return &transactionRepository{dbtx: ts.tx}
}

func (ts *txStore) TxSourceOfFundRepository() SourceOfFundRepository {
	return &sourceOfFundRepository{dbtx: ts.tx}
}
