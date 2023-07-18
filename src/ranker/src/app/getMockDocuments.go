package app

import util "ranker/src/util/core"

func getMockDocuments() []util.Document {
	return []util.Document{
		{
			DocId: "doc1",
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
			},
		},
		{
			DocId: "doc2",
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
		},
		{
			DocId: "doc3",
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
		},
		{
			DocId: "doc4",
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
		},
		{
			DocId: "doc5",
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
		},
	}
}
