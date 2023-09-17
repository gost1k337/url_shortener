package service

import (
	"time"
)

type CreateURLShortResp struct {
	ID          int64     `json:"id,omitempty"`
	OriginalURL string    `json:"originalUrl,omitempty"`
	ShortURL    string    `json:"shortUrl,omitempty"`
	ExpireAt    time.Time `json:"expireAt,omitempty"`
	Visits      int64     `json:"visits,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
}

type GetURLShortResp struct {
	ID          int64     `json:"id,omitempty"`
	OriginalURL string    `json:"originalUrl,omitempty"`
	ShortURL    string    `json:"shortUrl,omitempty"`
	ExpireAt    time.Time `json:"expireAt,omitempty"`
	Visits      int64     `json:"visits,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
}

type CreateUserResp struct {
	ID        int64     `json:"id,omitempty"`
	Username  string    `json:"username,omitempty"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type GetUserResp struct {
	ID        int64     `json:"id,omitempty"`
	Username  string    `json:"username,omitempty"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type DeleteUserResp struct {
	ID        int64     `json:"id,omitempty"`
	Username  string    `json:"username,omitempty"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}
