package main

import (
	"fmt"
	"github.com/aldarisbm/ltm/pkg"
	"github.com/aldarisbm/ltm/pkg/datasource/sqlite"
	"github.com/aldarisbm/ltm/pkg/embeddings/openai"
	"github.com/aldarisbm/ltm/pkg/shared"
	"github.com/aldarisbm/ltm/pkg/vectorstore/pinecone"
	"github.com/google/uuid"
	"os"
	"time"
)

func main() {

	vs := pc.NewStorer(
		pc.WithApiKey(os.Getenv("PINECONE_API_KEY")),
		pc.WithIndexName(os.Getenv("PINECONE_INDEX_NAME")),
		pc.WithProjectName(os.Getenv("PINECONE_PROJECT_NAME")),
		pc.WithEnvironment(os.Getenv("PINECONE_ENVIRONMENT")),
	)

	emb := oai.NewOpenAIEmbedder(
		oai.WithApiKey(os.Getenv("OPENAI_API_KEY")),
	)

	ls := sqlite.NewLocalStorer()
	ltm := pkg.NewLTM(ls, emb, vs)

	id := uuid.New()

	text := "mi mama me mima"
	if err := ltm.StoreDocument(&shared.Document{
		ID:         id,
		Text:       text,
		CreatedAt:  time.Now(),
		LastReadAt: time.Now(),
	}); err != nil {
		panic(err)
	}

	docs, err := ltm.RetrieveSimilarDocumentsByText(text, 1)
	if err != nil {
		panic(err)
	}
	for _, d := range docs {
		fmt.Printf("doc: %+v\n", d)
	}
}
