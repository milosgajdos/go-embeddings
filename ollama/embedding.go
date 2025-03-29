package ollama

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/milosgajdos/go-embeddings"
	"github.com/milosgajdos/go-embeddings/request"
)

// EmbeddingRequest is serialized and sent to the API server.
type EmbeddingRequest struct {
	Prompt any    `json:"prompt"`
	Model  string `json:"model"`
}

// EmbeddingResponse received from API.
type EmbeddingResponse struct {
	Embedding []float64 `json:"embedding"`
}

// ToEmbeddings converts the API response,
// into a slice of embeddings and returns it.
func (e *EmbeddingResponse) ToEmbeddings() ([]*embeddings.Embedding, error) {
	floats := make([]float64, len(e.Embedding))
	copy(floats, e.Embedding)
	return []*embeddings.Embedding{
		{Vector: floats},
	}, nil
}

// Embed returns embeddings for every object in EmbeddingRequest.
func (c *Client) Embed(ctx context.Context, embReq *EmbeddingRequest) ([]*embeddings.Embedding, error) {
	u, err := url.Parse(c.opts.BaseURL + "/embeddings")
	if err != nil {
		return nil, err
	}

	var body = &bytes.Buffer{}
	enc := json.NewEncoder(body)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(embReq); err != nil {
		return nil, err
	}

	options := []request.Option{}
	req, err := request.NewHTTP(ctx, http.MethodPost, u.String(), body, options...)
	if err != nil {
		return nil, err
	}

	resp, err := request.Do[APIError](c.opts.HTTPClient, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	e := new(EmbeddingResponse)
	if err := json.NewDecoder(resp.Body).Decode(e); err != nil {
		return nil, err
	}

	return e.ToEmbeddings()
}
