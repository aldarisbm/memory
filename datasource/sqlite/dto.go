package sqlite

import (
	"github.com/aldarisbm/memory/datasource"
)

type DTO struct {
	Path string
}

func (d *DTO) ToDataSource() datasource.DataSourcer {
	return NewLocalStorer(
		WithPath(d.Path),
	)
}

var _ datasource.Converter = (*DTO)(nil)
