package shared

import (
	"github.com/google/uuid"
	"time"
)

type Document struct {
	ID         uuid.UUID `json:"id"`
	Text       string    `json:"text"`
	CreatedAt  time.Time `json:"created_at"`
	LastReadAt time.Time `json:"last_read_at"`
}
