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
	Passages []Passage
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
	docs := make([]*Document, 0)
	for _, docId := range docIds {
		docs = append(docs, NewDocumentFromId(docId))
	}
	return docs
}

func getDocumentPassages(docId string) []Passage {
	constants := GetConstants()
	docPath := constants.StorageDocsDir + "/" + docId
	files, err := os.ReadDir(docPath)
	ErrPanic(err)
	passages := make([]Passage, 0)
	for _, file := range files {
		var passageEmbeddings []Vector
		var passageId string = file.Name()
		storagePath := docPath + "/" + passageId
		storage.ReadStructFromFile(storagePath, &passageEmbeddings)
		passages = append(passages, Passage{
			PassageId:  passageId,
			Embeddings: passageEmbeddings,
		})
	}
	return passages
}
