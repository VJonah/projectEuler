package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(generatePrimes())
}

func generatePrimes() int {
	primes := []int{2}
	count := 1
	for i := 3; true; i += 2 {
		if !factorInArray(primes, i) {
			count++
			primes = append(primes, i)
		}
		if count == 10001 {
			break
		}
	}
	return primes[count-1]
}

func factorInArray(arr []int, val int) bool {
	root := int(math.Sqrt(float64(val))) + 1
	for _, n := range arr {
		if n > root {
			return false
		}
		if val%n == 0 {
			return true
		}
	}
	return false
}
