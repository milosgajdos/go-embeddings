package openai

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
)

// Usage tracks API token usage.
type Usage struct {
	PromptTokens int `json:"prompt_tokens"`
	TotalTokens  int `json:"total_tokens"`
}

// Embedding is openai API vector embedding.
type Embedding struct {
	Object    string    `json:"object"`
	Index     int       `json:"index"`
	Embedding []float64 `json:"embedding"`
}

// EmbeddingString is base64 encoded embedding.
type EmbeddingString string

func (s EmbeddingString) Decode() ([]float64, error) {
	decoded, err := base64.StdEncoding.DecodeString(string(s))
	if err != nil {
		return nil, err
	}

	if len(decoded)%8 != 0 {
		return nil, fmt.Errorf("invalid base64 encoded string length")
	}

	floats := make([]float64, len(decoded)/8)

	for i := 0; i < len(floats); i++ {
		bits := binary.LittleEndian.Uint64(decoded[i*8 : (i+1)*8])
		floats[i] = math.Float64frombits(bits)
	}

	return floats, nil
}

// EmbeddingRequest is serialized and sent to the API server.
type EmbeddingRequest struct {
	Input          any            `json:"input"`
	Model          Model          `json:"model"`
	User           string         `json:"user"`
	EncodingFormat EncodingFormat `json:"encoding_format,omitempty"`
}

// Data is used for deserializing response data.
type Data[T any] struct {
	Object    string `json:"object"`
	Index     int    `json:"index"`
	Embedding T      `json:"embedding"`
}

// EmbeddingResponse is the API response from a Create embeddings request.
type EmbeddingResponse[T any] struct {
	Object string    `json:"object"`
	Data   []Data[T] `json:"data"`
	Model  Model     `json:"model"`
	Usage  Usage     `json:"usage"`
}

func ToEmbeddings[T any](resp io.Reader) ([]*Embedding, error) {
	data := new(T)
	if err := json.NewDecoder(resp).Decode(data); err != nil {
		return nil, err
	}

	switch e := any(data).(type) {
	case *EmbeddingResponse[EmbeddingString]:
		embs := make([]*Embedding, 0, len(e.Data))
		for _, d := range e.Data {
			floats, err := d.Embedding.Decode()
			if err != nil {
				return nil, err
			}
			emb := &Embedding{
				Object:    d.Object,
				Index:     d.Index,
				Embedding: floats,
			}
			embs = append(embs, emb)
		}
		return embs, nil
	case *EmbeddingResponse[[]float64]:
		embs := make([]*Embedding, 0, len(e.Data))
		for _, d := range e.Data {
			emb := &Embedding{
				Object:    d.Object,
				Index:     d.Index,
				Embedding: d.Embedding,
			}
			embs = append(embs, emb)
		}
		return embs, nil
	}

	return nil, ErrInValidData
}

// CreateEmbeddings returns embeddings for every object in EmbeddingRequest.
func (c *Client) CreateEmbeddings(ctx context.Context, embReq *EmbeddingRequest) ([]*Embedding, error) {
	u, err := url.Parse(c.baseURL + "/embeddings")
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

	switch embReq.EncodingFormat {
	case EncodingBase64:
		return ToEmbeddings[EmbeddingResponse[EmbeddingString]](resp.Body)
	case EncodingFloat:
		return ToEmbeddings[EmbeddingResponse[[]float64]](resp.Body)
	}

	return nil, fmt.Errorf("unknown encoding: %v", embReq.EncodingFormat)
}
