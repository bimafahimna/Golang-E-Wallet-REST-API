package utils

import (
	"golang-e-wallet-rest-api/internal/apperrors"
	"golang-e-wallet-rest-api/internal/apperrors/errormsg"
	"net/http"
)

func StatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	if _, ok := err.(*apperrors.CustomValidationErrors); ok {
		return http.StatusBadRequest
	}

	if customErr, ok := err.(*apperrors.CustomError); ok {
		if customErr == nil {
			return http.StatusInternalServerError
		}
		code, ok := errormsg.ErrStatusCodes[customErr.Msg]
		if ok {
			return code
		}
	}

	return http.StatusInternalServerError
}
