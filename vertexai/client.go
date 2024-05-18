package vertexai

import (
	"os"

	"github.com/milosgajdos/go-embeddings"
	"github.com/milosgajdos/go-embeddings/client"
	"golang.org/x/oauth2"
)

const (
	// BaseURL is the Google Vertex AI HTTP API base URL
	BaseURL = "https://us-central1-aiplatform.googleapis.com/v1/projects"
	// ModelURI is the Google Vertex AI HTTP API model URI.
	ModelURI = "locations/us-central1/publishers/google/models"
	// EmbedAction is embedding API action.
	EmbedAction = ":predict"
)

// Client is a Google Vertex AI HTTP API client.
type Client struct {
	opts Options
}

// Options are client options
type Options struct {
	Token      string
	TokenSrc   oauth2.TokenSource
	ProjectID  string
	ModelID    string
	BaseURL    string
	HTTPClient *client.HTTP
}

// Option is functional option.
type Option func(*Options)

// NewClient creates a new HTTP client and returns it.
// By default it reads the following env vars:
// * VERTEXAI_TOKEN for setting the API token
// * VERTEXAI_MODEL_ID for settings the API model ID
// * GOOGLE_PROJECT_ID for setting the Google Project ID
// It uses the default Go http.Client for making API requests
// to the BaseURL. You can override the default client options
// via the client methods.
// NOTE: you must provide either the token or the token source
func NewClient(opts ...Option) *Client {
	options := Options{
		Token:      os.Getenv("VERTEXAI_TOKEN"),
		ModelID:    os.Getenv("VERTEXAI_MODEL_ID"),
		ProjectID:  os.Getenv("GOOGLE_PROJECT_ID"),
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

// WithToken sets the API token.
func WithToken(token string) Option {
	return func(o *Options) {
		o.Token = token
	}
}

// WithTokenSrc sets the API token source.
// The source can be used for generating the API token
// if no token has been set.
func WithTokenSrc(ts oauth2.TokenSource) Option {
	return func(o *Options) {
		o.TokenSrc = ts
	}
}

// WithProjectID sets the Google Project ID.
func WithProjectID(id string) Option {
	return func(o *Options) {
		o.ProjectID = id
	}
}

// WithModelID sets the Vertex AI model ID.
// https://cloud.google.com/vertex-ai/docs/generative-ai/learn/model-versioning
func WithModelID(id string) Option {
	return func(o *Options) {
		o.ModelID = id
	}
}

// WithBaseURL sets the API base URL.
func WithBaseURL(baseURL string) Option {
	return func(o *Options) {
		o.BaseURL = baseURL
	}
}

// WithHTTPClient sets the HTTP client.
func WithHTTPClient(httpClient *client.HTTP) Option {
	return func(o *Options) {
		o.HTTPClient = httpClient
	}
}
