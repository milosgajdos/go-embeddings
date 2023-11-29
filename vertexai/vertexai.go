package vertexai

// Model is embedding model.
type Model string

const (
	EmbedGecko      Model = "textembedding-gecko"
	EmbedMultiGecko Model = "multimodalembedding@001"
)

// TaskType is embedding task type.
type TaskType string

const (
	RetrQueryTask      TaskType = "RETRIEVAL_QUERY"
	RetrDocTask        TaskType = "RETRIEVAL_DOCUMENT"
	SemanticSimTask    TaskType = "SEMANTIC_SIMILARITY"
	ClassificationTask TaskType = "CLASSIFICATION"
	ClusteringTask     TaskType = "CLUSTERING"
)
