package document

// Document stores data and associated metadata.
type Document struct {
	Content  string         `json:"content"`
	Metadata map[string]any `json:"metadata"`
}

// Splitter splits documents into chunks.
type Splitter interface {
	Split(Document) []string
}
