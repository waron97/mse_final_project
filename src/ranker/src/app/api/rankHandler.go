package api

import (
	"encoding/json"
	"net/http"
	"ranker/src/util/ColbertLateInteraction"
	"ranker/src/util/bert"
	"ranker/src/util/core"
	"ranker/src/util/networking"
	"ranker/src/util/prerank"
)

var clusters []prerank.Cluster = prerank.ReadClusters()

func RankHandler(w http.ResponseWriter, r *http.Request) {
	skip, limit := networking.Pagination(r)
	query := r.URL.Query().Get("query")
	if query == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Query is empty"))
		return
	}
	queryEmbeddings := bert.GetEmbeddings(query)
	topk := prerank.Prerank(queryEmbeddings, clusters, 100)
	topkDocs := core.NewDocumentsFromIds(topk)
	ranking := ColbertLateInteraction.Rank(topkDocs, queryEmbeddings)
	ranking = ranking[skip : skip+limit]
	payload := networking.MapRankingResults(ranking)
	w.Header().Set("Content-Type", "application/json")
	body, err := json.MarshalIndent(payload, "", "  ")
	core.ErrPanic(err)
	w.Write(body)
}
