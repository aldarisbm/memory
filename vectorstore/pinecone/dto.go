package pc

type DTO struct {
	ApiKey      string
	IndexName   string
	Namespace   string
	ProjectName string
	Environment string
}

func (d *DTO) ToVectorStorer() *storer {
	return NewStorer(
		WithApiKey(d.ApiKey),
		WithIndexName(d.IndexName),
		WithNamespace(d.Namespace),
		WithProjectName(d.ProjectName),
		WithEnvironment(d.Environment),
	)
}
