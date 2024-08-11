package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	Id             int64
	SourceOfFundId *int64
	FromWalletId   *int64
	ToWalletId     int64
	Amount         decimal.Decimal
	Description    string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type TransactionsRow struct {
	Id               int64
	SourceOfFundName *string
	FromWalletNumber *int64
	ToWalletNumber   int64
	Amount           decimal.Decimal
	Description      string
	CreatedAt        time.Time
	TotalData        int32
}
