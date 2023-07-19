package clustering

import "indexer/src/util/core"

type DocEmbedding struct {
	DocId     string
	Embedding core.Vector
}

func NewDocEmbedding(docId string, embedding core.Vector) *DocEmbedding {
	return &DocEmbedding{
		DocId:     docId,
		Embedding: embedding,
	}
}
