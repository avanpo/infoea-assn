package main

import "math/rand"

func shuffle(src []int) []int {
	v := make([]int, len(src))
	for i := range src {
		j := rand.Intn(i + 1)
		v[i] = v[j]
		v[j] = src[i]
	}
	return v
}
