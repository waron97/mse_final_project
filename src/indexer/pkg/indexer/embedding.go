package indexer

import (
	"errors"
	"math/rand"
)

type Vector []float64

// ToDo - replace with BERT embeddings
func generateRandomVector(length int) Vector {
	slice := make(Vector, length)
	for i := 0; i < length; i++ {
		slice[i] = rand.Float64()
	}
	return slice
}

// ToDo - replace with BERT embeddings
func getEmbeddings(text string) []Vector {

	vectorLength := 768
	vectorAmount := 512

	results := make([]Vector, vectorAmount)
	for i := 0; i < len(results); i++ {
		results[i] = generateRandomVector(vectorLength)
	}

	return results
}

func GetEmbedding(text string) []Vector {
	return getEmbeddings(text)
}

func GetAvgEmbedding(emb []Vector) (Vector, error) {
	rows := len(emb)
	if rows == 0 {
		return nil, errors.New("embeddings empty")
	}

	columns := len(emb[0])
	averages := make(Vector, columns)

	for col := 0; col < columns; col++ {
		sum := 0.0
		for row := 0; row < rows; row++ {
			sum += emb[row][col]
		}
		averages[col] = sum / float64(rows)
	}
	return averages, nil
}
