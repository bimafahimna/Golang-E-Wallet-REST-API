package utils

import "golang-e-wallet-rest-api/internal/dtos"

type ResponseMessage struct {
	Err        error            `json:"error,omitempty"`
	Data       any              `json:"data,omitempty"`
	Pagination *dtos.Pagination `json:"pagination,omitempty"`
}

func ResponseMsgBody(err error, data any, pagination *dtos.Pagination) ResponseMessage {
	return ResponseMessage{Err: err, Data: data, Pagination: pagination}
}
