package main

import (
	"indexer/src/tasks/clustering"
	"indexer/src/tasks/encoding"
	"indexer/src/util/storage"
	"time"
)

func main() {
	storage.CreateStorageDirs()
	for {
		encoding.RunEncodingTask()
		clustering.RunClusteringTask()
		time.Sleep((60 * 24) * time.Minute)
	}
}
