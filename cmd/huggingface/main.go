package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/milosgajdos/go-embeddings/cohere"
	"github.com/milosgajdos/go-embeddings/huggingface"
)

var (
	input string
	model string
	wait  bool
)

func init() {
	flag.StringVar(&input, "input", "what is life", "input data")
	flag.StringVar(&model, "model", string(cohere.EnglishV3), "model name")
	flag.BoolVar(&wait, "wait", false, "wait for model to start")
}

func main() {
	flag.Parse()

	c := huggingface.NewClient().
		WithModel(model)

	embReq := &huggingface.EmbeddingRequest{
		Inputs: []string{input},
		Options: huggingface.Options{
			WaitForModel: &wait,
		},
	}

	embResp, err := c.Embeddings(context.Background(), embReq)
	if err != nil {
		log.Fatal(err)
	}

	embs, err := huggingface.ToEmbeddings(embResp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("got %d embeddings", len(embs))
}
