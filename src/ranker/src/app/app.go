package app

import (
	"fmt"
	"net/http"
	"ranker/src/app/api"
	"ranker/src/util/core"

	"github.com/rs/cors"
)

func RunTasks() {
	doc := core.NewDocumentFromId("64b7ef444db6e0d1c693764e")
	fmt.Println("DocID", doc.DocId)

	handler := http.NewServeMux()
	handler.HandleFunc("/rank", api.RankHandler)
	handler.HandleFunc("/health", api.HealthHandler)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", 5000),
		Handler: cors.Default().Handler(handler),
	}
	server.ListenAndServe()
}
