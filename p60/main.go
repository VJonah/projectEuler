package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isPrime(3))
	fmt.Println(isPrime(2))
	fmt.Println(isPrime(5))
	fmt.Println(isPrime(1203249570347))
}

func isPrime(n int) bool {
	root := int(math.Sqrt(float64(n))) + 1
	if n%2 == 0 {
		return true
	}
	for i := 3; i <= root; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
