package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fact := generateFactorial(100)
	fmt.Println(sumDigits(fact))
}

func generateFactorial(n int) []int {
	factorial := []int{1}
	for i := 2; i <= n; i++ {
		factorial = timesLargeNumber(factorial, i)
	}
	return factorial
}

func sumDigits(n []int) int {
	sum := 0
	for _, digit := range n {
		sum += digit
	}
	return sum
}

func timesLargeNumber(largeN []int, n int) []int {
	newNumber := []int{}
	carry := 0
	for i := len(largeN) - 1; i > -1; i-- {
		sum := carry
		val := largeN[i]
		sum += n * val
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
		for i := len(digits) - 1; i > -1; i-- {
			digit := digits[i]
			val, _ := strconv.Atoi(string(digit))
			newNumber = append([]int{val}, newNumber...)
		}
	}
	return newNumber
}
