package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/milosgajdos/go-embeddings/cohere"
)

var (
	input     string
	model     string
	truncate  string
	inputType string
)

func init() {
	flag.StringVar(&input, "input", "what is life", "input data")
	flag.StringVar(&model, "model", cohere.EnglishV3.String(), "model name")
	flag.StringVar(&truncate, "truncate", cohere.NoneTrunc.String(), "truncate type")
	flag.StringVar(&inputType, "input-type", cohere.ClusteringInput.String(), "input type")
}

func main() {
	flag.Parse()

	c := cohere.NewClient()

	embReq := &cohere.EmbeddingRequest{
		Texts:     []string{input},
		Model:     cohere.Model(model),
		InputType: cohere.InputType(inputType),
		Truncate:  cohere.Truncate(truncate),
	}

	embs, err := c.Embeddings(context.Background(), embReq)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("got %d embeddings", len(embs))
}
