package api

import (
	"encoding/json"
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
		topk := prerank.Prerank(queryEmbeddings, clusters, 100)
		core.MeasureTime(now, "Prerank")
		topkDocs := core.NewDocumentsFromIds(topk)
		core.MeasureTime(now, "NewDocumentsFromIds")
		ranking = ColbertLateInteraction.Rank(topkDocs, queryEmbeddings)
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
