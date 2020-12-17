package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findTriangleNumber())
}

func findTriangleNumber() int {
	count := 21
	for i := 7; true; i++ {
		count += i
		if numberOfDivisors(count) > 500 {
			return count
		}
	}
	return -1
}

func numberOfDivisors(n int) int {
	root := int(math.Sqrt(float64(n))) + 1
	//count starts at 2 because the number itself and 1 are always divisors beyond 3
	count := 2
	for i := 2; i < root; i++ {
		if n%i == 0 {
			count++
		}
	}
	return count * 2
}
