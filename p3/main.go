package main

import (
	"fmt"
	"math"
)

func main() {
	input := 600851475143
	fmt.Println(largestPrimeFactor(input))
}

func largestPrimeFactor(n int) int {
	root := int(math.Floor(math.Sqrt(float64(n)))) + 1
	largest := 0
	for i := 3; i < root; i += 2 {
		if n%i == 0 && isPrime(i) {
			if i > largest {
				largest = i
			}
		}
	}
	return largest
}

func isPrime(n int) bool {
	root := int(math.Floor(math.Sqrt(float64(n)))) + 1
	for i := 2; i < root; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
