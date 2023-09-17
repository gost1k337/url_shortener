package handlers

import "time"

type CreateURLShortInput struct {
	OriginalURL string        `json:"originalUrl"`
	ExpireAt    time.Duration `json:"expireAt"`
}

type CreateUserInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
