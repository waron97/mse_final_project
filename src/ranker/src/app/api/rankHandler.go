package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ranker/src/util/ColbertLateInteraction"
	"ranker/src/util/bert"
	"ranker/src/util/cache"
	"ranker/src/util/core"
	"ranker/src/util/networking"
	"ranker/src/util/prerank"
	"time"
)

type RankMeta struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
}

type RankResponse struct {
	Data []networking.ResultItem `json:"data"`
	Meta RankMeta                `json:"meta"`
}

var clusters []prerank.Cluster = prerank.ReadClusters()
var responseCache *cache.Cache = cache.NewCache()

func RankHandler(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	skip, limit, page := networking.Pagination(r)
	query := r.URL.Query().Get("query")
	if query == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Query is empty"))
		return
	}
	var ranking []ColbertLateInteraction.RankResultItem
	if cachedItem, ok := responseCache.Get(query); ok {
		ranking = cachedItem.([]ColbertLateInteraction.RankResultItem)
		core.MeasureTime(now, "Cache")
	} else {
		queryEmbeddings := bert.GetEmbeddings(query)
		core.MeasureTime(now, "BERT")

		n := 200
		topk := prerank.Prerank(queryEmbeddings, clusters, n)
		topk = topk[:n]

		core.MeasureTime(now, fmt.Sprintf("Preranked documents %d", len(topk)))

		// Load documents from disk, at most 4 at a time
		// to prevent running out of RAM
		topkDocs := make(chan *core.Document, 50)
		go core.NewDocumentsFromIdsChan(topk, topkDocs)
		core.MeasureTime(now, "NewDocumentsFromIds started")

		// Process documents as they come in
		ranking = ColbertLateInteraction.RankChan(topkDocs, len(topk), queryEmbeddings)
		core.MeasureTime(now, "ColbertLateInteraction")

		responseCache.Set(query, ranking)
	}
	numDocs := len(ranking)

	ranking = ranking[skip : skip+limit]
	mapped := networking.MapRankingResults(ranking)
	core.MeasureTime(now, "MapRankingResults")
	payload := RankResponse{
		Data: mapped,
		Meta: RankMeta{
			Page:  page,
			Limit: limit,
			Total: numDocs,
		}}
	w.Header().Set("Content-Type", "application/json")
	body, err := json.MarshalIndent(payload, "", "  ")
	core.ErrPanic(err)
	w.Write(body)
}

func RankAllHandler(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	query := r.URL.Query().Get("query")
	if query == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Query is empty"))
		return
	}
	var ranking []ColbertLateInteraction.RankResultItem
	if cachedItem, ok := responseCache.Get(query); ok {
		ranking = cachedItem.([]ColbertLateInteraction.RankResultItem)
		core.MeasureTime(now, "Cache")
	} else {
		queryEmbeddings := bert.GetEmbeddings(query)
		core.MeasureTime(now, "BERT")

		n := 200
		topk := prerank.Prerank(queryEmbeddings, clusters, n)
		topk = topk[:n]

		core.MeasureTime(now, fmt.Sprintf("Preranked documents %d", len(topk)))

		// Load documents from disk, at most 4 at a time
		// to prevent running out of RAM
		topkDocs := make(chan *core.Document, 50)
		go core.NewDocumentsFromIdsChan(topk, topkDocs)
		core.MeasureTime(now, "NewDocumentsFromIds started")

		// Process documents as they come in
		ranking = ColbertLateInteraction.RankChan(topkDocs, len(topk), queryEmbeddings)
		core.MeasureTime(now, "ColbertLateInteraction")

		responseCache.Set(query, ranking)
	}
	numDocs := len(ranking)

	mapped := networking.MapRankingResults(ranking)
	core.MeasureTime(now, "MapRankingResults")
	payload := RankResponse{
		Data: mapped,
		Meta: RankMeta{
			Total: numDocs,
		}}
	w.Header().Set("Content-Type", "application/json")
	body, err := json.MarshalIndent(payload, "", "  ")
	core.ErrPanic(err)
	w.Write(body)
}
