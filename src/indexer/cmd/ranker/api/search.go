package main

import (
	"fmt"
	"indexer/pkg/indexer"
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
		return
	}

	// get averaged query embedding
	queryEmb := indexer.GetEmbedding(input.Query)
	queryAvgEmb, err := indexer.GetAvgEmbedding(queryEmb)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	centroid, err := app.index.GetClosestCentroidEmb(queryAvgEmb)
	if err != nil {
		fmt.Println(err)
		app.serverErrorResponse(w, r, err)
		return
	}

	// ToDo - `centroid` is a []*DocEmbedding, consisting of `DocId` (id from MongoDB)
	// ToDo - and `Embedding`, the averaged embedding of the document

	for _, doc := range centroid {
		var embeddings []indexer.Vector
		err = indexer.ReadStructFromFile(app.index.GetDocIdPath(doc.DocId), &embeddings)
		if err != nil {
			fmt.Println(err)
			app.serverErrorResponse(w, r, err)
			return
		}
		fmt.Println(len(embeddings), len(embeddings[0]))
		break
	}

	resp := envelope{"documents": make([]string, 0), "query": input.Query}
	err = app.writeJSON(w, http.StatusOK, resp, make(http.Header))
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	fmt.Println(input)
	fmt.Println(input.Query)
}
