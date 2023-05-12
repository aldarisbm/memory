package datasource

import "github.com/aldarisbm/ltm/pkg/shared"

// DataSourcer is an interface for data sources
type DataSourcer interface {
	// GetDocument returns the document with the given id
	GetDocument(id string) (*shared.Document, error)
	// GetDocuments returns the documents with the given ids
	GetDocuments(ids []string) ([]*shared.Document, error)
	// StoreDocument stores the given document
	StoreDocument(document *shared.Document) error
}
