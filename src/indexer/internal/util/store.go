package util

import (
	"encoding/gob"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

type Store struct {
	logger          *Logger
	Timestamp       time.Time
	connString      string
	storePath       string
	infoPath        string
	documentPath    string
	avgDocumentPath string
	clusterDir      string
	clusterMapPath  string
	ClusterMap      *ClusterMap
}

func New(indexPath string, connString string, logger *Logger) *Store {
	timestamp, _ := readTimestamp(indexPath + "/info.index")

	var clusterMap *ClusterMap
	err := readStructFromFile(indexPath+"/cluster"+"/clusterMap.index", clusterMap)
	if err != nil {
		clusterMap = NewClusterMap(indexPath + "/cluster")
	}

	return &Store{
		logger:          logger,
		connString:      connString,
		storePath:       indexPath,
		Timestamp:       timestamp,
		documentPath:    indexPath + "/documents",
		avgDocumentPath: indexPath + "/documents_averaged",
		infoPath:        indexPath + "/info.index",
		clusterDir:      indexPath + "/cluster",
		clusterMapPath:  indexPath + "/cluster" + "/clusterMap.index",
		ClusterMap:      clusterMap,
	}
}

func (index *Store) Store() {
	documents, err := GetAllCrawlPages(index.connString, index.Timestamp)
	newTime := time.Now()
	if err != nil {
		index.logger.Critical("store.go;Store()", "cannot retrieve data from DB", err)
		panic(err)
	}
	fmt.Println(documents)

	for _, doc := range documents {
		emb := getEmbeddings(doc.MainText)
		err = writeStructToFile(index.documentPath+"/"+doc.ID.Hex(), emb)
		if err != nil {
			index.logger.Critical("store.go;Store()", "cannot write vector to file", err)
			panic(err)
		}

		avgEmb, err := averageDocument(emb)
		if err != nil {
			index.logger.Critical("store.go;Store()", "vector 0 length", err)
			panic("vector 0 length")
		}

		err = writeStructToFile(index.avgDocumentPath+"/"+doc.ID.Hex(), avgEmb)
		if err != nil {
			index.logger.Critical("store.go;Store()", "cannot write vector to file", err)
			panic(err)
		}

		if len(index.ClusterMap.Centroids) > 0 {
			docEmb := NewDocEmbedding(doc.ID.Hex(), avgEmb)
			index.ClusterMap.IndexDoc(docEmb)
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

func readStructFromFile(filename string, data interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(data)
	if err != nil {
		return err
	}
	return nil
}

// BuildIndex Build a new index by calculating k Centroids using n random documents
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
			panic("can't read " + file.Name())
		}
		sampledDocEmbeddings[i] = NewDocEmbedding(file.Name(), embedding)
	}

	Cluster(sampledDocEmbeddings, k, index.ClusterMap)

	// assign all existing documents to cluster
	for _, file := range files {
		filePath := filepath.Join(index.avgDocumentPath, file.Name())
		doc, err := readAvgDocument(filePath)
		if err != nil {
			panic(err)
		}

		docEmb := NewDocEmbedding(file.Name(), doc)
		index.ClusterMap.IndexDoc(docEmb)
	}

	err = writeStructToFile(index.clusterMapPath, index.ClusterMap)
	if err != nil {
		fmt.Println(err)
	}
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
