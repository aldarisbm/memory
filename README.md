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

import "github.com/aldarisbm/memory/pkg/datasource/sqlite"

func main() {
	// can pass options
	ls := sqlite.NewLocalStorer()
}
```

### Create a new embedder

```go
package main

import (
	oai "github.com/aldarisbm/memory/pkg/embeddings/openai"
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

```go
package main

import (
    "github.com/aldarisbm/memory/pkg/vectorstore/pinecone"
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

```go
package main

import (
	"fmt"
	"github.com/aldarisbm/memory/pkg"
	oai "github.com/aldarisbm/memory/pkg/embeddings/openai"
	pc "github.com/aldarisbm/memory/pkg/vectorstore/pinecone"
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

	// default local store
	// as of right now expects both vector store and embedder
	// in the future I'd like to use a default local embedder
	mem := memory.NewMemory(memory.WithVectorStore(vs), memory.WithEmbedder(emb))
	defer mem.Close()

	text := "seinfield is the best comedy show in the world"
	user := "my_user"

	doc := mem.NewDocument(text, user)
	if err := mem.StoreDocument(doc); err != nil {
		panic(err)
	}

	q := "what is the best show in the world?"
	docs, err := mem.RetrieveSimilarDocumentsByText(q, 1)
	if err != nil {
		panic(err)
	}
	for _, d := range docs {
		fmt.Printf("doc: %+v\n", d)
	}
}
```
