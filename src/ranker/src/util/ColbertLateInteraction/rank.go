package ColbertLateInteraction

import (
	util "ranker/src/util/core"
	"sort"
)

func Rank(documents []util.Document, query []util.Vector) []RankResultItem {
	var result []RankResultItem
	for _, document := range documents {
		var documentScore float64 = 0.0
		var bestPassageId string = ""
		var bestPassageScore float64 = 0.0

		for _, passage := range document.Passages {
			score := GetScore(passage.Embeddings, query)
			documentScore += score
			if score > bestPassageScore {
				bestPassageScore = score
				bestPassageId = passage.PassageId
			}
		}

		documentScore = documentScore / float64(len(document.Passages))

		item := RankResultItem{
			documentId:    document.DocId,
			documentScore: documentScore,
			bestPassageId: bestPassageId,
		}
		result = append(result, item)
	}

	sort.Sort(sort.Reverse(ByDocumentScore(result)))
	return result
}
