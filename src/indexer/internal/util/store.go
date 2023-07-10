package util

import (
	"encoding/gob"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

type Store struct {
	logger          *Logger
	storePath       string
	infoPath        string
	documentPath    string
	avgDocumentPath string
	Timestamp       time.Time
}

func New(indexPath string, logger *Logger) *Store {
	timestamp, err := readTimestamp(indexPath + "/info.index")
	if err != nil {
		fmt.Println(err)
	}
	return &Store{
		logger:          logger,
		storePath:       indexPath,
		Timestamp:       timestamp,
		documentPath:    indexPath + "/documents",
		avgDocumentPath: indexPath + "/documents_averaged",
		infoPath:        indexPath + "/info.index",
	}
}

func (index *Store) Store() {
	documents, err := GetAllCrawlPages(index.Timestamp)
	newTime := time.Now()
	if err != nil {
		index.logger.Critical("store.go;Store()", "cannot retrieve data from DB", err)
		fmt.Println(err)
	}
	fmt.Println(documents)

	for _, doc := range documents {
		emb := getEmbeddings(doc.MainText)
		err = writeStructToFile(index.documentPath+"/"+doc.ID.Hex()+".gob", emb)
		if err != nil {
			index.logger.Critical("store.go;Store()", "cannot write vector to file", err)
		}

		avgEmb, err := averageDocument(emb)
		if err != nil {
			index.logger.Critical("store.go;Store()", "vector 0 length", err)
			panic("vector 0 length")
		}

		err = writeStructToFile(index.avgDocumentPath+"/"+doc.ID.Hex()+".gob", avgEmb)
		if err != nil {
			index.logger.Critical("store.go;Store()", "cannot write vector to file", err)
			panic("can't write to file")
		}
	}
	index.updateTimestamp(newTime)
}

func (index *Store) updateTimestamp(newTime time.Time) {
	err := writeTimestamp(newTime, index.infoPath)
	if err != nil {
		panic("can't write file")
	}
	index.Timestamp = newTime
}

func readTimestamp(filepath string) (time.Time, error) {
	d, err := os.ReadFile(filepath)
	if err != nil {
		return time.Time{}, err
	}
	str := string(d)
	timestamp, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return time.Time{}, err
	}
	return timestamp, nil
}

func writeTimestamp(timestamp time.Time, filepath string) error {
	d := []byte(timestamp.Format(time.RFC3339))
	err := os.WriteFile(filepath, d, 0644)

	if err != nil {
		return nil
	}
	return nil
}

func averageDocument(embeddings []Vector) (Vector, error) {
	rows := len(embeddings)
	if rows == 0 {
		return nil, errors.New("embeddings empty")
	}

	columns := len(embeddings[0])
	averages := make(Vector, columns)

	for col := 0; col < columns; col++ {
		sum := 0.0
		for row := 0; row < rows; row++ {
			sum += embeddings[row][col]
		}
		averages[col] = sum / float64(rows)
	}
	return averages, nil
}

func writeStructToFile(filename string, data interface{}) error {
	err := os.MkdirAll(filepath.Dir(filename), 0755)
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return err
	}
	return nil
}

// BuildIndex Build a new index by calculating k centroids using n random documents
func (index *Store) BuildIndex(n, k int) {
	files, err := ioutil.ReadDir(index.avgDocumentPath)
	if err != nil {
		panic("can't read dir")
	}

	sampledFiles, _ := getDocumentSubset(files, n)

	// Read embeddings from files
	sampledDocEmbeddings := make([]*DocEmbedding, len(sampledFiles))
	for i, file := range sampledFiles {
		filePath := filepath.Join(index.avgDocumentPath, file.Name())
		fmt.Println(filePath)
		embedding, err := readAvgDocument(filePath)
		if err != nil {
			log.Printf("Failed to parse %s: %s", file.Name(), err)
			panic("can't read file")
		}
		sampledDocEmbeddings[i] = NewDocEmbedding(file.Name(), embedding)
	}
	Cluster(sampledDocEmbeddings, k)
}

func getDocumentSubset(files []os.FileInfo, count int) ([]os.FileInfo, []os.FileInfo) {
	numFiles := len(files)
	if numFiles <= count {
		return nil, files
	}

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	random.Shuffle(numFiles, func(i, j int) {
		files[i], files[j] = files[j], files[i]
	})

	return files[:count], files[count:]
}

func readAvgDocument(filePath string) (Vector, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	var data Vector
	err = decoder.Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
