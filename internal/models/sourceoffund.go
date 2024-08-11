package models

import "time"

type SourceOfFund struct {
	Id        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
