// Experiments
//////////////

package main

import "fmt"

// Fitness functions used in experiments
var exp1FitFuncs = []func(v []int) float64{
	uniformCO, linearCO, tightDeceptiveTF, randDeceptiveTF,
	tightNonDeceptiveTF, randNonDeceptiveTF,
}

var exp23FitFuncs = []func(v []int) float64{
	uniformCO, linearCO, tightDeceptiveTF, tightNonDeceptiveTF,
}

// Runs experiment given a slice of fitness functions, a crossover
// function, and whether or not mutation will be allowed.
func experiment(fitFuncs []func(v []int) float64, c func(p1, p2 []int) (o1, o2 []int), m bool) {
	for _, f := range fitFuncs {
		prev, n := 5, 10
		for {
			var p population
			success := 0

			//run genetic algorithm 30 times
			for j := 0; j < 30; j++ {
				p = fillPopulation(n, f, c, m)
				p = geneticAlgorithm(p)
				if containsOptimalSol(p) {
					success++
				}
			}
			//output results
			fmt.Printf("N=%-4d %2d/%d\n", n, success, 30)

			//update parameters or break
			n, prev = bisectionSearch(success, n, prev)
			if n % 10 != 0 || n > 1280 {
				break
			}
		}
	}
}

// Utils
////////

func bisectionSearch(success, n, prev int) (int, int) {
	if success >= 29 {
		return n - abs(n-prev)/2, n
	} else {
		if prev == n / 2 {
			return 2*n, n
		} else {
			return n + abs(n-prev)/2, n
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
