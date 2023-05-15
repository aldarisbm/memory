package oai

import "github.com/sashabaranov/go-openai"

type options struct {
	apiKey string
	model  openai.EmbeddingModel
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

func WithModel(model openai.EmbeddingModel) CallOptions {
	return CallOptions{
		applyFunc: func(o *options) {
			o.model = model
		},
	}
}
