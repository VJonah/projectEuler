package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(getCuriousNmbrs(10000000))
}

func getCuriousNmbrs(n int) []int {
	curious := []int{}
	for i := 145; i < n; i++ {
		if digitFactSum(i) == i {
			curious = append(curious, i)
		}
	}
	return curious
}

func digitFactSum(n int) int {
	factMap := map[int]int{
		0: 1,
		1: 1,
		2: 2,
		3: 6,
		4: 24,
		5: 120,
		6: 720,
		7: 5040,
		8: 40320,
		9: 362880,
	}
	sum := 0

	for _, digit := range strconv.Itoa(n) {
		val, _ := strconv.Atoi(string(digit))
		sum += factMap[val]
	}
	return sum
}
