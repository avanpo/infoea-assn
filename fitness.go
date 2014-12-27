// Fitness Functions
////////////////////

package main

import "math/rand"

// Uniformly scaled counting ones function.
func uniformCO(v []int) float64 {
	f := 0
	for _, x := range v {
		f += x
	}
	return float64(f)
}

// Linearly scaled counting ones function.
func linearCO(v []int) float64 {
	f := 0
	for i, x := range v {
		f += i * x
	}
	return float64(f)
}

// Tightly linked deceptive trap function.
func tightDeceptiveTF(v []int) float64 {
	return trapFunction(v, 4, 1)
}

// Randomly linked deceptive trap function.
func randDeceptiveTF(v []int) float64 {
	w := shuffle(v)
	return trapFunction(w, 4, 1)
}

// Tightly linked non-deceptive trap function.
func tightNonDeceptiveTF(v []int) float64 {
	return trapFunction(v, 4, 2.5)
}

// Randomly linked non-deceptive trap function.
func randNonDeceptiveTF(v []int) float64 {
	w := shuffle(v)
	return trapFunction(w, 4, 2.5)
}

// Helper functions
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

func shuffle(src []int) []int {
	v := make([]int, len(src))
	for i := range src {
		j := rand.Intn(i + 1)
		v[i] = v[j]
		v[j] = src[i]
	}
	return v // returns new copy, doesn't modify src
}
