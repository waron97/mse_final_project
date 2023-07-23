package main

import (
	"indexer/src/tasks/clustering"
	"indexer/src/tasks/encoding"
	"indexer/src/util/storage"
)

func main() {
	storage.CreateStorageDirs()
	doRun := false
	if doRun {
		encoding.RunEncodingTask()
		clustering.RunClusteringTask()
	}
}
