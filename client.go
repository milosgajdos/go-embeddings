package embeddings

import "context"

// Embedder fetches embeddings.
type Embedder[T any] interface {
	// Embeddings fetches embeddings and returns them.
	Embeddings(context.Context, T) ([]*Embedding, error)
}
