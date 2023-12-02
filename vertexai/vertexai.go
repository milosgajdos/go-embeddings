package vertexai

// Model is embedding model.
type Model string

const (
	// EmbedGeckoV* are English language embedding models
	// See for the information about the latest available version
	// https://cloud.google.com/vertex-ai/docs/generative-ai/learn/model-versioning#latest-version
	EmbedGeckoV1     Model = "textembedding-gecko@001"
	EmbedGeckoV2     Model = "textembedding-gecko@002"
	EmbedGeckoLatest Model = "textembedding-gecko@latest"
	// EmbedMultiGecko is a multilanguage embeddings model
	EmbedMultiGecko Model = "multimodalembedding@001"
)

// String implements stringer.
func (m Model) String() string {
	return string(m)
}

// TaskType is embedding task type.
// It can be used to improve the embedding quality
// when targeting a specific task.
// See: https://cloud.google.com/vertex-ai/docs/generative-ai/embeddings/get-text-embeddings
type TaskType string

const (
	RetrQueryTask      TaskType = "RETRIEVAL_QUERY"
	RetrDocTask        TaskType = "RETRIEVAL_DOCUMENT"
	SemanticSimTask    TaskType = "SEMANTIC_SIMILARITY"
	ClassificationTask TaskType = "CLASSIFICATION"
	ClusteringTask     TaskType = "CLUSTERING"
)

// String implements stringer.
func (t TaskType) String() string {
	return string(t)
}
