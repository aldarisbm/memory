package datasource

// DataSourcer is an interface for data sources
type DataSourcer interface {
	// GetDocument returns the document with the given id
	GetDocument(id string) ([]byte, error)
	// GetDocuments returns the documents with the given ids
	GetDocuments(ids []string) ([][]byte, error)
	// StoreDocument stores the given document
	StoreDocument(document []byte) error
}
