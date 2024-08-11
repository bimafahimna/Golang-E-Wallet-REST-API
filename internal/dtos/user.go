package dtos

import (
	"golang-e-wallet-rest-api/internal/models"
	"time"
)

type UserRegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserForgotPwdRequest struct {
	Email string `json:"email" binding:"required"`
}

type UserResetPwdRequest struct {
	Email       string `json:"email" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
	Token       string `json:"token" binding:"required"`
}

type UserDetailsResponse struct {
	Id           int64     `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	WalletNumber int64     `json:"wallet_number"`
	Balance      float64   `json:"balance"`
	GameAttempts int32     `json:"game_attempts"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func UserRequestToModel(reqUser *UserRegisterRequest) *models.User {
	return &models.User{
		Username: reqUser.Username,
		Email:    reqUser.Email,
		Password: reqUser.Password,
	}
}

func ModelToUserDetailsResponse(user *models.User, wallet *models.Wallet) *UserDetailsResponse {
	return &UserDetailsResponse{
		Id:           user.Id,
		Username:     user.Username,
		Email:        user.Email,
		WalletNumber: wallet.WalletNumber,
		Balance:      wallet.Balance.InexactFloat64(),
		GameAttempts: user.GameAttempts,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}
