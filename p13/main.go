package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	input := parseInput("data.txt")
	fmt.Println(firstTenDigits(input)[:10])
}

func firstTenDigits(numbers []string) string {
	firstTenDigits := ""
	carry := 0
	for i := len(numbers[0]) - 1; i > -1; i-- {
		sum := carry
		for _, line := range numbers {
			n, _ := strconv.Atoi(string(line[i]))
			sum += n
		}
		if sum > 10 {
			carry = int(math.Floor(float64(sum) / 10.0))
			sum = sum % 10
		} else {
			carry = 0
		}
		digit := strconv.Itoa(sum)
		firstTenDigits = digit + firstTenDigits
	}
	if carry > 0 {
		digits := strconv.Itoa(carry)
		firstTenDigits = digits + firstTenDigits
	}
	return firstTenDigits
}

func parseInput(href string) []string {
	data, _ := ioutil.ReadFile(href)
	return strings.Split(string(data), "\n")
}
