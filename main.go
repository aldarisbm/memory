package main

import (
	"fmt"
	"github.com/aldarisbm/ltm/pkg"
	"github.com/aldarisbm/ltm/pkg/datasource/sqlite"
	oai "github.com/aldarisbm/ltm/pkg/embeddings/openai"
	pc "github.com/aldarisbm/ltm/pkg/vectorstore/pinecone"
	"os"
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
	memory := ltm.NewLTM(ltm.WithVectorStore(vs), ltm.WithEmbedder(emb), ltm.WithDataSource(ls))

	text := "seinfield is the best comedy show in the world"
	user := "my_user"
	doc := memory.NewDocument(text, user)
	if err := memory.StoreDocument(doc); err != nil {
		panic(err)
	}

	q := "what is the best show in the world?"
	docs, err := memory.RetrieveSimilarDocumentsByText(q, 1)
	if err != nil {
		panic(err)
	}
	for _, d := range docs {
		fmt.Printf("doc: %+v\n", d)
	}
}
