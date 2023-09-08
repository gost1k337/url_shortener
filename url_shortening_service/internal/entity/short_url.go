package entity

import "time"

type ShortURL struct {
	Id          int64
	OriginalURL string
	ShortURL    string
	Visits      int64
	ExpireAt    time.Time
	CreatedAt   time.Time
}
