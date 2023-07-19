package encoding

import (
	"fmt"
	"indexer/src/util/db"
)

func RunEncodingTask() {
	documentCount := db.CountCrawlPages()
	tasks := make(chan db.PageCrawl, 10)
	results := make(chan string, documentCount)
	go db.GetCrawlPages(tasks)
	go startEncodingWorker(tasks, results)
	go startEncodingWorker(tasks, results)

	for i := 0; i < documentCount; i++ {
		<-results
	}

	fmt.Println("[RunEncodingTask] finished encoding")

}
