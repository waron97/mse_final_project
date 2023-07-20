package ColbertLateInteraction

import (
	"fmt"
	util "ranker/src/util/core"
	"sort"
	"time"
)

func Rank(documents []util.Document, query []util.Vector) []RankResultItem {
	// tested various concurrency approaches, this one is the fastest
	return RankConcurrentDocuments(documents, query)
}

// Execution time: 150-200ms
func RankNoConcurrency(documents []util.Document, query []util.Vector) []RankResultItem {
	start := time.Now()
	var result []RankResultItem
	for _, document := range documents {
		var documentScore float64 = 0.0
		var bestPassageId string = ""
		var bestPassageScore float64 = 0.0

		for _, passage := range document.Passages {
			score := GetScore(passage.Embeddings, query)
			documentScore += score
			if score > bestPassageScore {
				bestPassageScore = score
				bestPassageId = passage.PassageId
			}
		}

		documentScore = documentScore / float64(len(document.Passages))

		item := RankResultItem{
			documentId:    document.DocId,
			documentScore: documentScore,
			bestPassageId: bestPassageId,
		}
		result = append(result, item)
	}

	sort.Sort(sort.Reverse(ByDocumentScore(result)))
	elapsed := time.Since(start)
	fmt.Printf("\nExection took %s\n", elapsed)
	return result
}

// ------------------------------------------------------------

func processDocumentNoConcurrency(document util.Document, query []util.Vector, channel chan RankResultItem) {
	var bestPassageId string = ""
	var bestPassageScore float64 = 0.0

	fullDocument := []util.Vector{}
	for _, passage := range document.Passages {
		fullDocument = append(fullDocument, passage.Embeddings...)
	}
	documentScore := GetScore(fullDocument, query)

	for _, passage := range document.Passages {
		score := GetScore(passage.Embeddings, query)
		if score > bestPassageScore {
			bestPassageScore = score
			bestPassageId = passage.PassageId
		}
	}

	item := RankResultItem{
		documentId:    document.DocId,
		documentScore: documentScore,
		bestPassageId: bestPassageId,
	}

	channel <- item
}

// Execution time: around 50ms
func RankConcurrentDocuments(documents []util.Document, query []util.Vector) []RankResultItem {
	start := time.Now()

	var result []RankResultItem
	channel := make(chan RankResultItem)

	for _, document := range documents {
		go processDocumentNoConcurrency(document, query, channel)
	}

	for i := 0; i < len(documents); i++ {
		result = append(result, <-channel)
	}

	sort.Sort(sort.Reverse(ByDocumentScore(result)))
	elapsed := time.Since(start)
	fmt.Printf("\nExection took %s\n", elapsed)
	return result
}

// ------------------------------------------------------------

type PassageResultItem struct {
	PassageId string
	Score     float64
}

func processPassage(passage util.Passage, query []util.Vector, channel chan PassageResultItem) {
	score := GetScore(passage.Embeddings, query)
	item := PassageResultItem{
		PassageId: passage.PassageId,
		Score:     score,
	}
	channel <- item
}

func processDocumentConcurrent(document util.Document, query []util.Vector, channel chan RankResultItem) {
	var documentScore float64 = 0.0
	var bestPassageId string = ""
	var bestPassageScore float64 = 0.0

	passageChannel := make(chan PassageResultItem)

	for _, passage := range document.Passages {
		go processPassage(passage, query, passageChannel)
	}

	for i := 0; i < len(document.Passages); i++ {
		passageResult := <-passageChannel
		documentScore += passageResult.Score
		if passageResult.Score > bestPassageScore {
			bestPassageScore = passageResult.Score
			bestPassageId = passageResult.PassageId
		}
	}

	documentScore = documentScore / float64(len(document.Passages))

	item := RankResultItem{
		documentId:    document.DocId,
		documentScore: documentScore,
		bestPassageId: bestPassageId,
	}

	channel <- item
}

// Execution time: around 60ms
func RankConcurrentDocumentsAndPassages(documents []util.Document, query []util.Vector) []RankResultItem {
	start := time.Now()

	var result []RankResultItem
	channel := make(chan RankResultItem)

	for _, document := range documents {
		go processDocumentConcurrent(document, query, channel)
	}

	for i := 0; i < len(documents); i++ {
		result = append(result, <-channel)
	}

	sort.Sort(sort.Reverse(ByDocumentScore(result)))
	elapsed := time.Since(start)
	fmt.Printf("\nExection took %s\n", elapsed)
	return result
}
