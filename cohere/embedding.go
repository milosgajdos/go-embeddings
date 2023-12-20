package cohere

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/milosgajdos/go-embeddings"
	"github.com/milosgajdos/go-embeddings/request"
)

// EmbeddingRequest sent to API endpoint.
type EmbeddingRequest struct {
	Texts     []string  `json:"texts"`
	Model     Model     `json:"model,omitempty"`
	InputType InputType `json:"input_type"`
	Truncate  Truncate  `json:"truncate,omitempty"`
}

// EmbedddingResponse received from API.
type EmbedddingResponse struct {
	Embeddings [][]float64 `json:"embeddings"`
	Meta       *Meta       `json:"meta,omitempty"`
}

// ToEmbeddings converts the API response,
// into a slice of embeddings and returns it.
func (e *EmbedddingResponse) ToEmbeddings() ([]*embeddings.Embedding, error) {
	embs := make([]*embeddings.Embedding, 0, len(e.Embeddings))
	for _, e := range e.Embeddings {
		floats := make([]float64, len(e))
		copy(floats, e)
		emb := &embeddings.Embedding{
			Vector: floats,
		}
		embs = append(embs, emb)
	}
	return embs, nil
}

// Meta stores API response metadata.
type Meta struct {
	APIVersion *APIVersion `json:"api_version,omitempty"`
}

// APIVersion stores metadata API version.
type APIVersion struct {
	Version string `json:"version"`
}

// Embeddings returns embeddings for every object in EmbeddingRequest.
func (c *Client) Embeddings(ctx context.Context, embReq *EmbeddingRequest) (*EmbedddingResponse, error) {
	u, err := url.Parse(c.baseURL + "/" + c.version + "/embed")
	if err != nil {
		return nil, err
	}

	var body = &bytes.Buffer{}
	enc := json.NewEncoder(body)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(embReq); err != nil {
		return nil, err
	}

	options := []request.Option{
		request.WithBearer(c.apiKey),
	}

	req, err := request.NewHTTP(ctx, http.MethodPost, u.String(), body, options...)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	e := new(EmbedddingResponse)
	if err := json.NewDecoder(resp.Body).Decode(e); err != nil {
		return nil, err
	}

	return e, nil
}
