package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	input := parseInput("data.txt")
	fmt.Println(countNameScores(input))
}

func countNameScores(names []string) int {
	sum := 0
	for i, name := range names {
		nameValue := 0
		for i := 1; i < len(name)-1; i++ {
			letter := name[i]
			nameValue += int(letter) - 64
		}
		sum += nameValue * (i + 1)
	}
	return sum
}

func parseInput(href string) []string {
	data, _ := ioutil.ReadFile(href)
	names := strings.Split(string(data), ",")
	sort.Strings(names)
	return names
}
