package local

type DTO struct {
	Host              string
	EmbeddingEndpoint string
}

func (d *DTO) ToEmbedder() *embedder {
	return New(
		WithEmbeddingEndpoint(d.EmbeddingEndpoint),
		WithHost(d.Host),
	)
}
