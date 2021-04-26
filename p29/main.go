package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(len(intCombinations(2, 100)))
}

func intCombinations(start, end int) map[string]int {
	terms := map[string]int{}
	for i := start; i <= end; i++ {
		digits := getDigits(i)
		for j := start; j <= end; j++ {
			digits = timesLargeNumber(digits, i)
			str := joinInt(digits)
			if terms[str] == 0 {
				terms[str]++
			}
		}
	}
	return terms
}

func joinInt(arr []int) string {
	s := ""
	for _, n := range arr {
		s += strconv.Itoa(n)
	}
	return s
}

func getDigits(n int) []int {
	strForm := strconv.Itoa(n)
	digits := []int{}
	for _, digit := range strForm {
		val, _ := strconv.Atoi(string(digit))
		digits = append(digits, val)
	}
	return digits
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
