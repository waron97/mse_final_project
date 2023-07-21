package networking

import (
	"ranker/src/util/ColbertLateInteraction"
	"ranker/src/util/db"
)

func MapRankingResults(result []ColbertLateInteraction.RankResultItem) []ResultItem {
	mapped := make([]ResultItem, 0)
	for _, item := range result {
		page := db.GetPage(item.DocumentId)
		mapped = append(mapped, ResultItem{
			DocumentId:    item.DocumentId,
			DocumentScore: item.DocumentScore,
			BestPassageId: item.BestPassageId,

			DocumentTitle:       page.Title,
			DocumentDescription: page.Description,
			DocumentUrl:         page.URL,
			BestPassageText:     page.GetPassage(item.BestPassageId),
		})
	}
	return mapped
}
