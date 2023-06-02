package oai

type DTO struct {
	ApiKey string
}

func (d *DTO) ToEmbedder() *embedder {
	return NewOpenAIEmbedder(
		WithApiKey(d.ApiKey),
	)
}
