package ColbertLateInteraction

import "fmt"

type RankResultItem struct {
	DocumentId    string  `json:"documentId"`
	DocumentScore float64 `json:"documentScore"`
	BestPassageId string  `json:"bestPassageId"`
}

func (r RankResultItem) String() string {
	return fmt.Sprintf(
		"{RankResultItem [documentId=%v, documentScore=%v, bestPassageId=%v]}",
		r.DocumentId, r.DocumentScore, r.BestPassageId)
}

type ByDocumentScore []RankResultItem

func (s ByDocumentScore) Len() int {
	return len(s)
}

func (s ByDocumentScore) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByDocumentScore) Less(i, j int) bool {
	return s[i].DocumentScore < s[j].DocumentScore
}
