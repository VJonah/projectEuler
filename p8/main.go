package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input := parseInput("data.txt")
	fmt.Println(findLargestProduct(input, 13))

}

func findLargestProduct(digits string, numberOfDigits int) int {
	largest := 0
	for i := 0; i < len(digits)-numberOfDigits+1; i++ {
		section := digits[i : i+numberOfDigits]
		product := 1
		for _, digit := range section {
			if string(digit) == "0" {
				product = 0
				break
			}
			val, _ := strconv.Atoi(string(digit))
			product *= val
		}
		if product > largest {
			largest = product
		}
	}
	return largest
}

func parseInput(href string) string {
	data, _ := ioutil.ReadFile(href)
	input := string(data)
	lines := strings.Split(input, "\n")
	longStr := strings.Join(lines, "")
	fmt.Println(longStr)
	return longStr
}
