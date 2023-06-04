package heisenberg

import "github.com/aldarisbm/memory/vectorstore"

type DTO struct {
	Dimensions int       `json:"dimensions"`
	Path       string    `json:"path"`
	SpaceType  SpaceType `json:"space_type"`
	Collection string    `json:"collection"`
}

func (d *DTO) ToVectorStore() vectorstore.VectorStorer {
	return New(
		WithDimensions(d.Dimensions),
		WithPath(d.Path),
		WithSpaceType(d.SpaceType),
		WithCollectionName(d.Collection),
	)
}

func (d *DTO) GetType() string {
	return "heisenberg"
}

var _ vectorstore.Converter = (*DTO)(nil)
