package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(getPandigitalPrimes(1000000000))
}

func getPandigitalPrimes(n int) int {
	largestP := -1
	for i := 123456789; i < n; i += 2 {
		if isPrime(i) && i > largestP {
			fmt.Println("here")
			largestP = i
		}
	}
	return largestP
}

func isPrime(n int) bool {
	root := int(math.Sqrt(float64(n))) + 1
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i < root; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
