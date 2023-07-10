package util

import (
	"github.com/muesli/clusters"
	"github.com/muesli/kmeans"
	"golang.org/x/exp/slices"
	"strconv"
)

type DocEmbedding struct {
	DocId     string
	Embedding Vector
}

type ClusterMap struct {
	Centroids    []Vector
	centroidsStr []string
	baseDir      string
}

func NewDocEmbedding(docId string, embedding Vector) *DocEmbedding {
	return &DocEmbedding{
		DocId:     docId,
		Embedding: embedding,
	}
}

func NewClusterMap(baseDir string) *ClusterMap {
	return &ClusterMap{
		Centroids:    make([]Vector, 0),
		centroidsStr: make([]string, 0),
		baseDir:      baseDir,
	}
}

func (cm *ClusterMap) GetCentroidCluster(cIdx int) []*DocEmbedding {
	var centroidEmb []*DocEmbedding
	centroidPath := cm.baseDir + "/centroid_" + strconv.Itoa(cIdx)
	err := readStructFromFile(centroidPath, &centroidEmb)
	if err != nil {
		panic("can't read file")
	}
	return centroidEmb
}

func (cm *ClusterMap) addCentroid(v Vector) {
	vStr := v.ToString()
	if !slices.Contains(cm.centroidsStr, vStr) {
		cm.centroidsStr = append(cm.centroidsStr, vStr)
		cm.Centroids = append(cm.Centroids, v)
	}
}

func (cm *ClusterMap) GetClosetCentroid(d *DocEmbedding) int {
	cIdx := -1
	maxCosim := 0.0

	// get centroid with smalles cosine similarity
	for i, centroid := range cm.Centroids {
		cos := Cosim(d.Embedding, centroid)
		if cos > maxCosim {
			maxCosim = cos
			cIdx = i
		}
	}
	return cIdx
}

func (cm *ClusterMap) IndexDoc(d *DocEmbedding) {
	cIdx := cm.GetClosetCentroid(d)
	centroidPath := cm.baseDir + "/centroid_" + strconv.Itoa(cIdx)

	// ToDo - Find serialization format that allows to append to existing file, without bad read performance
	var centroidEmb []*DocEmbedding
	err := readStructFromFile(centroidPath, &centroidEmb)

	// Create new slice, if file does not exist
	if err != nil {
		centroidEmb = make([]*DocEmbedding, 1)
		centroidEmb[0] = d
	} else {
		centroidEmb = append(centroidEmb, d)
	}

	err = writeStructToFile(centroidPath, centroidEmb)
	if err != nil {
		panic(err)
	}
}

func Cluster(embeddings []*DocEmbedding, k int, cm *ClusterMap) {
	var d clusters.Observations
	for _, x := range embeddings {
		d = append(d, clusters.Coordinates(x.Embedding))
	}

	// calculate Centroids
	km := kmeans.New()
	clusters, err := km.Partition(d, k)
	if err != nil {
		panic(err)
	}

	for _, c := range clusters {
		cm.addCentroid(Vector(c.Center))
	}
}
