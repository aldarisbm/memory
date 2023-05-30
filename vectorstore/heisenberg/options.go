package heisenberg

type options struct {
	path       string
	dimensions int
	indexName  string
}

// CallOptions provides a way to set optional parameters to various methods
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

func WithDimensions(dimensions int) CallOptions {
	return CallOptions{
		applyFunc: func(o *options) {
			o.dimensions = dimensions
		},
	}
}
