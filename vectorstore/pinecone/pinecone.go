package pinecone

import (
	"github.com/aldarisbm/ltm/options"
	"github.com/nekomeowww/go-pinecone"
)

type PineconeClient struct {
	c         *pinecone.Client
	indexName string
	namespace string
}

func NewPineconeClient(opts ...options.CallOptions) *PineconeClient {
	return nil
}

func (p *PineconeClient) StoreVector(id string, vector []float32) {}

func (p *PineconeClient) GetVector(id string) ([]float32, error) {
	return nil, nil
}

func (p *PineconeClient) QueryVector(vector []float32, k int) ([]float32, error) {
	return nil, nil
}
