package heisenbergvs

import "github.com/aldarisbm/memory/vectorstore"

type DTO struct {
	Dimensions  int       `json:"dimensions"`
	Path        string    `json:"path"`
	SpaceType   SpaceType `json:"space_type"`
	Collection  string    `json:"collection"`
	HasBeenInit bool      `json:"has_been_init"`
}

func (d *DTO) ToVectorStore() vectorstore.VectorStorer {
	return New(
		WithDimensions(d.Dimensions),
		WithPath(d.Path),
		WithSpaceType(d.SpaceType),
		WithCollectionName(d.Collection),
		WithHasBeenInit(d.HasBeenInit),
	)
}

func (d *DTO) GetType() string {
	return "heisenberg"
}

var _ vectorstore.Converter = (*DTO)(nil)
