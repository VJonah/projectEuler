package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(pythagoreanTriplet())
}

func pythagoreanTriplet() int {
	for i := 1; i < 1000; i++ {
		aSqrd := i * i
		for j := 1; j < i; j++ {
			bSqrd := j * j
			sum := aSqrd + bSqrd
			root := math.Sqrt(float64(sum))
			if isIntegerRoot(root) {
				if int(root)+i+j == 1000 {
					fmt.Println(i, j, int(root))
					return i * j * int(root)
				}
			}
		}
	}
	return -1
}

func isIntegerRoot(root float64) bool {
	if root*root == float64(int(root)*int(root)) {
		return true
	}
	return false
}
