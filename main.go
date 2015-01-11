// Evolutionary Algorithms Assignment 1
// Author: Alex van Poppelen
///////////////////////////////////////

package main

import (
	"math/rand"
	//"time"
	//"fmt"
)

const l = 100

func main() {
	// Initialize rand and random linkage
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

	// Test execution time for single algorithm run
	/*start := time.Now()

	p := fillPopulation(1280, tightDeceptiveTF, twoPointCrossover, false)
	p = geneticAlgorithm(p, false)

	elapsed := time.Since(start)

	if containsOptimalSol(p) {
		fmt.Printf("Found optimal solution in generation %d (%s)\n", p.t, elapsed)
	} else {
		fmt.Printf("Did not find optimal solution\n")
	}*/
}
