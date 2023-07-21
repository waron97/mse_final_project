package prerank

import (
	"ranker/src/util/core"
	"ranker/src/util/storage"
	"strconv"
)

func ReadClusters() []Cluster {
	constants := core.GetConstants()
	var clusterMap []core.Vector
	storage.ReadStructFromFile(constants.StorageClusterMapPath, &clusterMap)
	var clusters []Cluster
	for i, cluster := range clusterMap {
		clusterId := strconv.Itoa(i)
		clusterPath := constants.StorageClustersDir + "/centroid_" + clusterId
		var documents []DocEmbedding
		storage.ReadStructFromFile(clusterPath, &documents)
		docIds := make([]string, len(documents))
		for i, doc := range documents {
			docIds[i] = doc.DocId
		}
		clusters = append(clusters, Cluster{
			ID:        i,
			Centroid:  cluster,
			Documents: docIds,
		})
	}
	return clusters
}
