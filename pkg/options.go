package memory

import (
	datasource2 "github.com/aldarisbm/memory/pkg/datasource"
	"github.com/aldarisbm/memory/pkg/embeddings"
	"github.com/aldarisbm/memory/pkg/vectorstore"
)

type options struct {
	datasource  datasource2.DataSourcer
	embedder    embeddings.Embedder
	vectorStore vectorstore.VectorStorer
}

type CallOptions struct {
	applyFunc func(o *options)
}

func applyCallOptions(callOptions []CallOptions, defaultOptions ...options) *options {
	o := new(options)
	if len(defaultOptions) > 0 {
		*o = defaultOptions[0]
	}

	for _, callOption := range callOptions {
		callOption.applyFunc(o)
	}

	return o
}

func WithDataSource(ds datasource2.DataSourcer) CallOptions {
	return CallOptions{
		applyFunc: func(o *options) {
			o.datasource = ds
		},
	}
}

func WithEmbedder(e embeddings.Embedder) CallOptions {
	return CallOptions{
		applyFunc: func(o *options) {
			o.embedder = e
		},
	}
}

func WithVectorStore(vs vectorstore.VectorStorer) CallOptions {
	return CallOptions{
		applyFunc: func(o *options) {
			o.vectorStore = vs
		},
	}
}
