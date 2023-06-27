package oai

import "github.com/aldarisbm/memory/embeddings"

type DTO struct {
	ApiKey string `json:"api_key"`
}

func (d *DTO) ToEmbedder() embeddings.Embedder {
	return NewOpenAIEmbedder(
		WithApiKey(d.ApiKey),
	)
}

func (d *DTO) GetType() string {
	return "openai"
}

var _ embeddings.Converter = (*DTO)(nil)
