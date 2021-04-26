package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findConsecDistinctPrimes(3))
}

func findConsecDistinctPrimes(n int) []int {
	for i := 600; i > 0; i++ {
		seenPrimes := map[int]int{}
		for j := 0; j < n; j++ {
			current := i + j
			factors := primeFactor(current, []int{})
			prev := factors[0]
			product := prev
			notDistinct := false
			for k := 1; k < len(factors); k++ {
				factor := factors[k]
				if factor == prev {
					product *= factor
					continue
				} else {
					if product > 1 {
						seenPrimes[product]++
					} else if seenPrimes[factor] == 0 {
						seenPrimes[factor]++
					} else {
						notDistinct = true
						break
					}
				}
				product = 1
			}
			if notDistinct {
				break
			} else {
				return []int{i, i + 1, i + 2}
			}
		}
	}
	return []int{}
}

func primeFactor(n int, factors []int) []int {
	if n == 1 {
		return factors
	}
	root := int(math.Sqrt(float64(n))) + 1
	for i := 2; i < root; i++ {
		if isPrime(i) && n%i == 0 {
			factors = append(factors, i)
			return primeFactor(n/i, factors)
		}
	}
	factors = append(factors, n)
	return factors
}

func isPrime(n int) bool {
	root := int(math.Sqrt(float64(n))) + 1
	for i := 2; i < root; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
