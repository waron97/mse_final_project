package clustering

import (
	"indexer/src/util/core"
	"indexer/src/util/storage"
	"strconv"
)

func getClosestCentroid(document core.Vector, centroids []core.Vector) int {
	centroidId := -1
	max := 0.0

	for i, c := range centroids {
		cos := c.CosSim(document)
		if cos > max {
			max = cos
			centroidId = i
		}
	}
	return centroidId
}

func clusterDocuments(documents []*Document, clusters []core.Vector) {
	constants := core.GetConstants()
	for _, doc := range documents {
		centroidId := getClosestCentroid(doc.Embedding, clusters)

		if centroidId == -1 {
			panic("no centroid found")
		}

		centroidPath := constants.StorageClustersDir + "/centroid_" + strconv.Itoa(centroidId)
		entry := NewDocEmbedding(doc.DocID)

		var centroidData []*DocEmbedding
		err := storage.ReadStructFromFile(centroidPath, &centroidData)
		if err != nil {
			centroidData = make([]*DocEmbedding, 0)
		}

		if !centroidContains(centroidData, doc.DocID) {
			centroidData = append(centroidData, entry)
			err = storage.WriteStructToFile(centroidPath, centroidData)
			core.ErrPanic(err)
		}

	}
}
