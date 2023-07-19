package clustering

import (
	"fmt"
	"indexer/src/util/core"
)

func RunClusteringTask() {
	fmt.Println("Running clustering task")
	constants := core.GetConstants()
	documents := loadDocuments()
	documents = getDocumentSubset(documents, len(documents))

	centroids := cluster(getDocumentEmbeddings(documents), constants.ClusterCount)
	writeClusterMap(centroids)
	clusterDocuments(documents, centroids)
	fmt.Println("Documents clustered")
}
