package shared

import (
	"github.com/google/uuid"
	"time"
)

type Document struct {
	User       string         `json:"user"`
	ID         uuid.UUID      `json:"id"`
	Text       string         `json:"text"`
	CreatedAt  time.Time      `json:"created_at"`
	LastReadAt time.Time      `json:"last_read_at"`
	Vector     []float32      `json:"vector"`
	Metadata   map[string]any `json:"metadata"`
}
