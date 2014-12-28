// Evolutionary Algorithms Assignment 1
// Author: Alex van Poppelen
///////////////////////////////////////

package main

import "math/rand"

const l = 100

func main() {
	rand.Seed(42)
	initRandomLink()

	// Experiment 1: 2-point crossover, no mutation
	data := experiment(exp1FitFuncs, twoPointCrossover, false)
	writeData(1, data)

	// Experiment 2: uniform crossover, no mutation
	//experiment(exp23FitFuncs, uniformCrossover, false)

	// Experiment 3: uniform crossover, mutation
	//experiment(exp23FitFuncs, uniformCrossover, true)
}
