package ColbertLateInteraction

import (
	util "ranker/src/util/core"
	"sort"
)

func Rank(documents []*util.Document, query []util.Vector) []RankResultItem {
	// tested various concurrency approaches, this one is the fastest

	var result []RankResultItem
	channel := make(chan RankResultItem)

	for _, document := range documents {
		go processDocument(document, query, channel)
	}

	for i := 0; i < len(documents); i++ {
		result = append(result, <-channel)
	}

	sort.Sort(sort.Reverse(ByDocumentScore(result)))
	return result
}

func RankChan(documents chan *util.Document, docsCount int, query []util.Vector) []RankResultItem {
	var result []RankResultItem
	channel := make(chan RankResultItem)

	numWorkers := 10
	for i := 0; i < numWorkers; i++ {
		go startDocumentWorker(documents, query, channel)
	}

	for i := 0; i < docsCount; i++ {
		doc := <-channel
		result = append(result, doc)
	}

	sort.Sort(sort.Reverse(ByDocumentScore(result)))
	return result
}

func startDocumentWorker(documents chan *util.Document, query []util.Vector, channel chan RankResultItem) {
	for {
		document, ok := <-documents
		if !ok {
			return
		}
		processDocument(document, query, channel)
	}
}

func processDocument(document *util.Document, query []util.Vector, channel chan RankResultItem) {
	var bestPassageId string = ""
	var bestPassageScore float64 = 0.0

	fullDocument := []util.Vector{}
	for _, passage := range document.Passages {
		fullDocument = append(fullDocument, passage.Embeddings...)
	}
	documentScore := GetScore(query, fullDocument)

	for _, passage := range document.Passages {
		score := GetScore(query, passage.Embeddings)
		if score > bestPassageScore {
			bestPassageScore = score
			bestPassageId = passage.PassageId
		}
	}

	item := RankResultItem{
		DocumentId:    document.DocId,
		DocumentScore: documentScore,
		BestPassageId: bestPassageId,
	}

	channel <- item
}
