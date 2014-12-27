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
	rand.Seed(42)

	p := fillPopulation(200, uniformCO, twoPointCrossover, false)

	geneticAlgorithm(p)
}
