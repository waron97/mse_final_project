package storage

import (
	"indexer/src/util/core"
	"os"
)

func CreateStorageDirs() {
	constants := core.GetConstants()
	os.MkdirAll(constants.StorageBaseDir, os.ModePerm)
	os.MkdirAll(constants.StorageDocsDir, os.ModePerm)
	os.MkdirAll(constants.StorageAverageDocsDir, os.ModePerm)
}
