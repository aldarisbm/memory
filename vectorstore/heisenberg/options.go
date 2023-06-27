package heisenberg

type SpaceType int

const (
	Ip SpaceType = iota
	Cosine
	L2
)

type options struct {
	path        string
	collection  string
	dimensions  int
	spaceType   SpaceType
	hasBeenInit bool
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

func WithPath(path string) CallOptions {
	return CallOptions{
		applyFunc: func(o *options) {
			o.path = path
		},
	}
}

func WithDimensions(dimensions int) CallOptions {
	return CallOptions{
		applyFunc: func(o *options) {
			o.dimensions = dimensions
		},
	}
}

func WithSpaceType(spaceType SpaceType) CallOptions {
	return CallOptions{
		applyFunc: func(o *options) {
			o.spaceType = spaceType
		},
	}
}

func WithCollectionName(collection string) CallOptions {
	return CallOptions{
		applyFunc: func(o *options) {
			o.collection = collection
		},
	}
}

func WithHasBeenInit(hasBeenInit bool) CallOptions {
	return CallOptions{
		applyFunc: func(o *options) {
			o.hasBeenInit = hasBeenInit
		},
	}
}
