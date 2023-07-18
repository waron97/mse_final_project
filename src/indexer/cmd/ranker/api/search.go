package main

import (
	"fmt"
	"net/http"
)

func (app *application) searchHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Query string `json:"query"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		fmt.Println("ERROR", err)
		app.badRequestResponse(w, r, err)
	}

	// ToDo - get query embedding

	// ToDo - Retrieve top k documents from cluster

	resp := envelope{"documents": make([]string, 0), "query": input.Query}
	err = app.writeJSON(w, http.StatusOK, resp, make(http.Header))
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	fmt.Println(input)
	fmt.Println(input.Query)
}
