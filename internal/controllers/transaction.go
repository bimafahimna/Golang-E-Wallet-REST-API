package controllers

import (
	"golang-e-wallet-rest-api/internal/apperrors"
	"golang-e-wallet-rest-api/internal/apperrors/errormsg"
	"golang-e-wallet-rest-api/internal/constants"
	"golang-e-wallet-rest-api/internal/dtos"
	"golang-e-wallet-rest-api/internal/pkgs/utils"
	"golang-e-wallet-rest-api/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

const transactionControl = "transaction controller"

type TransactionController interface {
	Transfer(ctx *gin.Context)
	TopUp(ctx *gin.Context)
	GetAllTransactions(ctx *gin.Context)
}

type transactionController struct {
	transactionService services.TransactionService
}

func NewTransactionController(ts services.TransactionService) *transactionController {
	return &transactionController{
		transactionService: ts,
	}
}

func (tc *transactionController) Transfer(ctx *gin.Context) {
	userId := ctx.Value("user_id").(int64)

	transferReq := new(dtos.TransferRequest)
	err := ctx.ShouldBindJSON(transferReq)
	if err != nil {
		ctx.Error(err)
		return
	}

	if !transferReq.Amount.Valid {
		ctx.Error(apperrors.NewCustomError(nil, errormsg.ErrMsgEmptyAmount, transactionControl))
		return
	}

	if len(transferReq.Description) > constants.DescriptionCharLimit {
		ctx.Error(apperrors.NewCustomError(nil, errormsg.ErrMsgDescriptionExceedMaxChars, transactionControl))
		return
	}

	resData, err := tc.transactionService.TransferBalance(ctx, userId, transferReq)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(utils.StatusCode(err), utils.ResponseMsgBody(nil, resData, nil))
}

func (tc *transactionController) TopUp(ctx *gin.Context) {
	userId := ctx.Value("user_id").(int64)

	topUpReq := new(dtos.TopUpRequest)
	err := ctx.ShouldBindJSON(topUpReq)
	if err != nil {
		ctx.Error(err)
		return
	}

	if !topUpReq.Amount.Valid {
		ctx.Error(apperrors.NewCustomError(nil, errormsg.ErrMsgEmptyAmount, transactionControl))
		return
	}

	resData, err := tc.transactionService.TopUpBalance(ctx, userId, topUpReq)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(utils.StatusCode(err), utils.ResponseMsgBody(nil, resData, nil))
}

func (tc *transactionController) GetAllTransactions(ctx *gin.Context) {
	userId := ctx.Value("user_id").(int64)

	strPage := ctx.DefaultQuery("page", "1")
	page, err := strconv.Atoi(strPage)
	if err != nil {
		ctx.Error(apperrors.NewCustomError(err, errormsg.ErrMsgValueIsNotInt, transactionControl))
		return
	}

	resData, err := tc.transactionService.GetAllTransactions(ctx, userId, page)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(utils.StatusCode(err), utils.ResponseMsgBody(nil, resData, &resData[0].Pagination))
}
