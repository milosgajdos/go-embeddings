package voyage

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/milosgajdos/go-embeddings"
	"github.com/milosgajdos/go-embeddings/request"
)

// EmbeddingRequest sent to API endpoint.
type EmbeddingRequest struct {
	Input          []string       `json:"input"`
	Model          Model          `json:"model"`
	InputType      InputType      `json:"input_type,omitempty"`
	EncodingFormat EncodingFormat `json:"encoding_format,omitempty"`
	Truncation     bool           `json:"truncation,omitempty"`
}

// Data stores vector embeddings.
type Data struct {
	Object    string    `json:"object"`
	Index     int       `json:"index"`
	Embedding []float64 `json:"embedding"`
}

// EmbeddingResponseGen is the API response.
type EmbeddingResponse struct {
	Object string `json:"object"`
	Data   []Data `json:"data"`
	Model  Model  `json:"model"`
	Usage  Usage  `json:"usage"`
}

// ToEmbeddings converts the API response,
// into a slice of embeddings and returns it.
func (e *EmbeddingResponse) ToEmbeddings() ([]*embeddings.Embedding, error) {
	embs := make([]*embeddings.Embedding, 0, len(e.Data))
	for _, d := range e.Data {
		floats := make([]float64, len(d.Embedding))
		copy(floats, d.Embedding)
		emb := &embeddings.Embedding{
			Vector: floats,
		}
		embs = append(embs, emb)
	}
	return embs, nil
}

// Usage tracks API token usage.
type Usage struct {
	TotalTokens int `json:"total_tokens"`
}

// DataGen is a generic struct used for deserializing vector embeddings.
type DataGen[T any] struct {
	Object    string `json:"object"`
	Index     int    `json:"index"`
	Embedding T      `json:"embedding"`
}

// EmbeddingResponseGen is a generic struct used for deserializing API response.
type EmbeddingResponseGen[T any] struct {
	Object string       `json:"object"`
	Data   []DataGen[T] `json:"data"`
	Model  Model        `json:"model"`
	Usage  Usage        `json:"usage"`
}

// toEmbeddingResp decodes the raw API response,
// parses it into a slice of embeddings and returns it.
func toEmbeddingResp[T any](resp io.Reader) (*EmbeddingResponse, error) {
	data := new(T)
	if err := json.NewDecoder(resp).Decode(data); err != nil {
		return nil, err
	}

	switch e := any(data).(type) {
	case *EmbeddingResponseGen[embeddings.Base64]:
		embData := make([]Data, 0, len(e.Data))
		for _, d := range e.Data {
			floats, err := d.Embedding.Decode()
			if err != nil {
				return nil, err
			}
			embData = append(embData, Data{
				Object:    d.Object,
				Index:     d.Index,
				Embedding: floats,
			})
		}
		return &EmbeddingResponse{
			Object: e.Object,
			Data:   embData,
			Model:  e.Model,
			Usage:  e.Usage,
		}, nil
	case *EmbeddingResponseGen[[]float64]:
		embData := make([]Data, 0, len(e.Data))
		for _, d := range e.Data {
			embData = append(embData, Data(d))
		}
		return &EmbeddingResponse{
			Object: e.Object,
			Data:   embData,
			Model:  e.Model,
			Usage:  e.Usage,
		}, nil
	}

	return nil, ErrInValidData
}

// Embed returns embeddings for every object in EmbeddingRequest.
func (c *Client) Embed(ctx context.Context, embReq *EmbeddingRequest) ([]*embeddings.Embedding, error) {
	u, err := url.Parse(c.opts.BaseURL + "/" + c.opts.Version + "/embeddings")
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
		request.WithBearer(c.opts.APIKey),
	}

	req, err := request.NewHTTP(ctx, http.MethodPost, u.String(), body, options...)
	if err != nil {
		return nil, err
	}

	resp, err := request.Do[APIError](c.opts.HTTPClient, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var embs *EmbeddingResponse

	switch embReq.EncodingFormat {
	case EncodingBase64:
		embs, err = toEmbeddingResp[EmbeddingResponseGen[embeddings.Base64]](resp.Body)
	case EncodingNone, "":
		embs, err = toEmbeddingResp[EmbeddingResponseGen[[]float64]](resp.Body)
	default:
		return nil, ErrUnsupportedEncoding
	}
	if err != nil {
		return nil, err
	}

	return embs.ToEmbeddings()
}
