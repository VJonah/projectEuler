package main

import (
	"fmt"
	"strconv"
)

func main() {
	findRecurring(10.00)
}

func findRecurring(n float64) int {
	var val float64
	for i := 2.00; i < n; i++ {
		val = 1.00 / i
		fmt.Println(getRecurringLength(val))
	}
	return -1
}

func getRecurringLength(n float64) int {
	digits := strconv.FormatFloat(n, 'f', -1, 64)
	for i, digit := range digits[2:] {
		nextIdx := findNextOccurence(digits, string(digit), i)
		if nextIdx != -1 {
			return nextIdx - i
		}
	}
	return -1
}

func findNextOccurence(baseString, s string, startIdx int) int {
	for i := startIdx + 1; i < len(baseString); i++ {
		if s == string(baseString[i]) {
			return i
		}
	}
	return -1
}

func contains(str, val string) bool {
	for _, char := range str {
		if val == string(char) {
			return true
		}
	}
	return false
}
