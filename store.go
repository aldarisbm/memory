package memory

import (
	"github.com/aldarisbm/memory/datasource"
	"github.com/aldarisbm/memory/embeddings"
	"github.com/aldarisbm/memory/vectorstore"
)

const (
	BoltDB     = "memoryinternal"
	BucketName = "memories"
)

type DTO struct {
	VS  vectorstore.Converter
	Emb embeddings.Converter
	DS  datasource.Converter
}

type storer interface {
	saveMemoryToStore(name string, mem *Memory) error
	getMemoryFromStore(name string) (*Memory, error)
	close() error
}
