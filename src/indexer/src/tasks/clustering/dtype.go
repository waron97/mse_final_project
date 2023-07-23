package clustering

import "indexer/src/util/core"

type DocEmbedding struct {
	DocId string
}

func NewDocEmbedding(docId string) *DocEmbedding {
	return &DocEmbedding{
		DocId: docId,
	}
}

type Cluster struct {
	ID        int
	Centroid  core.Vector
	Documents []string
}
