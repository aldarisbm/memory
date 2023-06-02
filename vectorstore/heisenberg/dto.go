package heisenberg

import "github.com/aldarisbm/memory/vectorstore"

type DTO struct {
	Dimensions int
	Path       string
	SpaceType  SpaceType
	Collection string
}

func (d *DTO) ToVectorStore() vectorstore.VectorStorer {
	return New(
		WithDimensions(d.Dimensions),
		WithPath(d.Path),
		WithSpaceType(d.SpaceType),
		WithCollectionName(d.Collection),
	)
}
