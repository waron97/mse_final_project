package tojson

import (
	"encoding/json"
	"fmt"
	"indexer/src/tasks/encoding"
	"indexer/src/util/core"
	"indexer/src/util/storage"
	"os"
	"strconv"
)

func startWorker(tasks chan string, contants *core.Constants, results chan string) {
	for {
		fileName, ok := <-tasks
		if !ok {
			return
		}
		path := contants.StorageFullDocsDir + "/" + fileName
		var doc encoding.StoredDocument
		err := storage.ReadStructFromFile(path, &doc)
		core.ErrPanic(err)
		jsonDoc, err := json.MarshalIndent(doc, "", " ")
		core.ErrPanic(err)
		target := contants.StorageJsonDocsDir + "/" + doc.DocId + ".json"
		os.WriteFile(target, jsonDoc, os.ModePerm)
		results <- doc.DocId
	}
}

func IndexToJson() {
	fmt.Println("[IndexToJson] start")
	constants := core.GetConstants()
	numWorkers := 10
	files, err := os.ReadDir(constants.StorageFullDocsDir)
	core.ErrPanic(err)
	target := constants.StorageJsonDocsDir
	os.MkdirAll(target, os.ModePerm)

	tasks := make(chan string, len(files))
	results := make(chan string, len(files))

	for _, file := range files {
		tasks <- file.Name()
	}
	close(tasks)

	for i := 0; i < numWorkers; i++ {
		go startWorker(tasks, constants, results)
	}

	for i := 0; i < len(files); i++ {
		<-results
		if i%100 == 0 && i != 0 {

			fmt.Println("Indexed " + strconv.Itoa(i) + " documents")
		}
	}

}
