package heisenberg

type DTO struct {
	Dimensions int
	Path       string
	SpaceType  SpaceType
	Collection string
}

func (d *DTO) ToVectorStore() *vectorStorer {
	return New(
		WithDimensions(d.Dimensions),
		WithPath(d.Path),
		WithSpaceType(d.SpaceType),
		WithCollectionName(d.Collection),
	)
}
