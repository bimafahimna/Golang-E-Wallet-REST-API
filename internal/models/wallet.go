package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Wallet struct {
	Id           int64
	UserId       int64
	WalletNumber int64
	Balance      decimal.Decimal
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
