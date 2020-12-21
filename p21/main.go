package main

import "fmt"

func main() {
	amicable := getAmicableNumbers(10000)
	fmt.Println(amicable)
	fmt.Println(sumArr(amicable))
}

func getAmicableNumbers(n int) []int {
	amicable := []int{}
	for a := 1; a < n; a++ {
		divisorsA := getDivisors(a)
		b := sumArr(divisorsA)
		divisorsB := getDivisors(b)
		sumB := sumArr(divisorsB)
		//fmt.Printf("a: %d, b: %d\n", a, b)
		if sumB == a && b != a {
			amicable = append(amicable, a)
		}
	}
	return amicable
}

func sumArr(arr []int) int {
	sum := 0
	for _, n := range arr {
		sum += n
	}
	return sum
}

func getDivisors(n int) []int {
	divisors := []int{1}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors
}
