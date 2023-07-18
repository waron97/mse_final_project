package ColbertLateInteraction

import (
	util "ranker/src/util/core"
)

func GetScore(doc1 []util.Vector, doc2 []util.Vector) float64 {
	mtx := getInteractionMatrix(doc1, doc2)
	var score float64 = 0.0
	for _, row := range mtx {
		max_sim := getMax(row)
		score += max_sim
	}
	return score
}
