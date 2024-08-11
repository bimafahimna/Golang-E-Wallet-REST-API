package controllers_test

import (
	"encoding/json"
	"fmt"
	"golang-e-wallet-rest-api/internal/apperrors"
	"golang-e-wallet-rest-api/internal/apperrors/errormsg"
	"golang-e-wallet-rest-api/internal/controllers"
	"golang-e-wallet-rest-api/internal/dtos"
	"golang-e-wallet-rest-api/internal/mocks"
	"golang-e-wallet-rest-api/internal/pkgs/utils"
	"golang-e-wallet-rest-api/internal/router/middlewares"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func strPointer(text string) *string {
	return &text
}

func TestNewUserController(t *testing.T) {
	t.Run("Should return not nil userController struct pointer", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		mockUserService := mocks.NewUserService(t)

		userController := controllers.NewUserController(mockUserService)

		assert.NotNil(t, userController)
	})
}

func TestUserControllerRegister(t *testing.T) {

	randomString := func(length int) string {
		rand.Seed(time.Now().UnixNano())
		chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		result := make([]byte, length)
		for i := 0; i < length; i++ {
			result[i] = chars[rand.Intn(len(chars))]
		}
		return string(result)
	}

	tests := []struct {
		name        string
		requestBody dtos.UserRegisterRequest
		expectedRes *string
		expectedErr error
	}{
		{
			name: "should return success message 'account registered successfully'",
			requestBody: dtos.UserRegisterRequest{
				Username: "lol",
				Email:    "lmao@gmail.com",
				Password: "321123",
			},
			expectedRes: strPointer("account registered successfully"),
			expectedErr: nil,
		},
		{
			name: "should return error this field is required",
			requestBody: dtos.UserRegisterRequest{
				Username: "lol",
			},
			expectedRes: nil,
			expectedErr: &apperrors.CustomValidationErrors{
				apperrors.ValidationError{
					Field: "Email",
					Msg:   "this field is required",
				},
				apperrors.ValidationError{
					Field: "Password",
					Msg:   "this field is required",
				},
			},
		},
		{
			name: "should return error invalid username not alphanumeric",
			requestBody: dtos.UserRegisterRequest{
				Username: "!@#!@",
				Email:    "lmao@gmail.com",
				Password: "321123",
			},
			expectedRes: nil,
			expectedErr: apperrors.NewCustomError(nil, errormsg.ErrMsgInvalidUsernameNotAlphaNum, ""),
		},
		{
			name: "should return error invalid username, must be equal or less then 254 char",
			requestBody: dtos.UserRegisterRequest{
				Username: randomString(256),
				Email:    "lmao@gmail.com",
				Password: "321123",
			},
			expectedRes: nil,
			expectedErr: apperrors.NewCustomError(nil, errormsg.ErrMsgInvalidUsernameExceedsMaxCharLimit, ""),
		},
		{
			name: "should return error invalid email address format",
			requestBody: dtos.UserRegisterRequest{
				Username: "lol",
				Email:    "lmao",
				Password: "321123",
			},
			expectedRes: nil,
			expectedErr: apperrors.NewCustomError(nil, errormsg.ErrMsgInvalidEmail, ""),
		},
		{
			name: "should return error invalid password not alphanumeric",
			requestBody: dtos.UserRegisterRequest{
				Username: "lol",
				Email:    "lmao@gmail.com",
				Password: "!#@!...  ",
			},
			expectedRes: nil,
			expectedErr: apperrors.NewCustomError(nil, errormsg.ErrMsgInvalidPasswordNotAlphaNum, ""),
		},
		{
			name: "should return error from user service",
			requestBody: dtos.UserRegisterRequest{
				Username: "lol",
				Email:    "lmao@gmail.com",
				Password: "321123",
			},
			expectedRes: nil,
			expectedErr: fmt.Errorf("error from user service"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			reqBody, _ := json.Marshal(test.requestBody)
			ctx.Request = httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(string(reqBody)))

			mockUserService := mocks.NewUserService(t)
			mockUserService.On("RegisterAccount", ctx, &test.requestBody).Return(test.expectedRes, test.expectedErr).Maybe()
			userController := controllers.NewUserController(mockUserService)

			res := utils.ResponseMsgBody(test.expectedErr, test.expectedRes, nil)
			if test.expectedRes == nil {
				res = utils.ResponseMsgBody(test.expectedErr, nil, nil)
			}
			expectedResJSON, _ := json.Marshal(res)

			userController.Register(ctx)
			middlewares.ErrorMiddleware(ctx)

			assert.Equal(t, utils.StatusCode(test.expectedErr), w.Code)
			assert.Equal(t, string(expectedResJSON), w.Body.String())
		})
	}
}

