package controllers

import (
	"golang-e-wallet-rest-api/internal/apperrors"
	"golang-e-wallet-rest-api/internal/apperrors/errormsg"
	"golang-e-wallet-rest-api/internal/dtos"
	"golang-e-wallet-rest-api/internal/pkgs/utils"
	"golang-e-wallet-rest-api/internal/services"

	"github.com/gin-gonic/gin"
)

const userControl = "user controller"

type UserController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	GetDetails(ctx *gin.Context)
	ForgotPassword(ctx *gin.Context)
	ResetPassword(ctx *gin.Context)
}

type userController struct {
	userService services.UserService
}

func NewUserController(us services.UserService) *userController {
	return &userController{
		userService: us,
	}
}

func (uc *userController) Register(ctx *gin.Context) {
	registUser := new(dtos.UserRegisterRequest)
	err := ctx.ShouldBindJSON(registUser)
	if err != nil {
		ctx.Error(err)
		return
	}

	err = utils.IsUsernameValid(registUser.Username)
	if err != nil {
		ctx.Error(err)
		return
	}
	valid := utils.IsEmailValid(registUser.Email)
	if !valid {
		ctx.Error(apperrors.NewCustomError(nil, errormsg.ErrMsgInvalidEmail, userControl))
		return
	}
	err = utils.IsPasswordValid(registUser.Password)
	if err != nil {
		ctx.Error(err)
		return
	}

	resData, err := uc.userService.RegisterAccount(ctx, registUser)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(utils.StatusCode(err), utils.ResponseMsgBody(nil, resData, nil))
}

func (uc *userController) Login(ctx *gin.Context) {
	loginUser := new(dtos.UserLoginRequest)
	err := ctx.ShouldBindJSON(loginUser)
	if err != nil {
		ctx.Error(err)
		return
	}

	validEmail := utils.IsEmailValid(loginUser.Email)
	if !validEmail {
		ctx.Error(apperrors.NewCustomError(nil, errormsg.ErrMsgInvalidEmail, userControl))
		return
	}

	resData, err := uc.userService.LoginAccount(ctx, loginUser)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(utils.StatusCode(err), utils.ResponseMsgBody(nil, resData, nil))
}

func (uc *userController) GetDetails(ctx *gin.Context) {
	userId := ctx.Value("user_id").(int64)

	resData, err := uc.userService.GetDetails(ctx, userId)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(utils.StatusCode(err), utils.ResponseMsgBody(nil, resData, nil))
}

func (uc *userController) ForgotPassword(ctx *gin.Context) {
	forgotPwdReq := new(dtos.UserForgotPwdRequest)
	err := ctx.ShouldBind(forgotPwdReq)
	if err != nil {
		ctx.Error(err)
		return
	}

	validEmail := utils.IsEmailValid(forgotPwdReq.Email)
	if !validEmail {
		ctx.Error(apperrors.NewCustomError(nil, errormsg.ErrMsgInvalidEmail, userControl))
		return
	}

	resData, err := uc.userService.ForgotPassword(ctx, forgotPwdReq.Email)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(utils.StatusCode(err), utils.ResponseMsgBody(nil, resData, nil))
}

func (uc *userController) ResetPassword(ctx *gin.Context) {
	resetPwdReq := new(dtos.UserResetPwdRequest)
	err := ctx.ShouldBind(resetPwdReq)
	if err != nil {
		ctx.Error(err)
		return
	}

	validEmail := utils.IsEmailValid(resetPwdReq.Email)
	if !validEmail {
		ctx.Error(apperrors.NewCustomError(nil, errormsg.ErrMsgInvalidEmail, userControl))
		return
	}

	err = utils.IsPasswordValid(resetPwdReq.NewPassword)
	if err != nil {
		ctx.Error(err)
		return
	}

	resData, err := uc.userService.ResetPassword(ctx, resetPwdReq)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(utils.StatusCode(err), utils.ResponseMsgBody(nil, resData, nil))
}
