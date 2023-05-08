package openai

import (
	"github.com/aldarisbm/ltm/options"
	"github.com/sashabaranov/go-openai"
)

type Embedder struct {
	c *openai.Client
}

func NewOpenAIEmbedder(opts ...options.CallOptions) *Embedder {
	return &Embedder{}
}

func (e *Embedder) EmbedDocument(document []byte) ([]float32, error) {
	return nil, nil
}

func (e *Embedder) EmbedDocuments(documents [][]byte) ([][]float32, error) {
	return nil, nil
}
