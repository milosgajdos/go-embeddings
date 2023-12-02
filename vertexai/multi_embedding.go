package vertexai

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/milosgajdos/go-embeddings/request"
)

// MultiEmbeddingRequest is multimodal embedding request.
type MultiEmbeddingRequest struct {
	Instances []MultiInstance `json:"instances"`
	Params    *MultiParams    `json:"parameters,omitempty"`
}

// MultiInstance contains the request payload.
type MultiInstance struct {
	Text  *string `json:"text,omitempty"`
	Image any     `json:"image,omitempty"`
}

// ImageGCS contains GCS URI to image file.
type ImageGCS struct {
	URI string `json:"gcsUri"`
}

// ImageBase64 contains image encoded as base64 string.
type ImageBase64 struct {
	Bytes string `json:"bytesBase64Encoded"`
}

// MultiParams are additional API request parameters.
type MultiParams struct {
	Dimension uint `json:"dimension"`
}

// MultiEmbedddingResponse received from API.
type MultiEmbedddingResponse struct {
	Predictions []MultiPrediction `json:"predictions"`
	ModelID     string            `json:"deployedModelId"`
}

// MultiPrediction for a given request.
type MultiPrediction struct {
	Image []float64 `json:"imageEmbedding"`
	Text  []float64 `json:"textEmbedding"`
}

// MultiEmbeddings returns multimodal embeddings for every object in EmbeddingRequest.
func (c *Client) MultiEmbeddings(ctx context.Context, embReq *MultiEmbeddingRequest) (*MultiEmbedddingResponse, error) {
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

	if c.token == "" {
		var err error
		c.token, err = GetToken(c.tokenSrc)
		if err != nil {
			return nil, err
		}
	}

	options := []request.Option{
		request.WithBearer(c.token),
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

	e := new(MultiEmbedddingResponse)
	if err := json.NewDecoder(resp.Body).Decode(e); err != nil {
		return nil, err
	}

	return e, nil
}
