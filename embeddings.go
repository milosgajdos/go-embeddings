package embeddings

import (
	"context"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"math"
)

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

// Base64 is base64 encoded embedding string.
type Base64 string

// Decode decodes base64 encoded string into a slice of floats.
func (s Base64) Decode() ([]float64, error) {
	decoded, err := base64.StdEncoding.DecodeString(string(s))
	if err != nil {
		return nil, err
	}

	if len(decoded)%8 != 0 {
		return nil, fmt.Errorf("invalid base64 encoded string length")
	}

	floats := make([]float64, len(decoded)/8)

	for i := 0; i < len(floats); i++ {
		bits := binary.LittleEndian.Uint64(decoded[i*8 : (i+1)*8])
		floats[i] = math.Float64frombits(bits)
	}

	return floats, nil
}
