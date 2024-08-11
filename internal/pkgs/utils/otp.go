package utils

import (
	"crypto/rand"
	"golang-e-wallet-rest-api/internal/constants"
)

const otpChars = "1234567890"

func GenerateOTP() (string, error) {

	buffer := make([]byte, constants.OTPLength)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(otpChars)
	for i := 0; i < constants.OTPLength; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}

	return string(buffer), nil
}
