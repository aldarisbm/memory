package oai

import (
	"context"
	"github.com/aldarisbm/memory/pkg/embeddings"
	"github.com/sashabaranov/go-openai"
)

type embedder struct {
	c     *openai.Client
	model openai.EmbeddingModel
}

func NewOpenAIEmbedder(opts ...CallOptions) *embedder {
	o := applyCallOptions(opts, options{
		model: openai.AdaEmbeddingV2,
	})
	c := openai.NewClient(o.apiKey)
	return &embedder{
		c:     c,
		model: o.model,
	}
}

func (e *embedder) EmbedDocumentText(text string) ([]float32, error) {
	ctx := context.Background()
	req := openai.EmbeddingRequest{
		Input: []string{text},
		Model: e.model,
	}
	resp, err := e.c.CreateEmbeddings(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.Data[0].Embedding, nil
}

func (e *embedder) EmbedDocuments(texts []string) ([][]float32, error) {
	ctx := context.Background()
	req := openai.EmbeddingRequest{
		Input: make([]string, len(texts)),
		Model: e.model,
	}
	for i, text := range texts {
		req.Input[i] = text
	}
	resp, err := e.c.CreateEmbeddings(ctx, req)
	if err != nil {
		return nil, err
	}
	embeddings := make([][]float32, len(resp.Data))
	for i, data := range resp.Data {
		embeddings[i] = data.Embedding
	}
	return embeddings, nil
}

// Ensure embedder implements embeddings.Embedder
var _ embeddings.Embedder = (*embedder)(nil)
