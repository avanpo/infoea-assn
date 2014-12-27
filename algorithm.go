// Genetic Algorithm logic
//////////////////////////

package main

import (
"fmt"
"math/rand"
)

// Main algorithm loop. On return, the population is sorted.
func geneticAlgorithm(p population) population {
	for i := 0; true; i++ {
		//printPopulationStats(p)
		if !verifyPopulationFitness(p) {
			fmt.Println("FITNESS ERROR")
			break
		}

		parents := selectParents(p)
		offspring := generateOffspring(parents)

		stop := false
		p, stop = selectNewGeneration(p, offspring)
		if stop {
			break
		}

		shufflePopulation(p)
	}
	return p
}

// Returns sorted new generation, with a stop flag if no offspring
// solution has entered the new generation.
func selectNewGeneration(p population, po population) (population, bool) {
	newPop := initPopulation(p.n, p.fitness, p.crossover, p.mutation)
	sortPopulation(p)
	sortPopulation(po)
	a := 0
	b := 0
	for i := range newPop.sols {
		if p.fits[a] >= po.fits[b] {
			newPop.sols[i] = p.sols[a]
			newPop.fits[i] = p.fits[a]
			a++
		} else {
			newPop.sols[i] = po.sols[b]
			newPop.fits[i] = po.fits[b]
			b++
		}
	}
	newPop.t = p.t + 1
	if b == 0 {
		return newPop, true
	}
	return newPop, false
}

// Generates offspring from selected parents. Parent population must
// be shuffled. Returns shuffled offspring population.
func generateOffspring(pp population) population {
	po := initPopulation(pp.n, pp.fitness, pp.crossover, pp.mutation)
	for i := 0; i < pp.n; i += 2 {
		po.sols[i], po.sols[i+1] = po.crossover(pp.sols[i], pp.sols[i+1])
		if pp.mutation {
			mutate(po.sols[i])
			mutate(po.sols[i+1])
		}
		po.fits[i] = po.fitness(po.sols[i])
		po.fits[i+1] = po.fitness(po.sols[i+1])
	}
	return po
}

// Selects parents through tournament selection (s = 2). Population
// parameter must be shuffled. Returns shuffled parent population.
func selectParents(p population) population {
	pp := initPopulation(p.n, p.fitness, p.crossover, p.mutation)
	for i := 0; i < p.n; i += 2 {
		if p.fits[i] > p.fits[i+1] {
			pp.sols[i/2] = p.sols[i]
			pp.fits[i/2] = p.fits[i]
		} else {
			pp.sols[i/2] = p.sols[i+1]
			pp.fits[i/2] = p.fits[i+1]
		}
	}
	shufflePopulation(p)
	for i := 0; i < p.n; i += 2 {
		if p.fits[i] > p.fits[i+1] {
			pp.sols[pp.n/2+i/2] = p.sols[i]
			pp.fits[pp.n/2+i/2] = p.fits[i]
		} else {
			pp.sols[pp.n/2+i/2] = p.sols[i+1]
			pp.fits[pp.n/2+i/2] = p.fits[i+1]
		}
	}
	return pp
}

// Variation functions
//////////////////////

func mutate(v []int) {
	for {
		i := rand.Intn(len(v))
		v[i] = 1 - v[i]
		if rand.Intn(2) == 0 {
			break
		}
	}
}

func twoPointCrossover(p1, p2 []int) ([]int, []int) {
	o1 := make([]int, len(p1))
	o2 := make([]int, len(p1))

	left := rand.Intn(len(p1))
	right := rand.Intn(len(p1) + 1)
	if left > right {
		left, right = right, left
	}

	for i := range p1 {
		if i < left || i >= right {
			o1[i] = p1[i]
			o2[i] = p2[i]
		} else {
			o1[i] = p2[i]
			o2[i] = p1[i]
		}
	}
	return o1, o2
}

func uniformCrossover(p1, p2 []int) ([]int, []int) {
	o1 := make([]int, len(p1))
	o2 := make([]int, len(p1))
	for i := range p1 {
		if rand.Intn(2) == 0 {
			o1[i] = p1[i]
			o2[i] = p2[i]
		} else {
			o1[i] = p2[i]
			o2[i] = p1[i]
		}
	}
	return o1, o2
}
