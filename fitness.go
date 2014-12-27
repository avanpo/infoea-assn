// Fitness functions
////////////////////

package main

import "math/rand"

func uniformCO(v []int) float64 {
	f := 0
	for _, x := range v {
		f += x
	}
	return float64(f)
}

func linearCO(v []int) float64 {
	f := 0
	for i, x := range v {
		f += i * x
	}
	return float64(f)
}

func tightDeceptiveTF(v []int) float64 {
	return trapFunction(v, 4, 1)
}

func randDeceptiveTF(v []int) float64 {
	w := shuffle(v)
	return trapFunction(w, 4, 1)
}

func tightNonDeceptiveTF(v []int) float64 {
	return trapFunction(v, 4, 2.5)
}

func randNonDeceptiveTF(v []int) float64 {
	w := shuffle(v)
	return trapFunction(w, 4, 2.5)
}

// helper functions
///////////////////

func trapFunction(v []int, k float64, d float64) (f float64) {
	for i := 0; i < l; i += int(k) {
		co := float64(uniformCO(v[i : i+int(k)]))
		if co == k {
			f += k
		} else {
			f += k - d - ((k-d)/(k-1))*co
		}
	}
	return
}

// returns a new copy, doesn't modify source
func shuffle(src []int) []int {
	v := make([]int, len(src))
	for i := range src {
		j := rand.Intn(i + 1)
		v[i] = v[j]
		v[j] = src[i]
	}
	return v
}
