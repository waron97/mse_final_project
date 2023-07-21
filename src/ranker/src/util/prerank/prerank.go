package prerank

import (
	"ranker/src/util/core"
	"sort"
)

func Prerank(query []core.Vector, clusters []Cluster, n int) []string {
	queryVector, _ := core.GetAvgEmbedding(query)
	sorter := ByDistanceFromQuery{
		Clusters: clusters,
		Query:    queryVector,
	}
	sort.Sort(sorter)
	sorted := sorter.Clusters
	results := make([]string, 0)
	for _, cluster := range sorted {
		results = append(results, cluster.Documents...)
		if len(results) >= n {
			break
		}
	}
	return results
}
