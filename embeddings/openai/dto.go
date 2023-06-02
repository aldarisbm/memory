package oai

import "github.com/aldarisbm/memory/embeddings"

type DTO struct {
	ApiKey string
}

func (d *DTO) ToEmbedder() embeddings.Embedder {
	return NewOpenAIEmbedder(
		WithApiKey(d.ApiKey),
	)
}
