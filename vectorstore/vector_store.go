package vectorstore

type VectorStorer interface {
	GetVector(key string) ([]float32, error)
	StoreVector(key string, vector []float32) error
}
