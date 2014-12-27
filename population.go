// Population control logic
///////////////////////////

package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// Initialize empty population structure.
func initPopulation(n int, f func(v []int) float64, c func(p1, p2 []int) (o1, o2 []int), m bool) (p population) {
	p.n = n
	p.t = 0
	p.sols = make([][]int, n)
	p.fits = make([]float64, n)
	p.fitness = f
	p.crossover = c
	p.mutation = m
	return
}

// Initialize and fill population structure with randomly
// generated solutions.
func fillPopulation(n int, f func(v []int) float64, c func(p1, p2 []int) (o1, o2 []int), m bool) population {
	p := initPopulation(n, f, c, m)
	for i := range p.sols {
		p.sols[i] = make([]int, l)
		for j := range p.sols[i] {
			p.sols[i][j] = rand.Intn(2)
		}
		p.fits[i] = p.fitness(p.sols[i])
	}
	p.t = 1
	return p
}

// Randomizes the population order using the Fisher-Yates
// shuffle.
func shufflePopulation(p population) {
	for i := range p.sols {
		j := rand.Intn(i + 1)
		p.sols[i], p.sols[j] = p.sols[j], p.sols[i]
		p.fits[i], p.fits[j] = p.fits[j], p.fits[i]
	}
}

// Sorts the population according to fitness.
func sortPopulation(p population) {
	sort.Sort(p)
}

// Implementation of sort.Interface
///////////////////////////////////

func (s population) Len() int {
	return len(s.sols)
}

func (s population) Swap(i, j int) {
	s.sols[i], s.sols[j] = s.sols[j], s.sols[i]
	s.fits[i], s.fits[j] = s.fits[j], s.fits[i]
}

func (s population) Less(i, j int) bool {
	return s.fits[i] > s.fits[j] //sort high-to-low
}

// Printing/Debugging functions
///////////////////////////////

func printPopulation(p population) {
	for i := range p.sols {
		fmt.Printf("%6.1f  ", p.fits[i])
		for j := 0; j < l; j++ {
			fmt.Printf("%d", p.sols[i][j])
		}
		fmt.Printf("\n")
	}
}

func printPopulationShort(p population) {
	for i := range p.fits {
		if i > 0 && i % 11 == 0 {
			fmt.Printf("\n")
		}
		fmt.Printf("%6.1f ", p.fits[i])
	}
	fmt.Printf("\n")
}

func printPopulationStats(p population) {
	best := p.fits[0]
	worst := p.fits[0]
	total := 0.0
	for i := range p.fits {
		if p.fits[i] > best {
			best = p.fits[i]
		}
		if p.fits[i] < worst {
			worst = p.fits[i]
		}
		total += p.fits[i]
	}
	total = total / float64(p.n)
	fmt.Printf("Generation %d\n", p.t)
	fmt.Printf("  Average fitness: %5.1f\n", total)
	fmt.Printf("  Best fitness:    %5.1f  Worst: %5.1f\n", best, worst)
}

func verifyPopulationFitness(p population) bool {
	for i := range p.sols {
		if p.fitness(p.sols[i]) != p.fits[i] {
			return false
		}
	}
	return true
}
