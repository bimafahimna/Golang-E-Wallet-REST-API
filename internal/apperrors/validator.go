package apperrors

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field string `json:"field"`
	Msg   string `json:"message"`
}

func (ve *ValidationError) Error() string {
	return ve.Msg
}

type CustomValidationErrors []ValidationError

func (cve CustomValidationErrors) Error() string {
	buff := bytes.NewBufferString("")

	for i := 0; i < len(cve); i++ {

		buff.WriteString(cve[i].Error())
		buff.WriteString("\n")
	}

	return strings.TrimSpace(buff.String())
}

func NewCustomValidationErrors(ve validator.ValidationErrors) *CustomValidationErrors {
	valErrors := make([]ValidationError, len(ve))

	for i, fe := range ve {
		valErrors[i] = ValidationError{Field: fe.Field(), Msg: fmt.Sprintf("this field is %s", fe.Tag())}
	}

	return (*CustomValidationErrors)(&valErrors)
}
