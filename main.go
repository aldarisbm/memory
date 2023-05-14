package main

import (
	"fmt"
	"github.com/aldarisbm/ltm/pkg"
	"github.com/aldarisbm/ltm/pkg/datasource/sqlite"
	"github.com/aldarisbm/ltm/pkg/embeddings/openai"
	"github.com/aldarisbm/ltm/pkg/vectorstore/pinecone"
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
	ltm := pkg.NewLTM(ls, emb, vs)

	text := "You should always trust the puppers"
	user := "my_user"
	doc := ltm.NewDocument(text, user)
	if err := ltm.StoreDocument(doc); err != nil {
		panic(err)
	}

	q := "who should i trust?"
	docs, err := ltm.RetrieveSimilarDocumentsByText(q, 1)
	if err != nil {
		panic(err)
	}
	for _, d := range docs {
		fmt.Printf("doc: %+v\n", d)
	}
}
