package embeddings

import "context"

// Embedder fetches embeddings.
type Embedder[T any] interface {
	// Embeddings fetches embeddings and returns them.
	Embed(context.Context, T) ([]*Embedding, error)
}

// Embedding is vector embedding.
type Embedding struct {
	Vector []float64 `json:"vector"`
}

// ToFloat32 returns Embedding verctor as a slice of float32.
func (e Embedding) ToFloat32() []float32 {
	floats := make([]float32, len(e.Vector))
	for i, f := range e.Vector {
		floats[i] = float32(f)
	}
	return floats
}
