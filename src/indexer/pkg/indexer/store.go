package indexer

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/muesli/clusters"
	"github.com/muesli/kmeans"
)

type DocEmbedding struct {
	DocId     string
	Embedding Vector
}

func NewDocEmbedding(docId string, embedding Vector) *DocEmbedding {
	return &DocEmbedding{
		DocId:     docId,
		Embedding: embedding,
	}
}

type Index struct {
	indexPath string
	Centroids []Vector
}

func NewStore(indexPath string) *Index {
	cm := indexPath + "/cluster/clusterMap.index"
	var centroids []Vector

	err := ReadStructFromFile(cm, &centroids)
	if err != nil {
		fmt.Printf("warning: centroids not found: %s\n", err)
	}

	return &Index{
		indexPath: indexPath,
		Centroids: centroids,
	}
}

func (i *Index) getDocPath() string {
	return i.indexPath + "/documents"
}

func (i *Index) GetDocIdPath(docId string) string {
	return i.indexPath + "/documents/" + docId
}

func (i *Index) getAvgDocPath() string {
	return i.indexPath + "/documents_averaged"
}

func (i *Index) getClusterPath() string {
	return i.indexPath + "/cluster"
}

func (i *Index) getClusterMapPath() string {
	return i.indexPath + "/cluster/clusterMap.indexer"
}

// Store store document, averaged document to disk, indexer if available
func (i *Index) Store(doc *Document) {
	// fmt.Println("[Store] storing document", doc.Id)
	docPath := i.getDocPath() + "/" + doc.Id
	avgDocPath := i.getAvgDocPath() + "/" + doc.Id

	// Return if doc was already indexed
	_, err := os.Stat(docPath)
	if err == nil {
		return
	}

	// Store all embeddings of document + averaged embedding
	emb := GetEmbedding(doc.Text)
	avgEmb, err := GetAvgEmbedding(emb)
	errPanic(err)

	err = WriteStructToFile(docPath, emb)
	errPanic(err)

	err = WriteStructToFile(avgDocPath, avgEmb)
	errPanic(err)

	// documents can be stored before inverted-file-indexer was created
	if i.Centroids != nil {
		i.indexDoc(doc.Id, avgEmb)
	}
}

// BuildCluster (Re)build the indexer, by calculating clusters + centroids
func (i *Index) BuildCluster(n, k int) {
	fmt.Println("[BuildCluster] starting clustering routine")
	files, err := ioutil.ReadDir(i.getAvgDocPath())
	errPanic(err)

	// sample n documents and read their embeddings
	sampled := getDocumentSubset(files, n)
	embeddings := make([]Vector, len(sampled))
	for j, file := range sampled {
		path := filepath.Join(i.getAvgDocPath(), file.Name())

		var emb Vector
		err = ReadStructFromFile(path, &emb)
		errPanic(err)

		embeddings[j] = emb
	}

	fmt.Println("[BuildCluster] embeddings loaded")

	i.Centroids = cluster(embeddings, k)

	// assign all documents to cluster
	for _, file := range files {
		path := filepath.Join(i.getAvgDocPath(), file.Name())

		var emb Vector
		err = ReadStructFromFile(path, &emb)
		errPanic(err)

		i.indexDoc(file.Name(), emb)
	}

	centroidPath := i.getClusterMapPath()
	WriteStructToFile(centroidPath, i.Centroids)
}

func getDocumentSubset(files []os.FileInfo, count int) []os.FileInfo {
	numFiles := len(files)

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	random.Shuffle(numFiles, func(i, j int) {
		files[i], files[j] = files[j], files[i]
	})

	return files[:count]
}

// cluster apply kmeans
func cluster(embeddings []Vector, k int) []Vector {
	var d clusters.Observations
	for _, x := range embeddings {
		d = append(d, clusters.Coordinates(x))
	}

	km := kmeans.New()
	clusters, err := km.Partition(d, k)
	errPanic(err)

	centroids := make([]Vector, len(clusters))
	for i, c := range clusters {
		centroids[i] = Vector(c.Center)
	}

	return centroids
}

func (i *Index) indexDoc(docId string, emb Vector) error {
	docEmb := NewDocEmbedding(docId, emb)
	centroidId := i.GetClosestCentroid(emb)
	if centroidId == -1 {
		return errors.New("no centroid found")
	}
	centroidPath := i.getClusterPath() + "/centroid_" + strconv.Itoa(centroidId)

	// ToDo - Find serialization format that allows to append to existing file, without bad read performance
	// create new centroid file, if not exist
	var centroid []*DocEmbedding
	err := ReadStructFromFile(centroidPath, &centroid)
	if err != nil {
		centroid = make([]*DocEmbedding, 0)
	}
	centroid = append(centroid, docEmb)

	err = WriteStructToFile(centroidPath, centroid)
	errPanic(err)

	return nil
}

// GetClosestCentroid return id of closest centroid
func (i *Index) GetClosestCentroid(v Vector) int {
	centroidId := -1
	max := 0.0

	for i, c := range i.Centroids {
		cos := cosim(v, c)
		if cos > max {
			max = cos
			centroidId = i
		}
	}
	return centroidId
}

// GetClosestCentroidEmb Get all document embeddings of the closest centroid for Vector v
func (i *Index) GetClosestCentroidEmb(v Vector) ([]*DocEmbedding, error) {
	centroidId := i.GetClosestCentroid(v)
	centroidPath := i.getClusterPath() + "/centroid_" + strconv.Itoa(centroidId)

	var centroid []*DocEmbedding
	err := ReadStructFromFile(centroidPath, &centroid)
	if err != nil {
		return centroid, err
	}

	return centroid, nil
}
