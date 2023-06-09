package heisenbergvs

import (
	"github.com/aldarisbm/memory/internal"
	"github.com/aldarisbm/memory/types"
	"github.com/aldarisbm/memory/vectorstore"
	"github.com/google/uuid"
	"github.com/quantanotes/heisenberg/core"
	"github.com/quantanotes/heisenberg/utils"
)

type heisenberg struct {
	hb          *core.DB
	collection  string
	path        string
	hasBeenInit bool
	DTO         *DTO
}

func New(opts ...CallOptions) *heisenberg {
	o := applyCallOptions(opts, options{
		collection: "asltm",
		spaceType:  Cosine,
	})
	if o.dimensions == 0 {
		panic("dimensions cannot be 0")
	}
	if o.path == "" {
		o.path = internal.CreateFolderInsideMemoryFolder(internal.Generate(10))
	}
	hb := core.NewDB(o.path)
	if !o.hasBeenInit {
		if err := hb.NewCollection(o.collection, uint(o.dimensions), utils.SpaceType(o.spaceType)); err != nil {
			panic(err)
		}
	}

	vs := &heisenberg{
		hb:          hb,
		collection:  o.collection,
		path:        o.path,
		hasBeenInit: true,
		DTO: &DTO{
			Dimensions:  o.dimensions,
			Path:        o.path,
			SpaceType:   o.spaceType,
			Collection:  o.collection,
			HasBeenInit: true,
		},
	}
	return vs
}

func (h *heisenberg) StoreVector(document *types.Document) error {
	id := document.ID.String()
	if err := h.hb.Put(h.collection, id, document.Vector, document.Metadata); err != nil {
		return err
	}
	return nil
}

func (h *heisenberg) QuerySimilarity(vector []float32, k int64) ([]uuid.UUID, error) {
	entries, err := h.hb.Search(h.collection, vector, int(k))
	if err != nil {
		return nil, err
	}
	var uuids []uuid.UUID
	for _, entry := range entries {
		id, err := uuid.Parse(entry.Key)
		if err != nil {
			return nil, err
		}
		uuids = append(uuids, id)
	}
	return uuids, nil
}

func (h *heisenberg) Close() error {
	h.hb.Close()
	return nil
}

func (h *heisenberg) GetDTO() vectorstore.Converter {
	return h.DTO
}

var _ vectorstore.VectorStorer = (*heisenberg)(nil)
