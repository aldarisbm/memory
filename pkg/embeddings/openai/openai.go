package openai

import (
	"context"
	"github.com/aldarisbm/ltm/pkg/shared"
	goopenai "github.com/sashabaranov/go-openai"
)

type Embedder struct {
	c              *goopenai.Client
	embeddingModel goopenai.EmbeddingModel
}

func NewOpenAIEmbedder(opts ...CallOptions) *Embedder {
	o := applyCallOptions(opts, options{
		embeddingModel: goopenai.AdaEmbeddingV2,
	})
	c := goopenai.NewClient(o.apiKey)
	return &Embedder{
		c: c,
	}
}

func (e *Embedder) EmbedDocument(document shared.Document) ([]float32, error) {
	ctx := context.Background()
	req := goopenai.EmbeddingRequest{
		Input: []string{document.Text},
		Model: e.embeddingModel,
	}
	resp, err := e.c.CreateEmbeddings(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.Data[0].Embedding, nil
}

func (e *Embedder) EmbedDocuments(documents []shared.Document) ([][]float32, error) {
	ctx := context.Background()
	req := goopenai.EmbeddingRequest{
		Input: make([]string, len(documents)),
		Model: e.embeddingModel,
	}
	for i, document := range documents {
		req.Input[i] = document.Text
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
