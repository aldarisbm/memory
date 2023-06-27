package local

type options struct {
	host              string
	embeddingEndpoint string
}

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

// WithHost sets the host for the Embedder.
func WithHost(host string) CallOptions {
	return CallOptions{
		applyFunc: func(o *options) {
			o.host = host
		},
	}
}

func WithEmbeddingEndpoint(endpoint string) CallOptions {
	return CallOptions{
		applyFunc: func(o *options) {
			o.embeddingEndpoint = endpoint
		},
	}
}
