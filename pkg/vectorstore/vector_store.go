package vectorstore

import "github.com/google/uuid"

type VectorStorer interface {
	StoreVector(vector []float32) error
	QueryVector(vector []float32, k int64) ([]uuid.UUID, error)
}
