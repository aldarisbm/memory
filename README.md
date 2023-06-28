# Disclaimer: This is a work in progress, not ready for production, there will be breaking changes


# Memory is a long term memory implementation for golang

## Description
It is a simple abstraction that allows you to store and retrieve documents based on their text similarity. 
It uses a vector store to store the document embeddings and an embedder to generate the embeddings.
we use document.Document to represent a document, and we can add metadata if we want to.


We can create a new Memory instance by providing a vector store and an embedder. A local store can be 
provided but a default `sqlite` store will be created if not there

It's designed with simple interfaces to be easily extended to other vector stores, embedders and datasources.

### Create a new data source:
### If not path is given for sqlite or boltdb it will use a default path

```go
package main

import "github.com/aldarisbm/memory/datasource/sqliteds"

func main() {
	// can pass options
	ls := sqliteds.NewLocalStorer()
}
```

### Create a new embedder

```go
package main

import (
	oai "github.com/aldarisbm/memory/embeddings/openai"
	"os"
)

func main() {
	// can pass options such as model, but uses default if not provided
	embedder := oai.NewOpenAIEmbedder(
		oai.WithApiKey(os.Getenv("OPENAI_API_KEY")),
	)
}
```

### Create a new vector store

Here we are using pinecone, but there is `heisenberg` that can run locally

```go
package main

import (
    "github.com/aldarisbm/memory/vectorstore/pinecone"
    "os"
)

func main() { 
	// can pass options such as project name, index name, environment and api key 
	// it will create a default namespace if not provided
	vs := pinecone.NewStorer(
        pinecone.WithApiKey(os.Getenv("PINECONE_API_KEY")),
        pinecone.WithIndexName(os.Getenv("PINECONE_INDEX_NAME")),
        pinecone.WithProjectName(os.Getenv("PINECONE_PROJECT_NAME")),
        pinecone.WithEnvironment(os.Getenv("PINECONE_ENVIRONMENT")),
    )
}
```


### Example of using Memory

This example will show how to use the default sqlite store, heisenberg, and possibly the local embedder

```go
package main

import (
	"fmt"
	"github.com/aldarisbm/memory"
	"github.com/aldarisbm/memory/embeddings/local"
	"github.com/aldarisbm/memory/vectorstore/heisenberg"
	"os"
	"time"
)

func main() {
	// for this to work you should be running something like
	// https://github.com/aldarisbm/sentence_transformers locally
	emb := local.New(local.WithHost("http://127.0.0.1:5050"))
	// Uses default SQLite for storage
	// And default Heisenberg for vector store
	mem := memory.NewMemory(memory.WithEmbedder(emb))
	defer mem.Close()

	user := "my-user"
	texts := []string{
		"seinfield is the best comedy show in the world",
		"friends is an ok comedy show in the world",
		"the office is a good comedy show",
		"Wednesday is not really a comedy show",
		"Black Mirror is a suspenseful show",
		"If we are talking about suspenseful shows, then Twilight Zone is the best",
	}

	for _, t := range texts {
		doc := mem.NewDocument(t, user)
		if err := mem.StoreDocument(doc); err != nil {
			panic(err)
		}
	}

	q := "what is a good suspenseful show?"
	docs, err := mem.RetrieveSimilarDocumentsByText(q, 3)
	if err != nil {
		panic(err)
	}
	for _, doc := range docs {
		fmt.Println(doc.Text)
	}
	// Output:
	// the office is a good comedy show
	// Black Mirror is a suspenseful show
	// If we are talking about suspenseful shows, then Twilight Zone is the best

	// the last one being the closest to the query
}
```
