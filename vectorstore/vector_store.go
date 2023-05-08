package vectorstore

type VectorStorer interface {
	GetVector(id string) ([]float32, error)
	StoreVector(id string, vector []float32) error
	QueryVector(vector []float32, k int) ([]float32, error)
}
