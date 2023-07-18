package main

import "net/http"

func (app *application) logError(r *http.Request, err error, message string) {
	app.logger.Error(r.Method, "error", err)
}

// return a HTML 500 error response with structured JSON response containing error message
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	env := envelope{"error": message}

	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logError(r, err, "error")
		w.WriteHeader(500)
	}
}

// log error and return error response
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	message := "server encountered an error"
	app.logError(r, err, message)
	app.errorResponse(w, r, http.StatusNotFound, message)
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
}
