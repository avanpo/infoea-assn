package main

import (
	"fmt"
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
	crossover func(p1, p2 []int) (o1, o2 []int)
	mutation  bool
}

func main() {
	rand.Seed(42)

	p := fillPopulation(6, uniformCO, twoPointCrossover, false)

	fmt.Println(p)
	sortPopulation(p)
	fmt.Println(p)
}
