package buildfulldocs

import (
	"indexer/src/util/core"
	"indexer/src/util/storage"
	"os"
)

type Passage struct {
	PassageId  string
	Embeddings []core.Vector
}

type FullDocument struct {
	DocId    string
	Passages []Passage
}

func NewDocumentFromId(docId string) *FullDocument {
	passages := getDocumentPassages(docId)

	return &FullDocument{
		DocId:    docId,
		Passages: passages,
	}
}

func NewDocumentsFromIds(docIds []string) []*FullDocument {
	docsChan := make(chan *FullDocument)
	docs := make([]*FullDocument, len(docIds))

	for i, docId := range docIds {
		go func(i int, docId string) {
			docsChan <- NewDocumentFromId(docId)
		}(i, docId)
	}

	for i := 0; i < len(docIds); i++ {
		docs[i] = <-docsChan
	}

	return docs
}

func getDocumentPassages(docId string) []Passage {
	constants := core.GetConstants()
	docPath := constants.StorageDocsDir + "/" + docId
	files, err := os.ReadDir(docPath)
	core.ErrPanic(err)
	passages := make([]Passage, len(files))
	for i, file := range files {
		var passageEmbeddings []core.Vector
		var passageId string = file.Name()
		storagePath := docPath + "/" + passageId
		storage.ReadStructFromFile(storagePath, &passageEmbeddings)
		passages[i] = Passage{
			PassageId:  passageId,
			Embeddings: passageEmbeddings,
		}
	}
	return passages
}
