package clustering

import (
	"indexer/src/util/core"
)

func RunClusteringTask() {
	constants := core.GetConstants()
	documents := loadDocuments()
	documents = getDocumentSubset(documents, len(documents))

	centroids := cluster(getDocumentEmbeddings(documents), constants.ClusterCount)
	writeClusterMap(centroids)
	clusterDocuments(documents, centroids)
}
