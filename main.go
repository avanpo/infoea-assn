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
	data1 := experiment(exp1FitFuncs, twoPointCrossover, false)
	writeData(1, data1)

	// Experiment 2: uniform crossover, no mutation
	data2 := experiment(exp23FitFuncs, uniformCrossover, false)
	writeData(2, data2)

	// Experiment 3: uniform crossover, mutation
	data3 := experiment(exp23FitFuncs, uniformCrossover, true)
	writeData(3, data3)
}
