package clustering

import (
	"indexer/src/util/core"
	"indexer/src/util/storage"
)

func writeClusterMap(clusters []core.Vector) {
	constants := core.GetConstants()
	err := storage.WriteStructToFile(constants.StorageClusterMapPath, clusters)
	core.ErrPanic(err)
}
