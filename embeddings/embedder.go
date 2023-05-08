package embeddings

import (
	"github.com/aldarisbm/ltm"
)

type Embedder interface {
	EmbedDocument(document *ltm.Document) ([]float32, error)
	EmbedDocuments(documents [][]byte) ([][]float32, error)
}
