package pkg

import (
	"fmt"
	"github.com/aldarisbm/ltm/pkg/datasource"
	"github.com/aldarisbm/ltm/pkg/embeddings"
	"github.com/aldarisbm/ltm/pkg/shared"
	"github.com/aldarisbm/ltm/pkg/vectorstore"
)

// LTM is a long-term memory for a chatbot
type LTM struct {
	embedder    embeddings.Embedder
	vectorStore vectorstore.VectorStorer
	datasource  datasource.DataSourcer
}

// NewLTM creates or loads a new LTM instance from the given options
func NewLTM(dataSourcer datasource.DataSourcer, embedder embeddings.Embedder, vectorStorer vectorstore.VectorStorer) *LTM {
	if dataSourcer == nil || embedder == nil || vectorStorer == nil {
		panic("dataSourcer, embedder and vectorStorer must not be nil")
	}

	return &LTM{
		embedder:    embedder,
		vectorStore: vectorStorer,
		datasource:  dataSourcer,
	}
}

// StoreDocument stores a document in the LTM
func (l *LTM) StoreDocument(document *shared.Document) error {
	embedding, err := l.embedder.EmbedDocumentText(document.Text)
	if err != nil {
		return fmt.Errorf("embedding message: %w", err)
	}
	if err := l.vectorStore.StoreVector(embedding); err != nil {
		return fmt.Errorf("calling store vector: %w", err)
	}
	if err := l.datasource.StoreDocument(document); err != nil {
		return fmt.Errorf("storing message: %w", err)
	}
	return nil
}

// RetrieveSimilarDocumentsByText retrieves similar documents from the LTM
func (l *LTM) RetrieveSimilarDocumentsByText(documentText string, topK int64) ([]*shared.Document, error) {
	const TopKDefault int64 = 10
	if topK == 0 {
		topK = TopKDefault
	}
	embedding, err := l.embedder.EmbedDocumentText(documentText)
	if err != nil {
		return nil, fmt.Errorf("embedding message: %w", err)
	}
	ids, err := l.vectorStore.QueryVector(embedding, topK)
	if err != nil {
		return nil, fmt.Errorf("querying vector: %w", err)
	}
	documents, err := l.datasource.GetDocuments(ids)
	if err != nil {
		return nil, fmt.Errorf("getting documents: %w", err)
	}

	return documents, nil
}
