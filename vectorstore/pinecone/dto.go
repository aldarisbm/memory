package pineconevs

import "github.com/aldarisbm/memory/vectorstore"

type DTO struct {
	ApiKey      string `json:"api_key"`
	IndexName   string `json:"index_name"`
	Namespace   string `json:"namespace"`
	ProjectName string `json:"project_name"`
	Environment string `json:"environment"`
}

func (d *DTO) ToVectorStore() vectorstore.VectorStorer {
	return New(
		WithApiKey(d.ApiKey),
		WithIndexName(d.IndexName),
		WithNamespace(d.Namespace),
		WithProjectName(d.ProjectName),
		WithEnvironment(d.Environment),
	)
}

func (d *DTO) GetType() string {
	return "pinecone"
}

var _ vectorstore.Converter = (*DTO)(nil)
