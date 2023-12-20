package vertexai

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/milosgajdos/go-embeddings"
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
	token     string
	tokenSrc  oauth2.TokenSource
	projectID string
	modelID   string
	baseURL   string
	hc        *http.Client
}

// NewClient creates a new HTTP client and returns it.
// By default it reads the following env vars:
// * VERTEXAI_TOKEN for setting the API token
// * VERTEXAI_MODEL_ID for settings the API model ID
// * GOOGLE_PROJECT_ID for setting the Google Project ID
// It uses the default Go http.Client for making API requests
// to the BaseURL. You can override the default client options
// via the client methods.
// NOTE: you must provide either the token or the token source
func NewClient() *Client {
	return &Client{
		token:     os.Getenv("VERTEXAI_TOKEN"),
		modelID:   os.Getenv("VERTEXAI_MODEL_ID"),
		projectID: os.Getenv("GOOGLE_PROJECT_ID"),
		baseURL:   BaseURL,
		hc:        &http.Client{},
	}
}

// NewEmbedder creates a client that implements embeddings.Embedder
func NewEmbedder() embeddings.Embedder[*EmbeddingRequest] {
	return NewClient()
}

// WithToken sets the API token.
func (c *Client) WithToken(token string) *Client {
	c.token = token
	return c
}

// WithTokenSrc sets the API token source.
// The source can be used for generating the API token
// if no token has been set.
func (c *Client) WithTokenSrc(ts oauth2.TokenSource) *Client {
	c.tokenSrc = ts
	return c
}

// WithProjectID sets the Google Project ID.
func (c *Client) WithProjectID(id string) *Client {
	c.projectID = id
	return c
}

// WithModelID sets the Vertex AI model ID.
// https://cloud.google.com/vertex-ai/docs/generative-ai/learn/model-versioning
func (c *Client) WithModelID(id string) *Client {
	c.modelID = id
	return c
}

// WithBaseURL sets the API base URL.
func (c *Client) WithBaseURL(baseURL string) *Client {
	c.baseURL = baseURL
	return c
}

// WithHTTPClient sets the HTTP client.
func (c *Client) WithHTTPClient(httpClient *http.Client) *Client {
	c.hc = httpClient
	return c
}

func (c *Client) doRequest(req *http.Request) (*http.Response, error) {
	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusBadRequest {
		return resp, nil
	}
	defer resp.Body.Close()

	var apiErr APIError
	if err := json.NewDecoder(resp.Body).Decode(&apiErr); err != nil {
		return nil, err
	}

	return nil, apiErr
}
