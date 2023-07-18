package util

import (
	"fmt"
	"math"
	"math/rand"
)

type Vector []float64

func (v Vector) ToString() string {
	result := ""
	for _, e := range v {
		result += fmt.Sprintf("%f|", e)
	}
	return result
}

func (v Vector) Dot(other Vector) float64 {
	if len(v) != len(other) {
		panic(fmt.Sprintf("CosSim expected vectors of equal length, got %d and %d", len(v), len(other)))
	}
	var result float64 = 0.0
	for i, e := range v {
		result += e * other[i]
	}
	return result
}

func (v Vector) Norm() float64 {
	var result float64 = 0.0
	for _, e := range v {
		result += e * e
	}
	norm := math.Sqrt(result)
	return norm
}

func (v Vector) CosSim(other Vector) float64 {
	return v.Dot(other) / (v.Norm() * other.Norm())
}

func GetRandomVector(dims int) Vector {
	result := make(Vector, dims)
	for i := 0; i < dims; i++ {
		result[i] = rand.Float64()
	}
	return result
}
