package vectorstore

import (
	"github.com/aldarisbm/ltm/pkg/shared"
	"github.com/google/uuid"
)

type VectorStorer interface {
	StoreVector(document *shared.Document) error
	QueryVector(vector []float32, k int64) ([]uuid.UUID, error)
}
