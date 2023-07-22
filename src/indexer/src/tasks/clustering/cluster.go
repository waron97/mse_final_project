package clustering

import (
	"fmt"
	"indexer/src/util/core"
)

func RunClusteringTask() {
	fmt.Println("Running clustering task")
	constants := core.GetConstants()
	documents := loadDocuments()
	fmt.Println("Documents loaded", len(documents))
	documents = getDocumentSubset(documents, len(documents))
	centroids := cluster(getDocumentEmbeddings(documents), constants.ClusterCount)
	fmt.Println("Centroids found")
	writeClusterMap(centroids)
	clusterDocuments(documents, centroids)
	fmt.Println("Documents clustered")
}
