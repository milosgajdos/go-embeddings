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
	flag.StringVar(&input, "input", "", "input data")
	flag.StringVar(&model, "model", string(cohere.EnglishV3), "model name")
	flag.StringVar(&truncate, "truncate", string(cohere.NoneTrunc), "truncate type")
	flag.StringVar(&inputType, "input-type", string(cohere.ClusteringInput), "input type")
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

	embResp, err := c.Embeddings(context.Background(), embReq)
	if err != nil {
		log.Fatal(err)
	}

	embs, err := cohere.ToEmbeddings(embResp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("got %d embeddings", len(embs))
}
