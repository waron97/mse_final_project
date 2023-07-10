package util

import (
	"fmt"
	"github.com/muesli/clusters"
	"github.com/muesli/kmeans"
)

type DocEmbedding struct {
	DocId     string
	Embedding Vector
}

type CentroidMap struct {
	data      map[string][]*DocEmbedding
	centroids []Vector
}

func NewDocEmbedding(docId string, embedding Vector) *DocEmbedding {
	return &DocEmbedding{
		DocId:     docId,
		Embedding: embedding,
	}
}

func NewCentroidMap() *CentroidMap {
	return &CentroidMap{
		data:      make(map[string][]*DocEmbedding),
		centroids: make([]Vector, 0),
	}
}

func (cm *CentroidMap) Insert(key Vector, value *DocEmbedding) {
	strKey := key.ToString()
	if _, ok := cm.data[strKey]; ok {
		cm.data[strKey] = append(cm.data[strKey], value)
		cm.centroids = append(cm.centroids, key)
	} else {
		cm.data[strKey] = []*DocEmbedding{value}
	}
}

func (cm *CentroidMap) addCentroid(v Vector) {
	strKey := v.ToString()
	if _, ok := cm.data[strKey]; !ok {
		cm.data[strKey] = []*DocEmbedding{}
		cm.centroids = append(cm.centroids, v)
	}
}

func (cm *CentroidMap) Get(key Vector) []*DocEmbedding {
	strKey := key.ToString()
	return cm.data[strKey]
}

func (cm *CentroidMap) IndexDoc(d *DocEmbedding) {
	maxCosim := 0.0
	var maxCentroid Vector

	// get closest centroid
	for _, centroid := range cm.centroids {
		cos := Cosim(d.Embedding, centroid)
		if cos > maxCosim {
			maxCosim = cos
			maxCentroid = centroid
		}
	}

	cm.Insert(maxCentroid, d)

	for key, value := range cm.data {
		fmt.Println("Documents with Vector", key[0], key[1], key[2], key[3])
		for _, doc := range value {
			fmt.Println("DOC", doc.DocId)
		}
	}
}

func Cluster(embeddings []*DocEmbedding, k int) {
	var d clusters.Observations
	for _, x := range embeddings {
		d = append(d, clusters.Coordinates(x.Embedding))
	}

	// calculate centroids
	km := kmeans.New()
	clusters, err := km.Partition(d, k)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	cm := NewCentroidMap()
	for _, c := range clusters {
		cm.addCentroid(Vector(c.Center))
	}

	for _, e := range embeddings {
		cm.IndexDoc(e)
	}
}
