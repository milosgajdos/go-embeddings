package embeddings

// Embedding is vector embedding.
type Embedding struct {
	Vector []float64 `json:"vector"`
}

// ToFloat32 returns Embedding verctor as a slice of float32.
func (e Embedding) ToFloat32() []float32 {
	floats := make([]float32, 0, len(e.Vector))
	for _, f := range e.Vector {
		floats = append(floats, float32(f))
	}
	return floats
}
