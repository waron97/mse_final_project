package core

import "errors"

func GetEmbeddings(docId string, passageId string) []Vector {
	var dims int = 100
	return []Vector{
		GetRandomVector(dims),
		GetRandomVector(dims),
		GetRandomVector(dims),
		GetRandomVector(dims),
		GetRandomVector(dims),
		GetRandomVector(dims),
		GetRandomVector(dims),
		GetRandomVector(dims),
		GetRandomVector(dims),
		GetRandomVector(dims),
	}
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
