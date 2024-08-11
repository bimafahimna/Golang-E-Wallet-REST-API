package apperrors

import "fmt"

type CustomError struct {
	Err    error  `json:"-"`
	Msg    string `json:"message,omitempty"`
	Source string `json:"-"`
}

func NewCustomError(err error, msg string, source string) *CustomError {
	if err == nil {
		err = fmt.Errorf("")
	}
	return &CustomError{
		Err:    err,
		Msg:    msg,
		Source: source,
	}
}

func (ce *CustomError) Error() string {
	res := fmt.Sprintf("\n%s,\nMESSAGE: %s,\nSOURCE: %s\n------------------------\n", ce.Err.Error(), ce.Msg, ce.Source)
	return res
}
