package encoding

import (
	"indexer/src/util/bert"
	"indexer/src/util/core"
	"indexer/src/util/db"
	"indexer/src/util/storage"
)

func processDocument(task db.PageCrawl) {
	docId := task.ID.Hex()
	constants := core.GetConstants()
	docPath := constants.StorageDocsDir + "/" + docId
	avgDocPath := constants.StorageAverageDocsDir + "/" + docId

	emb := bert.GetEmbedding(task.BodyText)
	avgEmb, err := bert.GetAvgEmbedding(emb)
	core.ErrPanic(err)

	err = storage.WriteStructToFile(docPath, emb)
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
