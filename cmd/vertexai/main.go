package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/milosgajdos/go-embeddings/vertexai"
	"golang.org/x/oauth2/google"
)

var (
	input    string
	model    string
	truncate bool
	taskType string
	title    string
)

func init() {
	flag.StringVar(&input, "input", "what is life", "input data")
	flag.StringVar(&model, "model", vertexai.EmbedGeckoV2.String(), "model name")
	flag.BoolVar(&truncate, "truncate", false, "truncate type")
	flag.StringVar(&taskType, "task-type", vertexai.RetrQueryTask.String(), "task type")
	flag.StringVar(&title, "title", "", "title: only relevant for retrival document tasks")
}

func main() {
	flag.Parse()

	ctx := context.Background()

	ts, err := google.DefaultTokenSource(ctx, vertexai.Scopes)
	if err != nil {
		log.Fatalf("token source: %v", err)
	}

	c := vertexai.NewClient().
		WithTokenSrc(ts).
		WithModelID(model)

	embReq := &vertexai.EmbeddingRequest{
		Instances: []vertexai.Instance{
			{
				Content:  input,
				TaskType: vertexai.TaskType(taskType),
				Title:    title,
			},
		},
		Params: vertexai.Params{
			AutoTruncate: truncate,
		},
	}

	embs, err := c.Embeddings(context.Background(), embReq)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("got %d embeddings", len(embs))
}
