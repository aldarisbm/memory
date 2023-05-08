package datasource

type DataSourcer interface {
	GetDocument(key string) ([]byte, error)
	StoreDocument(key string, document []byte) error
}
