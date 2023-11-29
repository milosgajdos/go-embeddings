package cohere

// Model is an embedding model.
type Model string

const (
	EnglishV3        Model = "embed-english-v3.0"
	MultiLingV3      Model = "embed-multilingual-v3.0"
	EnglishLightV3   Model = "embed-english-light-v3.0"
	MultiLingLightV3 Model = "embed-multilingual-light-v3.0"
	EnglishV2        Model = "embed-english-v2.0"
	EnglishLightV2   Model = "embed-english-light-v2.0"
	MultiLingV2      Model = "embed-multilingual-v2.0"
)

// InputType is an embedding input type.
type InputType string

const (
	SearchDocInput      InputType = "search_document"
	SearchQueryInput    InputType = "search_query"
	ClassificationInput InputType = "classification"
	ClusteringInput     InputType = "clustering"
)

// Truncate controls input truncating.
// It controls how the API handles inputs
// longer than the maximum token length (recommended: <512)
type Truncate string

const (
	StartTrunc Truncate = "START"
	EndTrunc   Truncate = "END"
	NoneTrunc  Truncate = "NONE"
)
