package prerank

import "ranker/src/util/core"

type DocEmbedding struct {
	DocId string
}

type Cluster struct {
	ID        int
	Centroid  core.Vector
	Documents []string
}
