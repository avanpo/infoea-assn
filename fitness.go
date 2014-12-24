package main

func uniformlyScaledCO(vector []int) (fitness int) {
	for _, x := range vector {
		fitness += x
	}
	return
}

func linearlyScaledCO(vector []int) (fitness int) {
	for i, x := range vector {
		fitness += i * x
	}
	return
}

func deceptiveTF(vector []int) float64 {
	return trapFunction(vector, 1)
}

func nonDeceptiveTF(vector []int) float64 {
	return trapFunction(vector, 2.5)
}

func trapFunction(vector []int, d float64) (f float64) {
	for i := 0; i < l; i += k {
		co := float64(uniformlyScaledCO(vector[i:i+k]))
		if co == k {
			f += k
		} else {
			f += k - d - ((k - d) / (k - 1)) * co
		}
	}
	return
}
