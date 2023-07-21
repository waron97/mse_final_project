package ColbertLateInteraction

import (
	util "ranker/src/util/core"
)

func getInteractionMatrix(doc1 []util.Vector, doc2 []util.Vector) [][]float64 {
	var result [][]float64
	rows := len(doc1)
	cols := len(doc2)
	for i := 0; i < rows; i++ {
		var row []float64
		for j := 0; j < cols; j++ {
			row = append(row, doc1[i].CosSim(doc2[j]))
		}
		result = append(result, row)
	}
	return result
}
