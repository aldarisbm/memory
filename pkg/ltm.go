package memory

import (
	"fmt"
	"github.com/aldarisbm/memory/pkg/datasource"
	"github.com/aldarisbm/memory/pkg/datasource/sqlite"
	"github.com/aldarisbm/memory/pkg/embeddings"
	"github.com/aldarisbm/memory/pkg/shared"
	"github.com/aldarisbm/memory/pkg/vectorstore"
	"github.com/google/uuid"
	"time"
)

// Memory is a long-term memory for a chatbot
type Memory struct {
	embedder    embeddings.Embedder
	vectorStore vectorstore.VectorStorer
	datasource  datasource.DataSourcer
}

// NewMemory creates or loads a new Memory instance from the given options
func NewMemory(opts ...CallOptions) *Memory {
	o := applyCallOptions(opts, options{
		datasource: sqlite.NewLocalStorer(),
	})

	if o.embedder == nil || o.vectorStore == nil {
		panic("embedder and vector store must be provided")
	}
	return &Memory{
		embedder:    o.embedder,
		vectorStore: o.vectorStore,
		datasource:  o.datasource,
	}
}

// StoreDocument stores a document in the Memory
func (m *Memory) StoreDocument(document *shared.Document) error {
	embedding, err := m.embedder.EmbedDocumentText(document.Text)
	if err != nil {
		return fmt.Errorf("embedding message: %w", err)
	}
	document.Vector = embedding
	if err := m.vectorStore.StoreVector(document); err != nil {
		return fmt.Errorf("calling store vector: %w", err)
	}
	if err := m.datasource.StoreDocument(document); err != nil {
		return fmt.Errorf("storing message: %w", err)
	}
	return nil
}

// RetrieveSimilarDocumentsByText retrieves similar documents from the Memory
func (m *Memory) RetrieveSimilarDocumentsByText(text string, topK int64) ([]*shared.Document, error) {
	const TopKDefault int64 = 10
	if topK == 0 {
		topK = TopKDefault
	}
	embedding, err := m.embedder.EmbedDocumentText(text)
	if err != nil {
		return nil, fmt.Errorf("embedding message: %w", err)
	}
	ids, err := m.vectorStore.QuerySimilarity(embedding, topK)
	if err != nil {
		return nil, fmt.Errorf("querying vector: %w", err)
	}
	documents, err := m.datasource.GetDocuments(ids)
	if err != nil {
		return nil, err
	}

	return documents, nil
}

func (m *Memory) NewDocument(text string, user string) *shared.Document {
	return &shared.Document{
		ID:         uuid.New(),
		Text:       text,
		User:       user,
		CreatedAt:  time.Now(),
		LastReadAt: time.Now(),
	}
}

func (m *Memory) Close() error {
	return m.datasource.Close()
}
