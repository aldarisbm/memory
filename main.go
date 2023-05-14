package main

import (
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

	emb := openai.NewOpenAIEmbedder(
		openai.WithApiKey(os.Getenv("OPENAI_API_KEY")),
	)

	ls := sqlite.NewLocalStorer()
	ltm := pkg.NewLTM(ls, emb, vs)

}
