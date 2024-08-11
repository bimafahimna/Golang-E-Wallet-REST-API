package errormsg

import (
	"fmt"
	"golang-e-wallet-rest-api/internal/constants"
)

const (
	ErrMsgInvalidQuery                       = "invalid query"
	ErrMsgFailedToScanData                   = "failed to scan data"
	ErrMsgFailedToGetData                    = "failed to get data"
	ErrMsgFailedToAuthenticate               = "not authorized"
	ErrMsgUserExist                          = "user is already exist"
	ErrMsgEmailExist                         = "email is already exist"
	ErrMsgEmailNotExist                      = "email doesn't exist"
	ErrMsgSourceOfFundNotExist               = "source of funds doesn't exist"
	ErrMsgInvalidRequest                     = "invalid request"
	ErrMsgFailedToEncryptPwd                 = "failed to encrypt password"
	ErrMsgIncorrectPasswor                   = "incorrect password"
	ErrMsgInvalidUsernameNotAlphaNum         = "invalid username, must be alphanumeric!"
	ErrMsgInvalidUsernameExceedsMaxCharLimit = "invalid username, must be equal or less than 254 characters!"
	ErrMsgInvalidEmail                       = "invalid email"
	ErrMsgInvalidPasswordNotAlphaNum         = "invalid password, must be alphanumeric!"
	ErrMsgInvalidPasswordExceedsMaxCharLimit = "invalid password, must be equal or less than 50 characters!"
	ErrMsgWalletNumberNotExist               = "wallet number doesn't exist"
	ErrMsgTransferToSelf                     = "can not transfer to the same wallet number"
	ErrMsgBalanceInsufficient                = "balance insufficient"
	ErrMsgEmptyAmount                        = "amount field can not be empty"
	ErrMsgValueIsNotInt                      = "failed to convert, value is not int"
	ErrMsgFailedToGenerateOTP                = "failed to generate OTP"
	ErrMsgInvalidResetPwdToken               = "invalid reset password token"
	ErrMsgResetPwdTokenExpired               = "reset password token has expired"
	ErrMsgDescriptionExceedMaxChars          = "description can not be more than 35 chars"
)

var (
	ErrMsgTransferAmountLessThanMinThreshold = fmt.Sprintf(
		"transfer amount can not be less than %s", constants.MinTransferAmount.StringFixed(4),
	)
	ErrMsgTopUpAmountLessThanMinThreshold = fmt.Sprintf(
		"top up amount can not be less than %s", constants.MinTopUpAmount.StringFixed(4),
	)
	ErrMsgTopUpAmountMoreThanMaxThreshold = fmt.Sprintf(
		"top up amount can not be more than %s", constants.MaxTopUpAmount.StringFixed(4),
	)
)
