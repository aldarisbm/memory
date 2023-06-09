package vectorstore

import (
	"github.com/aldarisbm/memory/types"
	"github.com/google/uuid"
)

// VectorStorer is an interface for vector stores
type VectorStorer interface {
	// StoreVector stores the given Document
	// We should make sure the length of the vector is the same as the dimension of the vector store
	StoreVector(document *types.Document) error
	// QuerySimilarity returns the k most similar documents to the given vector
	QuerySimilarity(vector []float32, k int64) ([]uuid.UUID, error)
	// Close closes the vector store
	Close() error

	// GetDTO Gets DTO
	GetDTO() Converter
}

type Converter interface {
	ToVectorStore() VectorStorer
	GetType() string
}
