package bert

import (
	"errors"
	"indexer/src/util/core"
	"math/rand"
)

// ToDo - replace with BERT embeddings
func generateRandomVector(length int) core.Vector {
	slice := make(core.Vector, length)
	for i := 0; i < length; i++ {
		slice[i] = rand.Float64()
	}
	return slice
}

// ToDo - replace with BERT embeddings
func getEmbeddings(text string) []core.Vector {

	vectorLength := 100
	vectorAmount := 60

	results := make([]core.Vector, vectorAmount)
	for i := 0; i < len(results); i++ {
		results[i] = generateRandomVector(vectorLength)
	}

	return results
}

func GetEmbedding(text string) []core.Vector {
	return getEmbeddings(text)
}

func GetAvgEmbedding(emb []core.Vector) (core.Vector, error) {
	rows := len(emb)
	if rows == 0 {
		return nil, errors.New("embeddings empty")
	}

	columns := len(emb[0])
	averages := make(core.Vector, columns)

	for col := 0; col < columns; col++ {
		sum := 0.0
		for row := 0; row < rows; row++ {
			sum += emb[row][col]
		}
		averages[col] = sum / float64(rows)
	}
	return averages, nil
}
