package datasource

type DataSourcer interface {
	GetDocument(id string) ([]byte, error)
	StoreDocument(id string, document []byte) error
}
