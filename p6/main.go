package main

import "fmt"

func main() {
	sumN := sumToN(100)
	sumN *= sumN
	sumNSqrd := sumToNSquared(100)
	fmt.Println(sumN - sumNSqrd)

}

func sumToN(n int) int {
	return (n * (n + 1)) / 2
}

func sumToNSquared(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}
