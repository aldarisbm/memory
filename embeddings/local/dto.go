package local

import (
	"github.com/aldarisbm/memory/embeddings"
)

type DTO struct {
	Host              string `json:"host"`
	EmbeddingEndpoint string `json:"embedding_endpoint"`
}

func (d *DTO) ToEmbedder() embeddings.Embedder {
	return New(
		WithEmbeddingEndpoint(d.EmbeddingEndpoint),
		WithHost(d.Host),
	)
}

func (d *DTO) GetType() string {
	return "local"
}

var _ embeddings.Converter = (*DTO)(nil)
