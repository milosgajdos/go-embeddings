package voyage

// Model is an embedding model.
type Model string

const (
	LargeV2        Model = "voyage-large-2"
	CodeV2         Model = "voyage-code-2"
	VoyageV2       Model = "voyage-2"
	LiteV2Instruct Model = "voyage-lite-02-instruct"
)

// String implements stringer.
func (m Model) String() string {
	return string(m)
}

// InputType is an embedding input type.
type InputType string

const (
	NoneInput  InputType = "None"
	QueryInput InputType = "query"
	DocInput   InputType = "document"
)

// String implements stringer.
func (i InputType) String() string {
	return string(i)
}

// EncodingFormat for embedding API requests.
type EncodingFormat string

const (
	EncodingNone EncodingFormat = "None"
	// EncodingBase64 makes Voyage API return embeddings
	// encoded as base64 string
	EncodingBase64 EncodingFormat = "base64"
)

// String implements stringer.
func (f EncodingFormat) String() string {
	return string(f)
}
