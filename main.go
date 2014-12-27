// Evolutionary Algorithms Assignment 1
///////////////////////////////////////
// Author: Alex van Poppelen
///////////////////////////////////////

package main

import (
	"math/rand"
)

const (
	l = 100
)

type population struct {
	n         int
	sols      [][]int
	fits      []float64
	fitness   func(v []int) float64
	crossover func(p1, p2 []int) ([]int, []int)
	mutation  bool
}

func main() {
	rand.Seed(1)

	p := fillPopulation(400, tightDeceptiveTF, twoPointCrossover, true)
	geneticAlgorithm(p)
}
