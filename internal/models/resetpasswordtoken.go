package models

import "time"

type ResetPasswordToken struct {
	Id        int64
	UserId    int64
	Token     string
	IsValid   bool
	CreatedAt time.Time
	ExpiredAt time.Time
}
