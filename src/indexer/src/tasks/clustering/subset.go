package clustering

import (
	"math/rand"
	"time"
)

func getDocumentSubset(files []*Document, count int) []*Document {
	numFiles := len(files)

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	random.Shuffle(numFiles, func(i, j int) {
		files[i], files[j] = files[j], files[i]
	})

	return files[:count]
}
