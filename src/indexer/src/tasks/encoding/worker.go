package encoding

import (
	"encoding/json"
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
	jsonDocPath := constants.StorageJsonDocsDir + "/" + docId + ".json"

	os.MkdirAll(docPath, os.ModePerm)

	encodedDocument := make([]core.Vector, 0)

	passages := make([]Passage, len(task.Passages))

	for i, passage := range task.Passages {
		passageId := passage.ID.Hex()
		passageText := passage.Text

		if len(passageText) == 0 {
			continue
		}

		encoded := bert.GetEmbeddings(passageText)
		if encoded == nil {
			continue
		}
		passages[i] = Passage{
			PassageId:  passageId,
			Embeddings: encoded,
		}
		encodedDocument = append(encodedDocument, encoded...)
	}

	if len(encodedDocument) == 0 {
		return
	}

	avgEmb, err := bert.GetAvgEmbedding(encodedDocument)
	core.ErrPanic(err)

	err = storage.WriteStructToFile(avgDocPath, avgEmb)
	core.ErrPanic(err)

	storedDocument := StoredDocument{
		DocId:    docId,
		Passages: passages,
	}
	err = storage.WriteStructToFile(docPath, storedDocument)
	core.ErrPanic(err)
	jsonDoc, err := json.MarshalIndent(storedDocument, "", " ")
	core.ErrPanic(err)
	err = os.WriteFile(jsonDocPath, jsonDoc, os.ModePerm)
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
