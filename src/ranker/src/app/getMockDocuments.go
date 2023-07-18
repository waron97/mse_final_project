package app

import (
	"fmt"
	util "ranker/src/util/core"
)

func getMockDocuments() []util.Document {
	var result []util.Document
	for i := 0; i < 1000; i++ {
		result = append(result, util.Document{
			DocId: "doc" + fmt.Sprintf("%v", i),
			Passages: []util.Passage{
				{
					PassageId:  "1",
					Embeddings: util.GetEmbeddings("hello", "world"),
				},
				{
					PassageId:  "2",
					Embeddings: util.GetEmbeddings("hello", "world"),
				},
				{
					PassageId:  "3",
					Embeddings: util.GetEmbeddings("hello", "world"),
				},
				{
					PassageId:  "4",
					Embeddings: util.GetEmbeddings("hello", "world"),
				},
				{
					PassageId:  "5",
					Embeddings: util.GetEmbeddings("hello", "world"),
				},
			},
		})
	}
	return result
}
