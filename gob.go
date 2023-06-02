package memory

import (
	"encoding/gob"
	"github.com/aldarisbm/memory/datasource/boltdb"
	"github.com/aldarisbm/memory/datasource/sqlite"
	"github.com/aldarisbm/memory/embeddings/local"
	oai "github.com/aldarisbm/memory/embeddings/openai"
	"github.com/aldarisbm/memory/vectorstore/heisenberg"
	pc "github.com/aldarisbm/memory/vectorstore/pinecone"
)

func init() {
	gob.Register(sqlite.DTO{})
	gob.Register(boltdb.DTO{})

	gob.Register(oai.DTO{})
	gob.Register(local.DTO{})

	gob.Register(heisenberg.DTO{})
	gob.Register(pc.DTO{})
}
