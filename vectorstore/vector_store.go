package vectorstore

type VectorStorer interface {
	StoreVector(vector []float32) error
	QueryVector(vector []float32, k int) ([]string, error)
}
