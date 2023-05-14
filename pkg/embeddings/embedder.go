package embeddings

type Embedder interface {
	EmbedDocumentText(text string) ([]float32, error)
	EmbedDocuments(texts []string) ([][]float32, error)
}
