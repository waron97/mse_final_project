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
	DocId    string
	Passages []Passage
}

func (d Document) String() string {
	return fmt.Sprintf("{Document %s, %d passages}", d.DocId, len(d.Passages))
}

func NewDocumentFromId(docId string) *Document {
	constants := GetConstants()
	docPath := constants.StorageFullDocsDir + "/" + docId
	var doc Document
	storage.ReadStructFromFile(docPath, &doc)
	return &doc
}

func startDocumentGetter(tasks chan string, docs chan *Document) {
	for {
		task, ok := <-tasks
		if !ok {
			return
		}
		doc := NewDocumentFromId(task)
		docs <- doc
	}
}

func NewDocumentsFromIds(docIds []string) []*Document {
	tasksChan := make(chan string, len(docIds))
	docsChan := make(chan *Document)
	docs := make([]*Document, len(docIds))

	for _, docId := range docIds {
		tasksChan <- docId
	}

	close(tasksChan)

	go startDocumentGetter(tasksChan, docsChan)
	go startDocumentGetter(tasksChan, docsChan)
	go startDocumentGetter(tasksChan, docsChan)
	go startDocumentGetter(tasksChan, docsChan)

	for i := 0; i < len(docIds); i++ {
		docs[i] = <-docsChan
	}

	return docs
}

func startDocumentsGetterChan(tasks chan string, docs chan *Document) {
	for {
		task, ok := <-tasks
		if !ok {
			return
		}
		n := time.Now()
		doc := NewDocumentFromId(task)
		fmt.Println("Doc loaded in", time.Since(n))
		docs <- doc
	}
}

func NewDocumentsFromIdsChan(docIds []string, out chan *Document) {
	tasksChan := make(chan string, len(docIds))

	for _, docId := range docIds {
		tasksChan <- docId
	}

	close(tasksChan)

	numWorkers := 20
	for i := 0; i < numWorkers; i++ {
		go startDocumentsGetterChan(tasksChan, out)
	}
}
