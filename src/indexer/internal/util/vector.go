package util

import (
	"fmt"
	"math/rand"
)

type Vector []float64

func (v Vector) ToString() string {
	result := ""
	for _, e := range v {
		result += fmt.Sprintf("%f|", e)
	}
	return result
}

// ToDo - Delete once BERT embeddings work
func generateRandomVector(length int) Vector {
	slice := make(Vector, length)
	for i := 0; i < length; i++ {
		slice[i] = rand.Float64()
	}
	return slice
}

func getEmbeddings(text string) []Vector {
	// ToDo - replace with BERT embeddings
	vectorLength := 768
	vectorAmount := 512

	results := make([]Vector, vectorAmount)
	for i := 0; i < len(results); i++ {
		results[i] = generateRandomVector(vectorLength)
	}

	return results
}
