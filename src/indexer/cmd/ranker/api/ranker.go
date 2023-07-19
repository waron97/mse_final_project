package main

import (
	"flag"
	"fmt"
	"indexer/pkg/indexer"
	"net/http"
	"os"
	"time"
)

type config struct {
	port int
	db   string
	dir  string
}

type application struct {
	config config
	logger *indexer.Logger
	index  *indexer.Index
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.db, "db", "mongodb://localhost:27017", "DB connection string")
	flag.StringVar(&cfg.dir, "b", "./offline-indexer/", "Base directory for indexer storage")
	flag.Parse()

	logger := indexer.GetLogger()
	index := indexer.NewStore(cfg.dir)
	app := &application{
		config: cfg,
		logger: logger,
		index:  index,
	}

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("CD", pwd)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Info("ranker", "starting server", map[string]string{
		"addr": srv.Addr,
	})

	_ = srv.ListenAndServe()
}
