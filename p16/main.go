package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(generatePowers())
}

func generatePowers() int {
	/*
		power := strconv.Itoa(int(math.Pow(2.0, 62.0)))
		powers := []int{}
		for _, char := range power {
			digit, _ := strconv.Atoi(string(char))
			powers = append([]int{digit}, powers...)
		}
	*/
	powers := []int{3, 2, 7, 6, 8}
	for i := 15; i <= 1000; i++ {
		powers = doubleLargeNumber(powers)
		fmt.Println(i)
		fmt.Println(powers)
		fmt.Println(sumDigits(powers))
	}
	return sumDigits(powers)
}

func sumDigits(n []int) int {
	sum := 0
	for _, digit := range n {
		sum += digit
	}
	return sum
}

func doubleLargeNumber(n []int) []int {
	newNumber := []int{}
	carry := 0
	for i := len(n) - 1; i > -1; i-- {
		sum := carry
		val := n[i]
		sum += 2 * val
		if sum >= 10 {
			carry = int(math.Floor(float64(sum) / 10.0))
			sum = sum % 10
		} else {
			carry = 0
		}
		newNumber = append([]int{sum}, newNumber...)
	}
	if carry > 0 {
		digits := strconv.Itoa(carry)
		for _, digit := range digits {
			val, _ := strconv.Atoi(string(digit))
			newNumber = append([]int{val}, newNumber...)
		}
	}
	return newNumber
}
