package pinecone

type options struct {
	apiKey string
}
type CallOptions struct {
	applyFunc func(o *options)
}

func WithApiKey(apiKey string) CallOptions {
	return CallOptions{
		applyFunc: func(o *options) {
			o.apiKey = apiKey
		},
	}
}
