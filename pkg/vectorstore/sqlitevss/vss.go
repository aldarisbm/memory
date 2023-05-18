package sqlitevss

import (
	"github.com/aldarisbm/memory/pkg/types"
	"github.com/aldarisbm/memory/pkg/vectorstore"
	"github.com/google/uuid"
)

// SQLiteVSS is a vector store that uses SQLite as the backend
type SQLiteVSS struct{}

// NewSQLiteVSS returns a new SQLiteVSS
func NewSQLiteVSS() *SQLiteVSS {
	return nil
}

// StoreVector stores the given Document
func (vss *SQLiteVSS) StoreVector(document *types.Document) error {
	return nil
}

// QuerySimilarity returns the k most similar documents to the given vector
func (vss *SQLiteVSS) QuerySimilarity(vector []float32, k int64) ([]uuid.UUID, error) {
	return nil, nil
}

// Close closes the SQLiteVSS
func (vss *SQLiteVSS) Close() error {
	return nil
}

// Ensure that SQLiteVSS implements VectorStorer
var _ vectorstore.VectorStorer = (*SQLiteVSS)(nil)
