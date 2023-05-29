package types

import (
	"github.com/google/uuid"
	"time"
)

// Document is a struct that represents a document in the system
// you can make this as long or as short as you want
type Document struct {
	ID         uuid.UUID      `json:"id"`
	GroupingID uuid.UUID      `json:"grouping_id"`
	User       string         `json:"user"`
	Text       string         `json:"text"`
	CreatedAt  time.Time      `json:"created_at"`
	LastReadAt time.Time      `json:"last_read_at"`
	Vector     []float32      `json:"vector"`
	Metadata   map[string]any `json:"metadata"`
}
