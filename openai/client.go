package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	// BaseURL is OpenAI HTTP API base URL.
	BaseURL = "https://api.openai.com"
	// EmbedAPIVersion is the latest stable embeddings API version.
	EmbedAPIVersion = "v1"
	// OrgHeader is an Organization header
	OrgHeader = "OpenAI-Organization"
)

// Client is an OpenAI HTTP API client.
type Client struct {
	apiKey  string
	baseURL string
	version string
	orgID   string
	hc      *http.Client
}

// NewClient creates a new HTTP API client and returns it.
// By default it reads the OpenAI API key from OPENAI_API_KEY
// env var and uses the default Go http.Client for making API requests.
// You can override the default options via the client methods.
func NewClient() *Client {
	return &Client{
		apiKey:  os.Getenv("OPENAI_API_KEY"),
		baseURL: BaseURL,
		version: EmbedAPIVersion,
		orgID:   "",
		hc:      &http.Client{},
	}
}

// WithAPIKey sets the API key.
func (c *Client) WithAPIKey(apiKey string) *Client {
	c.apiKey = apiKey
	return c
}

// WithBaseURL sets the API base URL.
func (c *Client) WithBaseURL(baseURL string) *Client {
	c.baseURL = baseURL
	return c
}

// WithVersion sets the API version.
func (c *Client) WithVersion(version string) *Client {
	c.version = version
	return c
}

// WithOrgID sets the organization ID.
func (c *Client) WithOrgID(orgID string) *Client {
	c.orgID = orgID
	return c
}

// WithHTTPClient sets the HTTP client.
func (c *Client) WithHTTPClient(httpClient *http.Client) *Client {
	c.hc = httpClient
	return c
}

// ReqOption is http requestion functional option.
type ReqOption func(*http.Request)

// WithSetHeader sets the header key to value val.
func WithSetHeader(key, val string) ReqOption {
	return func(req *http.Request) {
		if req.Header == nil {
			req.Header = make(http.Header)
		}
		req.Header.Set(key, val)
	}
}

// WithAddHeader adds the val to key header.
func WithAddHeader(key, val string) ReqOption {
	return func(req *http.Request) {
		if req.Header == nil {
			req.Header = make(http.Header)
		}
		req.Header.Add(key, val)
	}
}

func (c *Client) newRequest(ctx context.Context, method, uri string, body io.Reader, opts ...ReqOption) (*http.Request, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	if body == nil {
		body = &bytes.Reader{}
	}

	req, err := http.NewRequestWithContext(ctx, method, uri, body)
	if err != nil {
		return nil, err
	}

	for _, setOption := range opts {
		setOption(req)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	if len(c.orgID) != 0 {
		req.Header.Set("OpenAI-Organization", c.orgID)
	}

	req.Header.Set("Accept", "application/json; charset=utf-8")
	if body != nil {
		// if no content-type is specified we default to json
		if ct := req.Header.Get("Content-Type"); len(ct) == 0 {
			req.Header.Set("Content-Type", "application/json; charset=utf-8")
		}
	}

	return req, nil
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
