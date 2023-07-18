package main

import (
	"flag"
	"fmt"
	"indexer/internal/util"
	"net/http"
	"time"
)

type config struct {
	port int
}

type application struct {
	config config
	logger *util.Logger
	// ToDo - add indexer
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.Parse()

	logger := util.GetLogger()
	app := &application{
		config: cfg,
		logger: logger,
	}

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
