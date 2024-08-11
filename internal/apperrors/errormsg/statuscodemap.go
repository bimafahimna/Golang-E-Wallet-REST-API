package errormsg

import "net/http"

var ErrStatusCodes = map[string]int{
	ErrMsgInvalidQuery:                       http.StatusInternalServerError,
	ErrMsgFailedToScanData:                   http.StatusInternalServerError,
	ErrMsgFailedToGetData:                    http.StatusBadRequest,
	ErrMsgFailedToAuthenticate:               http.StatusUnauthorized,
	ErrMsgUserExist:                          http.StatusBadRequest,
	ErrMsgEmailExist:                         http.StatusBadRequest,
	ErrMsgEmailNotExist:                      http.StatusBadRequest,
	ErrMsgIncorrectPasswor:                   http.StatusBadRequest,
	ErrMsgInvalidUsernameNotAlphaNum:         http.StatusBadRequest,
	ErrMsgInvalidEmail:                       http.StatusBadRequest,
	ErrMsgInvalidPasswordNotAlphaNum:         http.StatusBadRequest,
	ErrMsgInvalidPasswordExceedsMaxCharLimit: http.StatusBadRequest,
	ErrMsgInvalidUsernameExceedsMaxCharLimit: http.StatusBadRequest,
	ErrMsgWalletNumberNotExist:               http.StatusBadRequest,
	ErrMsgTransferToSelf:                     http.StatusBadRequest,
	ErrMsgBalanceInsufficient:                http.StatusBadRequest,
	ErrMsgEmptyAmount:                        http.StatusBadRequest,
	ErrMsgTransferAmountLessThanMinThreshold: http.StatusBadRequest,
	ErrMsgTopUpAmountLessThanMinThreshold:    http.StatusBadRequest,
	ErrMsgTopUpAmountMoreThanMaxThreshold:    http.StatusBadRequest,
	ErrMsgValueIsNotInt:                      http.StatusBadRequest,
	ErrMsgFailedToGenerateOTP:                http.StatusInternalServerError,
	ErrMsgInvalidResetPwdToken:               http.StatusBadRequest,
	ErrMsgResetPwdTokenExpired:               http.StatusBadRequest,
	ErrMsgDescriptionExceedMaxChars:          http.StatusBadRequest,
}
