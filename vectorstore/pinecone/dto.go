package pc

import "github.com/aldarisbm/memory/vectorstore"

type DTO struct {
	ApiKey      string
	IndexName   string
	Namespace   string
	ProjectName string
	Environment string
}

func (d *DTO) ToVectorStore() vectorstore.VectorStorer {
	return NewStorer(
		WithApiKey(d.ApiKey),
		WithIndexName(d.IndexName),
		WithNamespace(d.Namespace),
		WithProjectName(d.ProjectName),
		WithEnvironment(d.Environment),
	)
}
