package ltm

import (
	"fmt"
	"github.com/aldarisbm/ltm/datasource"
	"github.com/aldarisbm/ltm/embeddings"
	"github.com/aldarisbm/ltm/options"
	"github.com/aldarisbm/ltm/vectorstore"
)

// LTM is a long-term memory for a chatbot
type LTM struct {
	embedder embeddings.Embedder
	vectorDB vectorstore.VectorStorer
	localDB  datasource.DataSourcer
}

// NewLTM creates or loads a new LTM instance from the given options
func NewLTM(opts ...options.CallOptions) *LTM {
	return &LTM{}
}

// StoreMessage stores a message in the LTM
func (l *LTM) StoreMessage(s string) error {
	embedding, err := l.embedder.EmbedDocument([]byte(s))
	if err != nil {
		return fmt.Errorf("embedding message: %w", err)
	}
	if err := l.vectorDB.StoreVector(embedding); err != nil {
		return fmt.Errorf("storing message vector: %w", err)
	}
	if err := l.localDB.StoreDocument([]byte(s)); err != nil {
		return fmt.Errorf("storing message: %w", err)
	}
	return nil
}

// RetrieveSimilarMessages retrieves similar messages from the LTM
func (l *LTM) RetrieveSimilarMessages(s string) ([]string, error) {
	embedding, err := l.embedder.EmbedDocument([]byte(s))
	if err != nil {
		return nil, fmt.Errorf("embedding message: %w", err)
	}
	ids, err := l.vectorDB.QueryVector(embedding, 10)
	if err != nil {
		return nil, fmt.Errorf("querying vector: %w", err)
	}
	messages, err := l.localDB.GetDocuments(ids)
	if err != nil {
		return nil, fmt.Errorf("getting documents: %w", err)
	}

	// here we should convert the documents into somethign standardized
	return messages, nil
}
