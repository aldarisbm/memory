package sqlitevss

type options struct {
	defaults bool
}

type CallOptions struct {
	applyFunc func(*options)
}

func applyCallOptions(opts []CallOptions, defaultOptions ...options) *options {
	o := new(options)
	if len(defaultOptions) > 0 {
		*o = defaultOptions[0]
	}
	for _, opt := range opts {
		opt.applyFunc(o)
	}
	return o
}

func WithDefaults() CallOptions {
	return CallOptions{applyFunc: func(o *options) {
		o.defaults = true
	}}
}
