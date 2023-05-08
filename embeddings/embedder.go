package embeddings

type Embedder interface {
	EmbedDocument(document []byte) ([]float32, error)
	EmbedDocuments(documents [][]byte) ([][]float32, error)
}
