package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(primeSum())
}

func primeSum() int {
	sum := 2
	for i := 3; i < 2000000; i += 2 {
		if isPrime(i) {
			sum += i
		}
	}
	return sum
}

func isPrime(n int) bool {
	root := int(math.Sqrt(float64(n))) + 1
	for i := 3; i < root; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
