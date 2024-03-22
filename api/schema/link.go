package schema

import (
	"time"

	"github.com/google/uuid"
)

type ShortLink struct {
	ID          string    `json:"id"`
	Code        string    `json:"code"`
	OriginalUrl string    `json:"originalUrl"`
	CreatedAt   time.Time `json:"createdAt"`
}

func NewShortLink() *ShortLink {
	shortLink := ShortLink{
		ID:        uuid.New().String(),
		CreatedAt: time.Now(),
	}
	return &shortLink
}