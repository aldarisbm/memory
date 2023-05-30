package heisenberg

import (
	"fmt"
	"github.com/aldarisbm/memory"
	"github.com/aldarisbm/memory/types"
	"github.com/aldarisbm/memory/vectorstore"
	"github.com/google/uuid"
	"github.com/quantanotes/heisenberg/core"
	"github.com/quantanotes/heisenberg/utils"
	"os"
	"os/user"
)

type vectorStorer struct {
	hb         *core.Heisenberg
	collection string
}

func NewHeisenberg(opts ...CallOptions) *vectorStorer {
	o := applyCallOptions(opts, options{
		collection: "asltm",
		spaceType:  Cosine,
	})
	if o.dimensions == 0 {
		panic("dimensions cannot be 0")
	}
	if o.path == "" {
		usr, _ := user.Current()
		dir := usr.HomeDir
		_ = os.Mkdir(fmt.Sprintf("%s/%s", dir, memory.DomainName), os.ModePerm)
		o.path = fmt.Sprintf("%s/%s/heisenberg", dir, memory.DomainName)
	}
	heisenberg := core.NewHeisenberg(o.path)
	if err := heisenberg.NewCollection(o.collection, o.dimensions, utils.SpaceType(o.spaceType)); err != nil {
		panic(err)
	}

	vs := &vectorStorer{
		hb: core.NewHeisenberg(o.path),
	}
	return vs
}

func (h vectorStorer) StoreVector(document *types.Document) error {
	id := document.ID.String()
	if err := h.hb.Put(h.collection, id, document.Vector, document.Metadata); err != nil {
		return err
	}
	return nil
}

func (h vectorStorer) QuerySimilarity(vector []float32, k int64) ([]uuid.UUID, error) {
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

func (h vectorStorer) Close() error {
	h.hb.Close()
	return nil
}

var _ vectorstore.VectorStorer = (*vectorStorer)(nil)
