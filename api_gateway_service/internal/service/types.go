package service

import (
	"time"
)

type CreateUrlShortResp struct {
	Id          int64     `json:"id,omitempty"`
	OriginalUrl string    `json:"original_url,omitempty"`
	ShortUrl    string    `json:"short_url,omitempty"`
	ExpireAt    time.Time `json:"expire_at,omitempty"`
	Visits      int64     `json:"visits,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}

type GetUrlShortResp struct {
	Id          int64     `json:"id,omitempty"`
	OriginalUrl string    `json:"original_url,omitempty"`
	ShortUrl    string    `json:"short_url,omitempty"`
	ExpireAt    time.Time `json:"expire_at,omitempty"`
	Visits      int64     `json:"visits,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}