package entity

import "time"

type User struct {
	Id           int64
	Username     string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
}
