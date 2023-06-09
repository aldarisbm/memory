package boltds

type options struct {
	path string
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

func WithPath(path string) CallOptions {
	return CallOptions{
		applyFunc: func(o *options) {
			o.path = path
		},
	}
}
