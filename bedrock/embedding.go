package bedrock

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/milosgajdos/go-embeddings"
)

type Request struct {
	InputText string `json:"inputText"`
}

type Response struct {
	Embedding           []float64 `json:"embedding"`
	InputTextTokenCount int       `json:"inputTextTokenCount"`
}

func (e *Response) ToEmbeddings() ([]*embeddings.Embedding, error) {
	vals := make([]float64, len(e.Embedding))
	copy(vals, e.Embedding)
	return []*embeddings.Embedding{
		{Vector: vals},
	}, nil
}

// Embed returns embeddings for every object in EmbeddingRequest.
func (c *Client) Embed(ctx context.Context, embReq *Request) ([]*embeddings.Embedding, error) {
	payload, err := json.Marshal(embReq)
	if err != nil {
		return nil, err
	}

	resp, err := c.opts.Client.InvokeModel(ctx, &bedrockruntime.InvokeModelInput{
		Body:        payload,
		ModelId:     aws.String(c.opts.ModelID),
		ContentType: aws.String("application/json"),
	})
	if err != nil {
		return nil, err
	}

	var embs Response
	if err = json.Unmarshal(resp.Body, &embs); err != nil {
		return nil, nil
	}

	return embs.ToEmbeddings()
}
