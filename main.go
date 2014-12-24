package main

import "fmt"

const (
	l = 100
	k = 4
)

func main() {
	vector := make([]int, 100)
	for i := 0; i < l; i += 2 {
		vector[i] = 1
	}
	
	fmt.Println(nonDeceptiveTF(vector))
}
