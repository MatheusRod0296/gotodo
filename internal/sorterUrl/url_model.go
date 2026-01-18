package sorterUrl

import "time"

type ShortURL struct {
	ID          int64
	Code        string
	OriginalURL string
	CreatedAt   time.Time
	ExpiresAt   *time.Time
	AccessCount int64
}
