package memory

import (
	"fmt"
	"github.com/aldarisbm/memory/datasource"
	"github.com/aldarisbm/memory/datasource/sqlite"
	"github.com/aldarisbm/memory/embeddings"
	"github.com/aldarisbm/memory/types"
	"github.com/aldarisbm/memory/vectorstore"
	"github.com/aldarisbm/memory/vectorstore/heisenberg"
	"github.com/google/uuid"
	"time"
)

// CacheSize is the size of the cache for saving recent documents
const CacheSize = 30

// Memory is a long-term memory for a chatbot
type Memory struct {
	embedder    embeddings.Embedder
	vectorStore vectorstore.VectorStorer
	datasource  datasource.DataSourcer
	cache       map[uuid.UUID]*types.Document
}

// NewMemory creates or loads a new Memory instance from the given options
func NewMemory(name string, opts ...CallOptions) *Memory {
	store := getStore()
	o := applyCallOptions(opts, options{
		datasource: sqlite.NewLocalStorer(),
		cacheSize:  CacheSize,
	})

	// check if memory exists in store
	// TODO we might want to check that the vectorstore and embedder are of the same type
	mem, err := store.getMemoryFromStore(name)
	if err == nil {
		return mem
	}

	if o.embedder == nil {
		panic("embedder must be provided")
	}
	if o.vectorStore == nil {
		o.vectorStore = heisenberg.New(
			heisenberg.WithDimensions(o.embedder.GetDimensions()),
			heisenberg.WithSpaceType(heisenberg.Cosine),
		)
	}
	m := &Memory{
		embedder:    o.embedder,
		vectorStore: o.vectorStore,
		datasource:  o.datasource,
		cache:       make(map[uuid.UUID]*types.Document),
	}

	if err = store.saveMemoryToStore(name, m); err != nil {
		panic(err)
	}
	return m
}

// StoreDocument stores a document in the Memory
func (m *Memory) StoreDocument(document *types.Document) error {
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
	m.addToCache(document)
	return nil
}

// RetrieveSimilarDocumentsByText retrieves similar documents from the Memory
// TODO Update LastReadAt for the document when it is retrieved
func (m *Memory) RetrieveSimilarDocumentsByText(text string, topK int64) ([]*types.Document, error) {
	var documents []*types.Document

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

	for _, id := range ids {
		document, ok := m.cache[id]
		if !ok {
			document, err = m.datasource.GetDocument(id)
			if err != nil {
				return nil, fmt.Errorf("getting message: %w", err)
			}
			m.addToCache(document)
		}
		documents = append(documents, document)
	}

	return documents, nil
}

// NewDocument creates a new document
// we should be able to accept CreatedAt and LastReadAt as parameters
// especially if someone already has many conversations that they want to load
// into this memory
func (m *Memory) NewDocument(text string, user string) *types.Document {
	return &types.Document{
		ID:         uuid.New(),
		Text:       text,
		User:       user,
		CreatedAt:  time.Now(),
		LastReadAt: time.Now(),
	}
}

// Close closes the Memory
func (m *Memory) Close() error {
	return m.datasource.Close()
}

// addToCache adds a document to the cache
// and evicts a document at random if the cache is full
// should probably use an LRU cache instead using Document.LastReadAt
func (m *Memory) addToCache(document *types.Document) {
	if len(m.cache) <= CacheSize {
		m.cache[document.ID] = document
		return
	}
	for k := range m.cache {
		delete(m.cache, k)
		break
	}
	m.cache[document.ID] = document
}
