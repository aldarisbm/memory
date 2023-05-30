package heisenberg

import (
	"fmt"
	"github.com/aldarisbm/memory"
	"github.com/aldarisbm/memory/types"
	"github.com/aldarisbm/memory/vectorstore"
	"github.com/google/uuid"
	hs "github.com/quantanotes/heisenberg/core"
	"os"
	"os/user"
)

type Heisenberg struct {
	h *hs.Heisenberg
}

func NewHeisenberg(opts ...CallOptions) *Heisenberg {
	o := applyCallOptions(opts, options{})
	if o.path == "" {
		usr, _ := user.Current()
		dir := usr.HomeDir
		_ = os.Mkdir(fmt.Sprintf("%s/%s", dir, memory.DomainName), os.ModePerm)
		o.path = fmt.Sprintf("%s/%s/boltdb", dir, memory.DomainName)
	}
	return &Heisenberg{
		h: hs.NewHeisenberg(o.path),
	}
}

func (h Heisenberg) StoreVector(document *types.Document) error {
	//TODO implement me
	panic("implement me")
}

func (h Heisenberg) QuerySimilarity(vector []float32, k int64) ([]uuid.UUID, error) {
	//TODO implement me
	panic("implement me")
}

var _ vectorstore.VectorStorer = (*Heisenberg)(nil)
