package constants

import "github.com/shopspring/decimal"

var (
	MinTotalBalance   = decimal.NewFromInt32(0)
	MinTransferAmount = decimal.NewFromInt32(1)
	MinTopUpAmount    = decimal.NewFromInt32(50000)
	MaxTopUpAmount    = decimal.NewFromInt32(10000000)
)
