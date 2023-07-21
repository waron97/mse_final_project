package core

import (
	"fmt"
	"os"
	"ranker/src/util/storage"
)

type Passage struct {
	PassageId  string
	Embeddings []Vector
}

type Document struct {
	DocId    string
	Passages []*Passage
}

func (d Document) String() string {
	return fmt.Sprintf("{Document %s, %d passages}", d.DocId, len(d.Passages))
}

func NewDocumentFromId(docId string) *Document {
	return &Document{
		DocId:    docId,
		Passages: getDocumentPassages(docId),
	}
}

func NewDocumentsFromIds(docIds []string) []*Document {
	docsChan := make(chan *Document)
	docs := make([]*Document, len(docIds))

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

func getDocumentPassages(docId string) []*Passage {
	constants := GetConstants()
	docPath := constants.StorageDocsDir + "/" + docId
	files, err := os.ReadDir(docPath)
	ErrPanic(err)
	passages := make([]*Passage, len(files))
	for i, file := range files {
		var passageEmbeddings []Vector
		var passageId string = file.Name()
		storagePath := docPath + "/" + passageId
		storage.ReadStructFromFile(storagePath, &passageEmbeddings)
		passages[i] = &Passage{
			PassageId:  passageId,
			Embeddings: passageEmbeddings,
		}
	}
	return passages
}
