package embedding

// Embedding is a vector embedding.
type Embedding interface {
	// Values returns vector embedding values.
	Values() []float64
}
