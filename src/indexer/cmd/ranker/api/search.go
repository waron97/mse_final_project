package main

import (
	"fmt"
	"indexer/internal/util"
	"net/http"
	"strconv"
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
	queryEmb := app.store.ComputeEmbedding(input.Query)
	queryAvgEmb, err := app.store.ComputeAvgEmbedding(queryEmb)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	//fmt.Println(queryAvgEmb)

	// ToDo - Retrieve top k documents from cluster
	centroidId := app.store.ClusterMap.GetClosetCentroid(queryAvgEmb)
	var centroidEmb []*util.DocEmbedding
	err = util.ReadStructFromFile(app.store.ClusterMap.BaseDir+"/centroid_"+strconv.Itoa(centroidId), &centroidEmb)
	if err != nil {
		fmt.Println(err)
		app.serverErrorResponse(w, r, err)
		return
	}
	fmt.Println(centroidEmb)

	resp := envelope{"documents": make([]string, 0), "query": input.Query}
	err = app.writeJSON(w, http.StatusOK, resp, make(http.Header))
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	fmt.Println(input)
	fmt.Println(input.Query)
}
