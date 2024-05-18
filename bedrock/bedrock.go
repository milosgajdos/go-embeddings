package bedrock

// Model is embedding model.
type Model string

const (
	TitanTextV1 Model = "amazon.titan-embed-text-v1"
	TitanTextV2 Model = "amazon.titan-embed-text-v2:0"
)

// String implements stringer.
func (m Model) String() string {
	return string(m)
}
