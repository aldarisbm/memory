package boltdb

type DTO struct {
	Path string
}

func (d *DTO) ToLocalStorer() *localStorer {
	return NewLocalStorer(
		WithPath(d.Path),
	)
}
