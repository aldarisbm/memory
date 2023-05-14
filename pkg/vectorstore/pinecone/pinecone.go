package pc

import (
	"context"
	"fmt"
	"github.com/aldarisbm/ltm/pkg/vectorstore"
	"github.com/google/uuid"
	"github.com/nekomeowww/go-pinecone"
)

type Storer struct {
	client    *pinecone.IndexClient
	namespace string
}

func NewStorer(opts ...CallOptions) *Storer {
	o := applyCallOptions(opts, options{
		namespace: "ltmllm",
	})
	c, err := pinecone.NewIndexClient(
		pinecone.WithAPIKey(o.apiKey),
		pinecone.WithIndexName(o.indexName),
		pinecone.WithEnvironment(o.environment),
		pinecone.WithProjectName(o.projectName),
	)
	if err != nil {
		panic(err)
	}
	return &Storer{
		client:    c,
		namespace: o.namespace,
	}
}

func (p *Storer) StoreVector(vector []float32) error {
	id := uuid.New()
	ctx := context.Background()
	req := pinecone.UpsertVectorsParams{
		Vectors: []*pinecone.Vector{
			{
				ID:     id.String(),
				Values: vector,
			},
		},
		Namespace: p.namespace,
	}

	resp, err := p.client.UpsertVectors(ctx, req)
	if err != nil {
		return fmt.Errorf("storing vector: %w", err)
	}
	if resp.UpsertedCount != 1 {
		return fmt.Errorf("storing vector: upserted count is not 1")
	}
	return nil
}

func (p *Storer) QueryVector(vector []float32, k int64) ([]uuid.UUID, error) {
	ctx := context.Background()
	req := pinecone.QueryParams{
		Vector:    vector,
		Namespace: p.namespace,
		TopK:      k,
	}
	resp, err := p.client.Query(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("querying vector: %w", err)
	}
	if len(resp.Matches) == 0 {
		// should we return an error here?
		return nil, nil
	}

	var uuids []uuid.UUID
	for _, match := range resp.Matches {
		id, err := uuid.Parse(match.ID)
		if err != nil {
			return nil, fmt.Errorf("querying vector: %w", err)
		}
		uuids = append(uuids, id)
	}
	// TODO The return should be changed to something more generic
	return uuids, nil
}

var _ vectorstore.VectorStorer = (*Storer)(nil)