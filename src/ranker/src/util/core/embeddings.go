package util

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
