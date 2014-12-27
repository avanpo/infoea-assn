package main

import "math/rand"

func recombine(p1, p2 []int, c func(p1, p2 []int) (o1, o2 []int), m bool) (output1, output2 []int) {
	return
}

func twoPointCrossover(p1, p2 []int) (o1, o2 []int) {
	left := rand.Intn(len(p1))
	right := rand.Intn(len(p2) + 1)
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
	return
}

func generateOffspring(pp population) population {
	po := initPopulation(pp.n, pp.fitness, pp.crossover, pp.mutation)
	for i := 0; i < pp.n; i += 2 {
		po.sols[i], po.sols[i+1] = recombine(pp.sols[i], pp.sols[i+1], pp.crossover, pp.mutation)
		po.fits[i] = po.fitness(po.sols[i])
		po.fits[i+1] = po.fitness(po.sols[i+1])
	}
	return po
}

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
