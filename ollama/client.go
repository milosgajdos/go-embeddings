package ollama

import (
	"github.com/milosgajdos/go-embeddings"
	"github.com/milosgajdos/go-embeddings/client"
)

const (
	// BaseURL is Ollama HTTP API embeddings base URL.
	BaseURL = "http://localhost:11434/api"
)

// Client is an OpenAI HTTP API client.
type Client struct {
	opts Options
}

type Options struct {
	BaseURL    string
	HTTPClient *client.HTTP
}

// Option is functional graph option.
type Option func(*Options)

// NewClient creates a new Ollama HTTP API client and returns it.
// You can override the default options via the client methods.
func NewClient(opts ...Option) *Client {
	options := Options{
		BaseURL:    BaseURL,
		HTTPClient: client.NewHTTP(),
	}

	for _, apply := range opts {
		apply(&options)
	}

	return &Client{
		opts: options,
	}
}

// NewEmbedder creates a client that implements embeddings.Embedder
func NewEmbedder(opts ...Option) embeddings.Embedder[*EmbeddingRequest] {
	return NewClient(opts...)
}

// WithBaseURL sets the API base URL.
func WithBaseURL(baseURL string) Option {
	return func(o *Options) {
		o.BaseURL = baseURL
	}
}

// WithVersion sets the API version.
// WithHTTPClient sets the HTTP client.
func WithHTTPClient(httpClient *client.HTTP) Option {
	return func(o *Options) {
		o.HTTPClient = httpClient
	}
}
