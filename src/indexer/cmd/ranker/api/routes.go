package main

import (
	"net/http"
)
import "github.com/julienschmidt/httprouter"

func (app *application) routes() http.Handler {
	router := httprouter.New()

	// Routes
	router.HandlerFunc(http.MethodGet, "/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/search", app.searchHandler)

	return app.recoverPanic(router)
}
