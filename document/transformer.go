package document

// Transformer transform a Document and returns it.
type Transformer interface {
	Transform(Document) Document
}
