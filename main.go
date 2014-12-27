// Evolutionary Algorithms Assignment 1
// Author: Alex van Poppelen
///////////////////////////////////////

package main

import "math/rand"

const l = 100

// Population container structure.
type population struct {
	n         int
	t         int
	sols      [][]int
	fits      []float64
	fitness   func(v []int) float64
	crossover func(p1, p2 []int) ([]int, []int)
	mutation  bool
}

func main() {
	rand.Seed(1)
	initRandomLink()

	p := fillPopulation(20, tightDeceptiveTF, twoPointCrossover, true)
	p = geneticAlgorithm(p)

	printPopulation(p)
	printPopulationStats(p)
}
