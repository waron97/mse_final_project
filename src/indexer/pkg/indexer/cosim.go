package indexer

import "math"

// Cosim Cosine Similarity between two vectors v1, v2
func cosim(v1 Vector, v2 Vector) float64 {
	dotProduct := dot(v1, v2)
	magnitudeV1 := magnitude(v1)
	magnitudeV2 := magnitude(v2)

	// Division by 0
	if magnitudeV1 == 0 || magnitudeV2 == 0 {
		return 0.0
	}

	// Calculate cosine similarity
	return dotProduct / (magnitudeV1 * magnitudeV2)
}

// dot calculates the dot product of two vectors.
func dot(v1 Vector, v2 Vector) float64 {
	// same length
	if len(v1) != len(v2) {
		return 0.0
	}

	dotProduct := 0.0
	for i := 0; i < len(v1); i++ {
		dotProduct += v1[i] * v2[i]
	}

	return dotProduct
}

// magnitude calculates the magnitude of a vector.
func magnitude(v Vector) float64 {
	sumOfSquares := 0.0
	for _, val := range v {
		sumOfSquares += val * val
	}

	return math.Sqrt(sumOfSquares)
}
