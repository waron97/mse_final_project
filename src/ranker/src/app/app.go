package app

import (
	"fmt"
	"net/http"
	"ranker/src/app/api"

	"github.com/rs/cors"
)

func RunTasks() {
	handler := http.NewServeMux()
	handler.HandleFunc("/rank", api.RankHandler)
	handler.HandleFunc("/health", api.HealthHandler)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", 5000),
		Handler: cors.Default().Handler(handler),
	}
	server.ListenAndServe()
}
