package encoding

import (
	"fmt"
	"indexer/src/util/bert"
	"indexer/src/util/db"
	"time"
)

func RunEncodingTask() {
	for !bert.IsAlive() {
		fmt.Println("[RunEncodingTask] waiting for bert")
		time.Sleep(2 * time.Second)
	}
	documentCount := db.CountCrawlPages()
	tasks := make(chan db.PageCrawl, 10)
	results := make(chan string, documentCount)
	now := time.Now()
	go db.GetCrawlPages(tasks)
	go startEncodingWorker(tasks, results)
	go startEncodingWorker(tasks, results)

	for i := 0; i < documentCount; i++ {
		<-results
		if i%100 == 0 && i != 0 {
			timeElapsed := time.Since(now)
			fmt.Printf("[RunEncodingTask] %d/%d :: %s\n", i, documentCount, timeElapsed)
		}
	}

	fmt.Println("[RunEncodingTask] finished encoding")

}
