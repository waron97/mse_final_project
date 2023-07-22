package core

import (
	"fmt"
	"ranker/src/util/storage"
	"time"
)

type Passage struct {
	PassageId  string
	Embeddings []Vector
}

type Document struct {
	DocId              string
	Passages           []Passage
	DocumentEmbeddings []Vector
}

func (d Document) String() string {
	return fmt.Sprintf("{Document %s, %d passages}", d.DocId, len(d.Passages))
}

func NewDocumentFromId(docId string) *Document {
	t := time.Now()
	constants := GetConstants()
	docPath := constants.StorageFullDocsDir + "/" + docId
	var doc Document
	storage.ReadStructFromFile(docPath, &doc)
	elapsed := time.Since(t)
	fmt.Println("Document loaded in", elapsed)
	return &doc
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