func TestUserControllerLogin(t *testing.T) {
	tests := []struct {
		name        string
		requestBody dtos.UserLoginRequest
		expectedRes gin.H
		expectedErr error
	}{
		{
			name: "should return Json Web Token",
			requestBody: dtos.UserLoginRequest{
				Email:    "lmao@gmail.com",
				Password: "321123",
			},
			expectedRes: gin.H{"token": "ezhawjbdlnubyib98y172bjk"},
			expectedErr: nil,
		},
		{
			name: "should return error this field is required",
			requestBody: dtos.UserLoginRequest{
				Email: "lmao@gmail.com",
			},
			expectedRes: nil,
			expectedErr: &apperrors.CustomValidationErrors{
				apperrors.ValidationError{
					Field: "password",
					Msg:   "this field is required",
				},
			},
		},
		{
			name: "should return error invalid email address format",
			requestBody: dtos.UserLoginRequest{
				Email:    "lmao",
				Password: "321123",
			},
			expectedRes: nil,
			expectedErr: apperrors.NewCustomError(nil, errormsg.ErrMsgInvalidEmail, ""),
		},
		{
			name: "should return error from user service",
			requestBody: dtos.UserLoginRequest{
				Email:    "lmao@gmail.com",
				Password: "321123",
			},
			expectedRes: nil,
			expectedErr: fmt.Errorf("error from user service"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			reqBody, _ := json.Marshal(test.requestBody)
			ctx.Request = httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(string(reqBody)))

			mockUserService := mocks.NewUserService(t)
			mockUserService.On("LoginAccount", ctx, &test.requestBody).Return(test.expectedRes, test.expectedErr).Maybe()
			userController := controllers.NewUserController(mockUserService)

			res := utils.ResponseMsgBody(test.expectedErr, test.expectedRes, nil)
			if test.expectedRes == nil {
				res = utils.ResponseMsgBody(test.expectedErr, nil, nil)
			}
			expectedResJSON, _ := json.Marshal(res)

			userController.Login(ctx)
			middlewares.ErrorMiddleware(ctx)

			assert.Equal(t, utils.StatusCode(test.expectedErr), w.Code)
			assert.Equal(t, string(expectedResJSON), w.Body.String())
		})
	}
}

func TestUserControllerGetDetails(t *testing.T) {
	tests := []struct {
		name        string
		JWTUserID   int64
		expectedRes *dtos.UserDetailsResponse
		expectedErr error
	}{
		{
			name:      "should return user detail response",
			JWTUserID: 1,
			expectedRes: &dtos.UserDetailsResponse{
				Id:           1,
				Username:     "lol",
				Email:        "lmao@gmail.com",
				WalletNumber: 3330000021022,
				Balance:      10000,
				GameAttempts: 0,
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			},
			expectedErr: nil,
		},
		{
			name:        "should return error failed to get user detail",
			JWTUserID:   1,
			expectedRes: nil,
			expectedErr: apperrors.NewCustomError(nil, errormsg.ErrMsgEmailNotExist, ""),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodPost, "/details", nil)

			ctx.Set("user_id", test.JWTUserID)

			mockUserService := mocks.NewUserService(t)
			mockUserService.On("GetDetails", ctx, test.JWTUserID).Return(test.expectedRes, test.expectedErr).Maybe()
			userController := controllers.NewUserController(mockUserService)

			res := utils.ResponseMsgBody(test.expectedErr, test.expectedRes, nil)
			if test.expectedRes == nil {
				res = utils.ResponseMsgBody(test.expectedErr, nil, nil)
			}
			expectedResJSON, _ := json.Marshal(res)

			userController.GetDetails(ctx)
			middlewares.ErrorMiddleware(ctx)

			assert.Equal(t, utils.StatusCode(test.expectedErr), w.Code)
			assert.Equal(t, string(expectedResJSON), w.Body.String())
		})
	}
}

func TestUserControllerForgotPassword(t *testing.T) {
	tests := []struct {
		name        string
		requestBody dtos.UserForgotPwdRequest
		expectedRes *string
		expectedErr error
	}{
		{
			name:        "should return error Email field is required",
			requestBody: dtos.UserForgotPwdRequest{},
			expectedRes: nil,
			expectedErr: &apperrors.CustomValidationErrors{
				apperrors.ValidationError{
					Field: "email",
					Msg:   "this field is required",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			reqBody, _ := json.Marshal(test.requestBody)
			ctx.Request = httptest.NewRequest(http.MethodPost, "/forgot-password", strings.NewReader(string(reqBody)))

			mockUserService := mocks.NewUserService(t)
			mockUserService.On("ForgotPassword", ctx, test.requestBody.Email).Return(test.expectedRes, test.expectedErr).Maybe()
			userController := controllers.NewUserController(mockUserService)

			res := utils.ResponseMsgBody(test.expectedErr, test.expectedRes, nil)
			if test.expectedRes == nil {
				res = utils.ResponseMsgBody(test.expectedErr, nil, nil)
			}
			expectedResJSON, _ := json.Marshal(res)

			userController.ForgotPassword(ctx)
			middlewares.ErrorMiddleware(ctx)

			assert.Equal(t, utils.StatusCode(test.expectedErr), w.Code)
			assert.Equal(t, string(expectedResJSON), w.Body.String())
		})
	}
}
