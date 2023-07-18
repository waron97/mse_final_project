package ColbertLateInteraction

import (
	"fmt"
	util "ranker/src/util/core"
)

func getInteractionMatrix(doc1 []util.Vector, doc2 []util.Vector) [][]float64 {
	if len(doc1) != len(doc2) {
		panic(fmt.Sprintf("ColbertLateInteraction expected vectors of equal length, got %d and %d", len(doc1), len(doc2)))
	}
	var result [][]float64
	d := len(doc1)
	for i := 0; i < d; i++ {
		var row []float64
		for j := 0; j < d; j++ {
			row = append(row, doc1[i].CosSim(doc2[j]))
		}
		result = append(result, row)
	}
	return result
}
