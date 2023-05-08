package datasource

import "github.com/aldarisbm/ltm"

// DataSourcer is an interface for data sources
type DataSourcer interface {
	// GetDocument returns the document with the given id
	GetDocument(id string) (*ltm.Document, error)
	// GetDocuments returns the documents with the given ids
	GetDocuments(ids []string) ([]*ltm.Document, error)
	// StoreDocument stores the given document
	StoreDocument(document *ltm.Document) error
}
