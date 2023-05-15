package vectorstore

import (
	"github.com/aldarisbm/memory/pkg/shared"
	"github.com/google/uuid"
)

type VectorStorer interface {
	// StoreVector stores the given Document
	StoreVector(document *shared.Document) error
	// QuerySimilarity returns the k most similar documents to the given vector
	QuerySimilarity(vector []float32, k int64) ([]uuid.UUID, error)
}
