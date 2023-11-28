package openai

// Model is embedding model.
type Model string

const (
	TextAdaV2 Model = "text-embedding-ada-002"
)

// EncodingFormat for embedding API requests.
type EncodingFormat string

const (
	EncodingFloat  EncodingFormat = "float"
	EncodingBase64 EncodingFormat = "base64"
)