package handlers

import "time"

type CreateShortURLInput struct {
	OriginalURL string        `json:"original_url"`
	ExpireAt    time.Duration `json:"expire_at"`
}
