package pinecone

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	pc "github.com/nekomeowww/go-pinecone"
)

type PineconeStore struct {
	client    *pc.IndexClient
	namespace string
}

func NewPineconeStore(opts ...CallOptions) *PineconeStore {
	o := applyCallOptions(opts, options{
		namespace: "ltmllm",
	})
	c, err := pc.NewIndexClient(
		pc.WithAPIKey(o.apiKey),
		pc.WithIndexName(o.indexName),
		pc.WithEnvironment(o.environment),
		pc.WithProjectName(o.projectName),
	)
	if err != nil {
		panic(err)
	}
	return &PineconeStore{
		client: c,
	}
}

func (p *PineconeStore) StoreVector(id uuid.UUID, vector []float32) error {
	ctx := context.Background()
	req := pc.UpsertVectorsParams{
		Vectors: []*pc.Vector{
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

// The return should be changed to something more generic
func (p *PineconeStore) QueryVector(vector []float32, k int64) ([]uuid.UUID, error) {
	ctx := context.Background()
	req := pc.QueryParams{
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
	return uuids, nil
}
