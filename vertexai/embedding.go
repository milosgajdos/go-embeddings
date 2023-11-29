package vertexai

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
// https://cloud.google.com/vertex-ai/docs/generative-ai/embeddings/get-text-embeddings#generative-ai-get-text-embedding-drest
type EmbeddingRequest struct {
	Instances []Instance `json:"instances"`
	Params    Params     `json:"parameters"`
}

// Params are additional API request parameters passed via body.
type Params struct {
	// If set to false, text that exceeds the token limit (3.072)
	// causes the request to fail. The default value is true
	AutoTruncate bool `json:"autoTruncate"`
}

// NOTE: Title is only valid with TaskType set to RetrDocTask
// https://cloud.google.com/vertex-ai/docs/generative-ai/embeddings/get-text-embeddings#api_changes_to_models_released_on_or_after_august_2023
type Instance struct {
	TaskType TaskType `json:"task_type"`
	Title    string   `json:"title,omitempty"`
	Content  string   `json:"content"`
}

// EmbedddingResponse received from API endpoint.
// https://cloud.google.com/vertex-ai/docs/generative-ai/model-reference/text-embeddings#response_body
type EmbedddingResponse struct {
	Predictions []Predictions  `json:"predictions"`
	Metadata    map[string]any `json:"metadata"`
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

func ToEmbeddings(r io.Reader) ([]*Embedding, error) {
	var resp EmbedddingResponse
	if err := json.NewDecoder(r).Decode(&resp); err != nil {
		return nil, err
	}
	embs := make([]*Embedding, 0, len(resp.Predictions))
	for _, p := range resp.Predictions {
		floats := make([]float64, len(p.Embeddings.Values))
		copy(floats, p.Embeddings.Values)
		emb := &Embedding{
			Vector: floats,
		}
		embs = append(embs, emb)
	}
	return embs, nil
}

// Embeddings returns embeddings for every object in EmbeddingRequest.
func (c *Client) Embeddings(ctx context.Context, embReq *EmbeddingRequest) ([]*Embedding, error) {
	u, err := url.Parse(c.baseURL + "/" + c.projectID + "/" + ModelURI + "/" + c.modelID + EmbedAction)
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
