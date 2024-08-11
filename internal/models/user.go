package models

import "time"

type User struct {
	Id           int64
	Username     string
	Email        string
	Password     string
	GameAttempts int32
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
