package handlers

import "time"

type CreateUrlShortInput struct {
	OriginalUrl string        `json:"original_url"`
	ExpireAt    time.Duration `json:"expire_at"`
}
