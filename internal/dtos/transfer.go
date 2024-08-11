package dtos

import (
	"golang-e-wallet-rest-api/internal/models"

	"github.com/shopspring/decimal"
)

type TransferRequest struct {
	ToWalletNumber int64               `json:"to" binding:"required"`
	Amount         decimal.NullDecimal `json:"amount" binding:"required"`
	Description    string              `json:"description"`
}

func TransferReqToTransactionModel(transReq *TransferRequest) *models.Transaction {
	return &models.Transaction{
		Amount:      transReq.Amount.Decimal,
		Description: transReq.Description,
	}
}
