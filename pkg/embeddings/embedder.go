package embeddings

import "github.com/aldarisbm/ltm/pkg/shared"

type Embedder interface {
	EmbedDocument(document *shared.Document) ([]float32, error)
	EmbedDocuments(documents []*shared.Document) ([][]float32, error)
}
