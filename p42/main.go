package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	words := parseInput("data.txt")
	fmt.Println(countTriangleWords(words))
}

func countTriangleWords(words []string) int {
	count := 0
	for _, word := range words {
		value := calculateWordValue(word)
		triangle := 0
		for i := 1; value > triangle; i++ {
			triangle = (i * (i + 1)) / 2
			if triangle == value {
				fmt.Println(value, i)
				count++
				break
			}
		}
	}
	return count
}

func calculateWordValue(word string) int {
	sum := 0
	for i := 1; i < len(word)-1; i++ {
		sum += int(word[i]) - 64
	}
	return sum
}

func parseInput(href string) []string {
	data, _ := ioutil.ReadFile(href)
	words := strings.Split(string(data), ",")
	return words
}
