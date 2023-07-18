package main

import (
	"flag"
	"fmt"
	"indexer/internal/util"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// indexingGoRoutine Run indexer.Store() function continuously
func indexingGoRoutine(indexer *util.Store) {
	var wg sync.WaitGroup

	// Create a channel to coordinate termination
	done := make(chan struct{})

	// Handle termination signals (e.g., Ctrl+C)
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh

		// Termination signal received, close the done channel
		close(done)
	}()

	// initial run
	runIndexStore(&wg, indexer)
	fmt.Println("indexing", time.Now())

	ticker := time.NewTicker(1 * time.Hour)

	// run indexer.Store() at the scheduled intervals
	go func() {
		for {
			select {
			case <-done:
				// wait for completion upon receiving Termination signal
				wg.Wait()

				// Exit the goroutine
				fmt.Println("Terminating goroutine.")
				os.Exit(0)

			case <-ticker.C:
				wg.Wait()
				runIndexStore(&wg, indexer)
				fmt.Println("indexing", time.Now())
			}
		}
	}()
	select {}
}

// runIndexStore runs the indexer.Store() function and marks it as completed when finished
func runIndexStore(wg *sync.WaitGroup, indexer *util.Store) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		indexer.Store()
	}()
}

func main() {
	dFlag := flag.String("d", "mongodb://localhost:27017", "DB connection string")
	bFlag := flag.String("b", "../offline-index", "Base directory for index storage")
	cFlag := flag.Bool("c", false, "Assign and calculate document clusters")
	nFlag := flag.Int("n", 0, "Number of documents for kmeans clustering")
	kFlag := flag.Int("k", 0, "Number of clusters")

	flag.Parse()

	logger := util.GetLogger()
	logger.Info("Bootstrap", "Starting indexer", nil)

	indexer := util.New(*bFlag, *dFlag, logger)

	// calculate clusters before storing new documents
	if *cFlag {
		fmt.Println("Building index")
		if *nFlag == 0 || *kFlag == 0 {
			fmt.Println("Both -n and -k flags must be provided")
			flag.Usage()
			os.Exit(1)
		}
		indexer.BuildIndex(*nFlag, *kFlag)
	}

	fmt.Println("Syncing index with DB")
	indexingGoRoutine(indexer)
}
