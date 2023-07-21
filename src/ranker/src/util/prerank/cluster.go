package prerank

import "ranker/src/util/core"

type DocEmbedding struct {
	DocId     string
	Embedding core.Vector
}

type Cluster struct {
	ID        int
	Centroid  core.Vector
	Documents []string
}
