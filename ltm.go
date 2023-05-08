package ltm

import (
	"github.com/aldarisbm/ltm/datasource"
	"github.com/aldarisbm/ltm/embeddings"
	"github.com/aldarisbm/ltm/options"
	"github.com/aldarisbm/ltm/vectorstore"
)

// LTM is a long-term memory for a chatbot
type LTM struct {
	embedder *embeddings.Embedder
	vectorDB *vectorstore.VectorStorer
	localDB  *datasource.DataSourcer
}

// NewLTM creates or loads a new LTM instance from the given options
func NewLTM(opts ...options.CallOptions) *LTM {
	return nil
}
