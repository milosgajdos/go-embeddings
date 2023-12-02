package openai

// Model is embedding model.
type Model string

const (
	TextAdaV2 Model = "text-embedding-ada-002"
)

// String implements stringer.
func (m Model) String() string {
	return string(m)
}

// EncodingFormat for embedding API requests.
type EncodingFormat string

const (
	EncodingFloat EncodingFormat = "float"
	// EncodingBase64 makes OpenAI API return embeddings
	// encoded as base64 string
	EncodingBase64 EncodingFormat = "base64"
)

// String implements stringer.
func (f EncodingFormat) String() string {
	return string(f)
}
