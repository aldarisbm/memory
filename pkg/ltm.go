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
func NewLTM(embedder embeddings.Embedder, storer vectorstore.VectorStorer, dataSourcer datasource.DataSourcer) *LTM {
	return &LTM{
		embedder:    embedder,
		vectorStore: storer,
		datasource:  dataSourcer,
	}
}

// StoreDocument stores a document in the LTM
func (l *LTM) StoreDocument(document *shared.Document) error {
	embedding, err := l.embedder.EmbedDocument(document)
	if err != nil {
		return fmt.Errorf("embedding message: %w", err)
	}
	if err := l.vectorStore.StoreVector(embedding); err != nil {
		return fmt.Errorf("storing message vector: %w", err)
	}
	if err := l.datasource.StoreDocument(document); err != nil {
		return fmt.Errorf("storing message: %w", err)
	}
	return nil
}

// RetrieveSimilarDocuments retrieves similar documents from the LTM
func (l *LTM) RetrieveSimilarDocuments(document *shared.Document) ([]*shared.Document, error) {
	embedding, err := l.embedder.EmbedDocument(document)
	if err != nil {
		return nil, fmt.Errorf("embedding message: %w", err)
	}
	ids, err := l.vectorStore.QueryVector(embedding, 10)
	if err != nil {
		return nil, fmt.Errorf("querying vector: %w", err)
	}
	documents, err := l.datasource.GetDocuments(ids)
	if err != nil {
		return nil, fmt.Errorf("getting documents: %w", err)
	}

	// here we should convert the documents into something standardized
	return documents, nil
}
