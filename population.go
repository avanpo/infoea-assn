package main

import (
	"math/rand"
	"sort"
)

func initPopulation(n int, f func(v []int) float64, c func(p1, p2 []int) (o1, o2 []int), m bool) (p population) {
	p.n = n
	p.sols = make([][]int, n)
	p.fits = make([]float64, n)
	p.fitness = f
	p.crossover = c
	p.mutation = m
	return
}

func fillPopulation(n int, f func(v []int) float64, c func(p1, p2 []int) (o1, o2 []int), m bool) population {
	p := initPopulation(n, f, c, m)
	for i := range p.sols {
		p.sols[i] = make([]int, l)
		for j := range p.sols[i] {
			p.sols[i][j] = rand.Intn(2)
		}
		p.fits[i] = f(p.sols[i])
	}
	return p
}

func shufflePopulation(p population) {
	for i := range p.sols {
		j := rand.Intn(i + 1)
		p.sols[i], p.sols[j] = p.sols[j], p.sols[i]
		p.fits[i], p.fits[j] = p.fits[j], p.fits[i]
	}
}

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
