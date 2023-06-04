package sqlite

import (
	"github.com/aldarisbm/memory/datasource"
)

type DTO struct {
	Path string `json:"path"`
}

func (d *DTO) ToDataSource() datasource.DataSourcer {
	return NewLocalStorer(
		WithPath(d.Path),
	)
}

func (d *DTO) GetType() string {
	return "sqlite"
}

var _ datasource.Converter = (*DTO)(nil)
