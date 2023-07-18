package util

type Passage struct {
	PassageId  string
	Embeddings []Vector
}

type Document struct {
	DocId    string
	Passages []Passage
}
