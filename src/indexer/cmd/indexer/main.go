package main

import (
	"flag"
	"fmt"
	"indexer/pkg/indexer"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type config struct {
	path string
	k    int
	n    int
	c    bool
}

func indexingGoRoutine(i *indexer.Index) {
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
	runIndexStore(&wg, i)
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
				runIndexStore(&wg, i)
				fmt.Println("indexing", time.Now())
			}
		}
	}()
	select {}
}

// runIndexStore runs the indexer.Store() function and marks it as completed when finished
func runIndexStore(wg *sync.WaitGroup, i *indexer.Index) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		indexer.StoreAllCrawlPages(i)
	}()
}

func main() {
	var cfg config
	flag.StringVar(&cfg.path, "path", "./testIndex", "Index Path")
	flag.IntVar(&cfg.k, "k", 16, "Cluster amount for kmeans")
	flag.IntVar(&cfg.n, "n", 5000, "document sample size for kmeans")
	flag.BoolVar(&cfg.c, "c", false, "(re)build index cluster")

	flag.Parse()
	store := indexer.NewStore(cfg.path)
	if cfg.c {
		fmt.Println("TRUE", cfg.c)
		store.BuildCluster(cfg.n, cfg.k)
	}
	indexingGoRoutine(store)
}
