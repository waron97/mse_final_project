package clustering

import (
	"indexer/src/util/core"

	"github.com/muesli/clusters"
	"github.com/muesli/kmeans"
)

func cluster(embeddings []core.Vector, k int) []core.Vector {
	var d clusters.Observations
	for _, x := range embeddings {
		d = append(d, clusters.Coordinates(x))
	}

	km := kmeans.New()
	clusters, err := km.Partition(d, k)
	core.ErrPanic(err)

	centroids := make([]core.Vector, len(clusters))
	for i, c := range clusters {
		centroids[i] = core.Vector(c.Center)
	}

	return centroids
}
