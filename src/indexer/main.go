package main

import (
	"indexer/src/tasks/clustering"
	"indexer/src/tasks/encoding"
	"indexer/src/util/storage"
)

func main() {
	storage.CreateStorageDirs()
	encoding.RunEncodingTask()
	clustering.RunClusteringTask()
}
