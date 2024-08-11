package dtos

import (
	"golang-e-wallet-rest-api/internal/models"

	"github.com/shopspring/decimal"
)

type TopUpRequest struct {
	Amount       decimal.NullDecimal `json:"amount" binding:"required"`
	SourceOfFund string              `json:"source_of_funds" binding:"required"`
}

func TopUpReqToTransactionModel(topUpReq *TopUpRequest) *models.Transaction {
	return &models.Transaction{
		Amount: topUpReq.Amount.Decimal,
	}
}
