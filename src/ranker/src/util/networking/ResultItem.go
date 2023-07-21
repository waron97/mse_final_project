package networking

type ResultItem struct {
	DocumentId    string  `json:"documentId"`
	DocumentScore float64 `json:"documentScore"`
	BestPassageId string  `json:"bestPassageId"`

	DocumentTitle   string `json:"documentTitle"`
	BestPassageText string `json:"bestPassageText"`
	DocumentUrl     string `json:"documentUrl"`
}
