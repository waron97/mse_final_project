package clustering

import (
	"indexer/src/util/core"
	"indexer/src/util/storage"
	"os"
	"path/filepath"
)

type Document struct {
	Embedding core.Vector
	DocID     string
}

func loadDocuments() []Document {
	constants := core.GetConstants()
	files, err := os.ReadDir(constants.StorageAverageDocsDir)
	core.ErrPanic(err)

	documents := make([]Document, len(files))
	for j, file := range files {
		path := filepath.Join(constants.StorageAverageDocsDir, file.Name())

		var emb core.Vector
		err = storage.ReadStructFromFile(path, &emb)
		core.ErrPanic(err)

		documents[j] = Document{
			Embedding: emb,
			DocID:     file.Name(),
		}
	}
	return documents
}

func getDocumentEmbeddings(documents []Document) []core.Vector {
	embeddings := make([]core.Vector, len(documents))
	for j, doc := range documents {
		embeddings[j] = doc.Embedding
	}
	return embeddings
}
