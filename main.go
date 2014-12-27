// Evolutionary Algorithms Assignment 1
// Author: Alex van Poppelen
///////////////////////////////////////

package main

import "math/rand"

const l = 100

func main() {
	rand.Seed(1)
	initRandomLink()

	// Experiment 1: 2-point crossover, no mutation
	experiment(exp1FitFuncs, twoPointCrossover, false)
}
