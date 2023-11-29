package cohere

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

// Embedding is cohere API vector embedding.
type Embedding struct {
	Vector []float64 `json:"vector"`
}

// EmbeddingRequest sent to API endpoint.
type EmbeddingRequest struct {
	Texts     []string  `json:"texts"`
	Model     Model     `json:"model,omitempty"`
	InputType InputType `json:"input_type"`
	Truncate  Truncate  `json:"truncate,omitempty"`
}

// EmbedddingResponse received from API endpoint.
type EmbedddingResponse struct {
	Embeddings [][]float64 `json:"embeddings"`
	Meta       *Meta       `json:"meta,omitempty"`
}

// Meta stores API response metadata
type Meta struct {
	APIVersion *APIVersion `json:"api_version,omitempty"`
}

// APIVersion stores metadata API version.
type APIVersion struct {
	Version string `json:"version"`
}

func ToEmbeddings(r io.Reader) ([]*Embedding, error) {
	var resp EmbedddingResponse
	if err := json.NewDecoder(r).Decode(&resp); err != nil {
		return nil, err
	}
	embs := make([]*Embedding, 0, len(resp.Embeddings))
	for _, e := range resp.Embeddings {
		floats := make([]float64, len(e))
		copy(floats, e)
		emb := &Embedding{
			Vector: floats,
		}
		embs = append(embs, emb)
	}
	return embs, nil
}

// Embeddings returns embeddings for every object in EmbeddingRequest.
func (c *Client) Embeddings(ctx context.Context, embReq *EmbeddingRequest) ([]*Embedding, error) {
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

	req, err := c.newRequest(ctx, http.MethodPost, u.String(), body)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ToEmbeddings(resp.Body)
}
