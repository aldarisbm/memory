package datasource

import (
	"github.com/aldarisbm/ltm/pkg/shared"
	"github.com/google/uuid"
)

// DataSourcer is an interface for data sources
type DataSourcer interface {
	// GetDocument returns the document with the given id
	GetDocument(id uuid.UUID) (*shared.Document, error)
	// GetDocuments returns the documents with the given ids
	GetDocuments(ids []uuid.UUID) ([]*shared.Document, error)
	// StoreDocument stores the given document
	StoreDocument(document *shared.Document) error
	// Close closes the data source
	Close() error
}
