package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/milosgajdos/go-embeddings/voyage"
)

var (
	input      string
	model      string
	truncation bool
	inputType  string
)

func init() {
	flag.StringVar(&input, "input", "what is life", "input data")
	flag.StringVar(&model, "model", voyage.VoyageV2.String(), "model name")
	flag.StringVar(&inputType, "input-type", voyage.DocInput.String(), "input type")
	flag.BoolVar(&truncation, "truncate", false, "truncate type")
}

func main() {
	flag.Parse()

	c := voyage.NewClient()

	embReq := &voyage.EmbeddingRequest{
		Input:      []string{input},
		Model:      voyage.Model(model),
		InputType:  voyage.InputType(inputType),
		Truncation: truncation,
	}

	embs, err := c.Embed(context.Background(), embReq)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("got %d embeddings", len(embs))
}
