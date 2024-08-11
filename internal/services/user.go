package services

import (
	"context"
	"golang-e-wallet-rest-api/internal/apperrors"
	"golang-e-wallet-rest-api/internal/apperrors/errormsg"
	"golang-e-wallet-rest-api/internal/constants"
	"golang-e-wallet-rest-api/internal/dtos"
	"golang-e-wallet-rest-api/internal/pkgs/utils"
	"golang-e-wallet-rest-api/internal/pkgs/utils/encryption"
	"golang-e-wallet-rest-api/internal/repositories"
	"os"

	"github.com/gin-gonic/gin"
)

const userServ = "user service"

type UserService interface {
	RegisterAccount(ctx context.Context, registUser *dtos.UserRegisterRequest) (*string, error)
	LoginAccount(ctx context.Context, loginUser *dtos.UserLoginRequest) (gin.H, error)
	GetDetails(ctx context.Context, userId int64) (*dtos.UserDetailsResponse, error)
	ForgotPassword(ctx context.Context, email string) (*string, error)
	ResetPassword(ctx context.Context, resetPwdReq *dtos.UserResetPwdRequest) (*string, error)
}

type userService struct {
	userRepository   repositories.UserRepository
	walletRepository repositories.WalletRepsitory
	transactor       repositories.Transactor
}

func NewUserService(ur repositories.UserRepository, wr repositories.WalletRepsitory, t repositories.Transactor) *userService {
	return &userService{
		userRepository:   ur,
		walletRepository: wr,
		transactor:       t,
	}
}

func (us *userService) RegisterAccount(ctx context.Context, registUser *dtos.UserRegisterRequest) (*string, error) {
	res, err := us.transactor.WithTransaction(ctx, func(ts repositories.TxStore) (any, error) {
		txUserRepo := ts.TxUserRepsitory()
		txResetPasswordTokenRepo := ts.TxResetPasswordTokenRepository()
		txWalletRepo := ts.TxWalletRepository()

		encryptedPwd, err := encryption.HashPassword(registUser.Password, constants.SALT)
		if err != nil {
			return nil, apperrors.NewCustomError(err, errormsg.ErrMsgFailedToEncryptPwd, userServ)
		}

		registUser.Password = string(encryptedPwd)
		newUser := dtos.UserRequestToModel(registUser)

		userId, err := txUserRepo.SaveAccount(ctx, newUser)
		if err != nil {
			return nil, err
		}

		err = txResetPasswordTokenRepo.SetupByUserId(ctx, userId)
		if err != nil {
			return nil, err
		}

		err = txWalletRepo.SetupByUserId(ctx, userId)
		if err != nil {
			return nil, err
		}

		resMsg := "account registered successfully"

		return resMsg, nil
	})
	if err != nil {
		return nil, err
	}

	resMsg := res.(string)

	return &resMsg, nil
}

func (us *userService) LoginAccount(ctx context.Context, loginUser *dtos.UserLoginRequest) (gin.H, error) {
	user, err := us.userRepository.GetByEmail(ctx, &loginUser.Email)
	if err != nil {
		return nil, err
	}

	_, err = encryption.CheckPassword(loginUser.Password, user.Password)
	if err != nil {
		return nil, apperrors.NewCustomError(err, errormsg.ErrMsgIncorrectPasswor, userServ)
	}

	jwtProvider := utils.NewJWTProviderHS256(os.Getenv("ISSUER"), os.Getenv("SECRET_KEY"))
	jwt, err := jwtProvider.CreateToken(user.Id)
	if err != nil {
		return nil, err
	}

	return gin.H{"JWT": jwt}, nil
}

func (us *userService) GetDetails(ctx context.Context, userId int64) (*dtos.UserDetailsResponse, error) {
	user, err := us.userRepository.GetById(ctx, userId)
	if err != nil {
		return nil, err
	}

	wallet, err := us.walletRepository.GetByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	userDetailsRes := dtos.ModelToUserDetailsResponse(user, wallet)
	return userDetailsRes, nil
}

func (us *userService) ForgotPassword(ctx context.Context, email string) (*string, error) {
	res, err := us.transactor.WithTransaction(ctx, func(ts repositories.TxStore) (any, error) {
		txUserRepo := ts.TxUserRepsitory()
		txResetPasswordTokenRepo := ts.TxResetPasswordTokenRepository()

		user, err := txUserRepo.GetByEmail(ctx, &email)
		if err != nil {
			return nil, err
		}

		token, err := utils.GenerateOTP()
		if err != nil {
			return nil, apperrors.NewCustomError(err, errormsg.ErrMsgFailedToGenerateOTP, userServ)
		}

		err = txResetPasswordTokenRepo.SaveTokenByUserId(ctx, &token, &user.Id)
		if err != nil {
			return nil, err
		}

		return token, nil
	})
	if err != nil {
		return nil, err
	}

	token := res.(string)
	return &token, nil
}
func (us *userService) ResetPassword(ctx context.Context, resetPwdReq *dtos.UserResetPwdRequest) (*string, error) {
	res, err := us.transactor.WithTransaction(ctx, func(ts repositories.TxStore) (any, error) {
		txUserRepo := ts.TxUserRepsitory()
		txResetPwdTokenRepo := ts.TxResetPasswordTokenRepository()

		user, err := txUserRepo.GetByEmail(ctx, &resetPwdReq.Email)
		if err != nil {
			return nil, err
		}

		resetPwd, err := txResetPwdTokenRepo.GetByUserId(ctx, &user.Id)
		if err != nil {
			return nil, err
		}

		if resetPwdReq.Token != resetPwd.Token {
			return nil, apperrors.NewCustomError(nil, errormsg.ErrMsgInvalidResetPwdToken, userServ)
		}
		if !resetPwd.IsValid {
			return nil, apperrors.NewCustomError(nil, errormsg.ErrMsgResetPwdTokenExpired, userServ)
		}

		err = utils.IsPasswordValid(resetPwdReq.NewPassword)
		if err != nil {
			return nil, err
		}

		hashedPwd, err := encryption.HashPassword(resetPwdReq.NewPassword, constants.SALT)
		if err != nil {
			return nil, err
		}
		err = txUserRepo.SaveNewPassword(ctx, string(hashedPwd), &user.Id)
		if err != nil {
			return nil, err
		}

		err = txResetPwdTokenRepo.ResetToken(ctx, &user.Id)
		if err != nil {
			return nil, err
		}

		resMsg := "password changed successfully"
		return resMsg, nil
	})
	if err != nil {
		return nil, err
	}

	resMsg := res.(string)
	return &resMsg, nil
}
