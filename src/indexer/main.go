package main

import (
	buildfulldocs "indexer/src/tasks/buildFullDocs"
	"indexer/src/tasks/clustering"
	"indexer/src/tasks/encoding"
	"indexer/src/util/storage"
)

func main() {
	storage.CreateStorageDirs()
	buildfulldocs.Build()
	encoding.RunEncodingTask()
	clustering.RunClusteringTask()
}
