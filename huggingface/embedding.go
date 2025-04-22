package huggingface

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
	Inputs  []string `json:"inputs"`
	Options Options  `json:"options,omitempty"`
}

// Options
type Options struct {
	WaitForModel *bool `json:"wait_for_model,omitempty"`
}

// EmbedddingResponse is returned by API.
// TODO: hugging face APIs are a mess
type EmbedddingResponse [][][][]float64

// ToEmbeddings converts the raw API response,
// parses it into a slice of embeddings and returns it.
func ToEmbeddings(e *EmbedddingResponse) ([]*embeddings.Embedding, error) {
	emb := *e
	embs := make([]*embeddings.Embedding, 0, len(emb))
	// for i := range emb {
	// 	vals := emb[i]
	// 	floats := make([]float64, len(vals))
	// 	copy(floats, vals)
	// 	emb := &embeddings.Embedding{
	// 		Vector: floats,
	// 	}
	// 	embs = append(embs, emb)
	// }
	return embs, nil
}

// Embeddings returns embeddings for every object in EmbeddingRequest.
func (c *Client) Embeddings(ctx context.Context, embReq *EmbeddingRequest) (*EmbedddingResponse, error) {
	u, err := url.Parse(c.baseURL + "/" + c.model)
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
