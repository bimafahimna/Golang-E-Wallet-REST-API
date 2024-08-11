package utils

import (
	"golang-e-wallet-rest-api/internal/apperrors"
	"golang-e-wallet-rest-api/internal/apperrors/errormsg"
	"golang-e-wallet-rest-api/internal/constants"
	"net/mail"
	"regexp"
)

const requestValid = "request validator"

func IsEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func IsAlphaNumeric(word string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(word)
}

func IsPasswordValid(pwd string) error {
	if len(pwd) > constants.PasswordCharLimit {
		return apperrors.NewCustomError(nil, errormsg.ErrMsgInvalidPasswordExceedsMaxCharLimit, requestValid)
	}
	valid := IsAlphaNumeric(pwd)
	if !valid {
		return apperrors.NewCustomError(nil, errormsg.ErrMsgInvalidPasswordNotAlphaNum, requestValid)
	}

	return nil
}

func IsUsernameValid(username string) error {
	if len(username) > constants.UsernameCharLimit {
		return apperrors.NewCustomError(nil, errormsg.ErrMsgInvalidUsernameExceedsMaxCharLimit, requestValid)
	}
	valid := IsAlphaNumeric(username)
	if !valid {
		return apperrors.NewCustomError(nil, errormsg.ErrMsgInvalidUsernameNotAlphaNum, requestValid)
	}

	return nil
}
