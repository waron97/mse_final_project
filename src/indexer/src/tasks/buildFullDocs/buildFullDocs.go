package buildfulldocs

import (
	"fmt"
	"indexer/src/util/core"
	"indexer/src/util/storage"
	"os"
	"path/filepath"
)

func processDocument(docId string, constants *core.Constants, done chan bool) {
	docPath := filepath.Join(constants.StorageFullDocsDir, docId)
	doc := *NewDocumentFromId(docId)
	err := storage.WriteStructToFile(docPath, doc)
	core.ErrPanic(err)
	done <- true
}

func startDocWorker(tasks chan os.DirEntry, done chan bool) {
	constant := core.GetConstants()
	for {
		file, ok := <-tasks
		if !ok {
			return
		}
		processDocument(file.Name(), constant, done)
	}

}

func Build() {
	fmt.Println("Building full docs...")
	constants := core.GetConstants()
	files, err := os.ReadDir(constants.StorageAverageDocsDir)
	fmt.Println("Number of files to process:", len(files))
	core.ErrPanic(err)

	tasks := make(chan os.DirEntry, len(files))
	done := make(chan bool)

	for _, file := range files {
		tasks <- file
	}

	close(tasks)

	fmt.Println("Starting workers...")

	go startDocWorker(tasks, done)
	go startDocWorker(tasks, done)
	go startDocWorker(tasks, done)

	for i := 0; i < len(files); i++ {
		<-done
		if i%100 == 0 && i != 0 {
			fmt.Printf("Processed %d documents\n", i+1)
		}
	}

}