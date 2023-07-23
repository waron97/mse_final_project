package clustering

import (
	"fmt"
	"indexer/src/util/core"
	"os"
)

func RunClusteringTask() {
	fmt.Println("[RunClusteringTask] Running clustering task")
	constants := core.GetConstants()
	documents := loadDocuments()
	fmt.Println("[RunClusteringTask] Documents loaded", len(documents))
	documents = getDocumentSubset(documents, len(documents))

	if len(documents) < 1000 {
		fmt.Println("[RunClusteringTask] Not enough documents to cluster")
		return
	}

	var centroids []core.Vector

	if f, err := os.Stat(constants.StorageClusterMapPath); err == nil && !f.IsDir() {
		data := ReadClusters()
		for _, row := range data {
			centroids = append(centroids, row.Centroid)
		}
		fmt.Println("[RunClusteringTask] Centroids loaded from disc")
	} else {
		centroids = cluster(getDocumentEmbeddings(documents), constants.ClusterCount)
		fmt.Println("[RunClusteringTask] Centroids found")
		writeClusterMap(centroids)
	}

	clusterDocuments(documents, centroids)
	fmt.Println("[RunClusteringTask] Documents clustered")
}
