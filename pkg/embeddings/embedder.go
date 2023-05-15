package embeddings

type Embedder interface {
	// EmbedDocumentText returns the embedding of the given text
	EmbedDocumentText(text string) ([]float32, error)
	// EmbedDocuments returns the embeddings of the given texts
	EmbedDocuments(texts []string) ([][]float32, error)
}
