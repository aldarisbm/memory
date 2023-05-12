package pinecone

type options struct {
	apiKey      string
	indexName   string
	namespace   string
	projectName string
	environment string
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

func WithApiKey(apiKey string) CallOptions {
	return CallOptions{
		applyFunc: func(o *options) {
			o.apiKey = apiKey
		},
	}
}

func WithIndexName(indexName string) CallOptions {
	return CallOptions{
		applyFunc: func(o *options) {
			o.indexName = indexName
		},
	}
}

func WithNamespace(namespace string) CallOptions {
	return CallOptions{
		applyFunc: func(o *options) {
			o.namespace = namespace
		},
	}
}

func WithProjectName(projectName string) CallOptions {
	return CallOptions{
		applyFunc: func(o *options) {
			o.projectName = projectName
		},
	}
}

func WithEnvironment(environment string) CallOptions {
	return CallOptions{
		applyFunc: func(o *options) {
			o.environment = environment
		},
	}
}
