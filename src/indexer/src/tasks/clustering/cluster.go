package clustering

import (
	"fmt"
	"indexer/src/util/core"
	"os"
)

func RunClusteringTask() {
	fmt.Println("Running clustering task")
	constants := core.GetConstants()
	documents := loadDocuments()
	fmt.Println("Documents loaded", len(documents))
	documents = getDocumentSubset(documents, len(documents))

	var centroids []core.Vector

	if f, err := os.Stat(constants.StorageClusterMapPath); err == nil && !f.IsDir() {
		data := ReadClusters()
		for _, row := range data {
			centroids = append(centroids, row.Centroid)
		}
	} else {
		centroids = cluster(getDocumentEmbeddings(documents), constants.ClusterCount)
		fmt.Println("Centroids found")
		writeClusterMap(centroids)
	}

	clusterDocuments(documents, centroids)
	fmt.Println("Documents clustered")
}
