package local

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aldarisbm/memory/embeddings"
	"net/http"
)

type request struct {
	Text string `json:"text"`
}

type response struct {
	Embedding []float32 `json:"embedding"`
}

type embedder struct {
	host              string
	embeddingEndpoint string
	client            *http.Client
}

func New(opts ...CallOptions) *embedder {
	o := applyCallOptions(opts, options{
		host:              "http://localhost:5000",
		embeddingEndpoint: "/embeddings",
	})
	return &embedder{
		host:              o.host,
		embeddingEndpoint: o.embeddingEndpoint,
		client:            &http.Client{},
	}
}
func (e embedder) EmbedDocumentText(text string) ([]float32, error) {
	req := request{
		Text: text,
	}
	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	endpoint := fmt.Sprintf(e.host + e.embeddingEndpoint)
	resp, err := e.client.Post(endpoint, "application/json", bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var r response
	err = json.NewDecoder(resp.Body).Decode(&r)

	return r.Embedding, nil
}

func (e embedder) EmbedDocumentTexts(texts []string) ([][]float32, error) {
	// should make this better but for now, to loop over the texts
	// and call EmbedDocumentText
	embs := make([][]float32, len(texts))
	for _, text := range texts {
		emb, err := e.EmbedDocumentText(text)
		if err != nil {
			return nil, err
		}
		embs = append(embs, emb)
	}
	return embs, nil
}

func (e embedder) GetDimensions() uint {
	const SentenceTransformersDimensions = 384
	return SentenceTransformersDimensions
}

var _ embeddings.Embedder = (*embedder)(nil)
