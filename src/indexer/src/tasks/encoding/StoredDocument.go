package encoding

import (
	"indexer/src/util/core"
)

type Passage struct {
	PassageId  string        `json:"passageId"`
	Embeddings []core.Vector `json:"embeddings"`
}

type StoredDocument struct {
	DocId    string    `json:"docId"`
	Passages []Passage `json:"passages"`
}
