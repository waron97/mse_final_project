package prerank

import "ranker/src/util/core"

type ByDistanceFromQuery struct {
	Clusters []Cluster
	Query    core.Vector
}

func (b ByDistanceFromQuery) Len() int {
	return len(b.Clusters)
}

func (b ByDistanceFromQuery) Less(i, j int) bool {
	return b.Clusters[i].Centroid.EuclideanDistance(b.Query) < b.Clusters[j].Centroid.EuclideanDistance(b.Query)
}

func (b ByDistanceFromQuery) Swap(i, j int) {
	b.Clusters[i], b.Clusters[j] = b.Clusters[j], b.Clusters[i]
}
