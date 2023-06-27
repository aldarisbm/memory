package oaiembedder

type options struct {
	apiKey string
}

// CallOptions are the options for the Embedder.
type CallOptions struct {
	applyFunc func(o *options)
}

func applyCallOptions(callOptions []CallOptions, defaultOptions ...options) *options {
	o := new(options)
	if len(defaultOptions) > 0 {
		*o = defaultOptions[0]
	}

	for _, callOption := range callOptions {
		callOption.applyFunc(o)
	}

	return o
}

// WithApiKey sets the OpenAI API key for the Embedder.
func WithApiKey(apiKey string) CallOptions {
	return CallOptions{
		applyFunc: func(o *options) {
			o.apiKey = apiKey
		},
	}
}
