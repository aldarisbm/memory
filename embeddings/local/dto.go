package local

import (
	"github.com/aldarisbm/memory/embeddings"
)

type DTO struct {
	Host              string
	EmbeddingEndpoint string
}

func (d *DTO) ToEmbedder() embeddings.Embedder {
	return New(
		WithEmbeddingEndpoint(d.EmbeddingEndpoint),
		WithHost(d.Host),
	)
}

var _ embeddings.Converter = (*DTO)(nil)
