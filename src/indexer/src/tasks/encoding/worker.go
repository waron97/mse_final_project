package encoding

import (
	"fmt"
	"indexer/src/util/bert"
	"indexer/src/util/core"
	"indexer/src/util/db"
	"indexer/src/util/storage"
	"os"
)

func processDocument(task db.PageCrawl) {
	docId := task.ID.Hex()
	constants := core.GetConstants()
	avgDocPath := constants.StorageAverageDocsDir + "/" + docId
	docPath := constants.StorageDocsDir + "/" + docId

	os.MkdirAll(docPath, os.ModePerm)

	encodedDocument := make([]bert.Vector, 0)

	for _, passage := range task.Passages {
		passageId := passage.ID.Hex()
		passagePath := constants.StorageDocsDir + "/" + fmt.Sprintf("%s/%s", docId, passageId)
		passageText := passage.Text

		if len(passageText) == 0 {
			continue
		}

		encoded := bert.GetEmbedding(passageText)
		err := storage.WriteStructToFile(passagePath, encoded)
		core.ErrPanic(err)
		encodedDocument = append(encodedDocument, encoded...)
	}

	if len(encodedDocument) == 0 {
		return
	}

	avgEmb, err := bert.GetAvgEmbedding(encodedDocument)
	core.ErrPanic(err)

	err = storage.WriteStructToFile(avgDocPath, avgEmb)
	core.ErrPanic(err)

	db.MarkPageIndexed(task.ID.Hex())
}

func startEncodingWorker(tasks chan db.PageCrawl, results chan string) {
	for {
		task := <-tasks

		if len(task.BodyText) == 0 {
			results <- task.ID.Hex()
			continue
		}

		processDocument(task)
		results <- task.ID.Hex()
	}
}
