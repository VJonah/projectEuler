package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(largestPalindromeProduct(999))
}

func largestPalindromeProduct(n int) int {
	largest := -1
	for i := n; i >= 900; i-- {
		for j := n; j >= 900; j-- {
			product := i * j
			if isPalindrome(product) {
				if product > largest {
					largest = product
				}
			}
		}
	}
	return largest
}

func isPalindrome(n int) bool {
	digits := strconv.Itoa(n)
	reversed := reverse(digits)
	if reversed == digits {
		return true
	}
	return false
}

func reverse(str string) string {
	reversed := ""
	for _, char := range str {
		reversed = string(char) + reversed
	}
	return reversed
}
