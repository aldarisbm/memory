package sqliteds

import (
	"github.com/aldarisbm/memory/datasource"
)

type DTO struct {
	Path string `json:"path"`
}

func (d *DTO) ToDataSource() datasource.DataSourcer {
	return New(
		WithPath(d.Path),
	)
}

func (d *DTO) GetType() string {
	return "sqlite"
}

var _ datasource.Converter = (*DTO)(nil)
