package main

import (
	"indexer/src/tasks/clustering"
	"indexer/src/tasks/encoding"
	tojson "indexer/src/tasks/toJson"
	"indexer/src/util/storage"
)

func main() {
	storage.CreateStorageDirs()
	encoding.RunEncodingTask()
	tojson.IndexToJson()
	clustering.RunClusteringTask()
}
