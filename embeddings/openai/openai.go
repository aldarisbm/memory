package openai

import (
	"github.com/aldarisbm/ltm/options"
	"github.com/sashabaranov/go-openai"
)

type OpenAIEmbedder struct {
	c *openai.Client
}

func NewOpenAIEmbedder(opts ...options.CallOptions) *OpenAIEmbedder {
	return nil
}

func (e *OpenAIEmbedder) EmbedDocument(document []byte) ([]float32, error) {
	return nil, nil
}

func (e *OpenAIEmbedder) EmbedDocuments(documents [][]byte) ([][]float32, error) {
	return nil, nil
}
