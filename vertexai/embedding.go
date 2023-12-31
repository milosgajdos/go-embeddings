package vertexai

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
// https://cloud.google.com/vertex-ai/docs/generative-ai/embeddings/get-text-embeddings#generative-ai-get-text-embedding-drest
type EmbeddingRequest struct {
	Instances []Instance `json:"instances"`
	Params    Params     `json:"parameters"`
}

// NOTE: Title is only valid with TaskType set to RetrDocTask
// https://cloud.google.com/vertex-ai/docs/generative-ai/embeddings/get-text-embeddings#api_changes_to_models_released_on_or_after_august_2023
type Instance struct {
	TaskType TaskType `json:"task_type"`
	Title    string   `json:"title,omitempty"`
	Content  string   `json:"content"`
}

// Params are additional API request parameters passed via body.
type Params struct {
	// If set to false, text that exceeds the token limit (3.072)
	// causes the request to fail. The default value is true
	AutoTruncate bool `json:"autoTruncate"`
}

// EmbedddingResponse received from API endpoint.
// https://cloud.google.com/vertex-ai/docs/generative-ai/model-reference/text-embeddings#response_body
type EmbedddingResponse struct {
	Predictions []Predictions  `json:"predictions"`
	Metadata    map[string]any `json:"metadata"`
}

// ToEmbeddings converts the API response,
// into a slice of embeddings and returns it.
func (e *EmbedddingResponse) ToEmbeddings() ([]*embeddings.Embedding, error) {
	embs := make([]*embeddings.Embedding, 0, len(e.Predictions))
	for _, p := range e.Predictions {
		floats := make([]float64, len(p.Embeddings.Values))
		copy(floats, p.Embeddings.Values)
		emb := &embeddings.Embedding{
			Vector: floats,
		}
		embs = append(embs, emb)
	}
	return embs, nil
}

// Predictions is the generated response
type Predictions struct {
	Embeddings struct {
		Values     []float64  `json:"values"`
		Statistics Statistics `json:"statistics"`
	} `json:"embeddings"`
}

// Statistics define the statistics for a text embedding
type Statistics struct {
	TokenCount int  `json:"token_count"`
	Truncated  bool `json:"truncated"`
}

// Embed returns embeddings for every object in EmbeddingRequest.
func (c *Client) Embed(ctx context.Context, embReq *EmbeddingRequest) ([]*embeddings.Embedding, error) {
	u, err := url.Parse(c.opts.BaseURL + "/" + c.opts.ProjectID + "/" + ModelURI + "/" + c.opts.ModelID + EmbedAction)
	if err != nil {
		return nil, err
	}

	var body = &bytes.Buffer{}
	enc := json.NewEncoder(body)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(embReq); err != nil {
		return nil, err
	}

	if c.opts.Token == "" {
		var err error
		c.opts.Token, err = GetToken(c.opts.TokenSrc)
		if err != nil {
			return nil, err
		}
	}

	options := []request.Option{
		request.WithBearer(c.opts.Token),
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

	e := new(EmbedddingResponse)
	if err := json.NewDecoder(resp.Body).Decode(e); err != nil {
		return nil, err
	}

	return e.ToEmbeddings()
}
