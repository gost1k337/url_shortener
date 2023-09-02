package entity

import "time"

type ShortURL struct {
	UserId      int
	OriginalURL string
	ShortURL    string
	Visits      int
	ExpireAt    time.Duration
	CreatedAt   time.Time
}
