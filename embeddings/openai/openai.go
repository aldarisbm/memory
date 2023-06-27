package oai

import (
	"context"
	"github.com/aldarisbm/memory/embeddings"
	"github.com/sashabaranov/go-openai"
)

type embedder struct {
	c     *openai.Client
	model openai.EmbeddingModel
	DTO   *DTO
}

// NewOpenAIEmbedder returns an Embedder that uses OpenAI's API to embed text.
// it always uses the AdaEmbeddingV2 model per OpenAI recommendation.
func NewOpenAIEmbedder(opts ...CallOptions) *embedder {
	o := applyCallOptions(opts)
	c := openai.NewClient(o.apiKey)
	return &embedder{
		c:     c,
		model: openai.AdaEmbeddingV2,
	}
}

// EmbedDocumentText embeds the given text
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

// EmbedDocumentTexts embeds the given texts
func (e *embedder) EmbedDocumentTexts(texts []string) ([][]float32, error) {
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
	embs := make([][]float32, len(resp.Data))
	for i, data := range resp.Data {
		embs[i] = data.Embedding
	}
	return embs, nil
}

func (e *embedder) GetDimensions() int {
	const AdaEmbeddingV2Dimensions = 1536
	return AdaEmbeddingV2Dimensions
}

func (e *embedder) GetDTO() embeddings.Converter {
	return e.DTO
}

// Ensure embedder implements embeddings.Embedder
var _ embeddings.Embedder = (*embedder)(nil)
