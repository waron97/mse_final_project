package clustering

func centroidContains(data []*DocEmbedding, docId string) bool {
	for _, doc := range data {
		if doc.DocId == docId {
			return true
		}
	}
	return false
}
