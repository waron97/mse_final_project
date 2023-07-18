package ColbertLateInteraction

import "fmt"

type RankResultItem struct {
	documentId    string
	documentScore float64
	bestPassageId string
}

func (r RankResultItem) String() string {
	return fmt.Sprintf(
		"{RankResultItem [documentId=%v, documentScore=%v, bestPassageId=%v]}",
		r.documentId, r.documentScore, r.bestPassageId)
}

type ByDocumentScore []RankResultItem

func (s ByDocumentScore) Len() int {
	return len(s)
}

func (s ByDocumentScore) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByDocumentScore) Less(i, j int) bool {
	return s[i].documentScore < s[j].documentScore
}
