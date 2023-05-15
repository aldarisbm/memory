package datasource

import (
	"github.com/aldarisbm/memory/pkg/document"
	"github.com/google/uuid"
)

// DataSourcer is an interface for data sources
type DataSourcer interface {
	// GetDocument returns the document with the given id
	GetDocument(id uuid.UUID) (*document.Document, error)
	// GetDocuments returns the documents with the given ids
	GetDocuments(ids []uuid.UUID) ([]*document.Document, error)
	// StoreDocument stores the given document
	StoreDocument(document *document.Document) error
	// Close closes the data source
	Close() error
}
