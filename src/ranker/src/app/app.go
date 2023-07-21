package app

import (
	"fmt"
	"net/http"
	"ranker/src/app/api"
	util "ranker/src/util/core"
)

func RunTasks() {
	util.GetLogger().Info("Bootstrap", "Hello from GO ranker!", nil)
	handler := http.NewServeMux()
	handler.HandleFunc("/rank", api.RankHandler)
	handler.HandleFunc("/health", api.HealthHandler)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", 5000),
		Handler: handler,
	}
	server.ListenAndServe()
}
