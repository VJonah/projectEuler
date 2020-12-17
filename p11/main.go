package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input := parseInput("data.txt")
	fmt.Println(largestProduct(input))
}

func largestProduct(grid [][]int) int {
	rowLen := len(grid[0])
	largest := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < rowLen; j++ {
			directionsTried := tryDirections(i, j, grid)
			if directionsTried > largest {
				largest = directionsTried
			}
		}
	}
	return largest
}

func tryDirections(x, y int, grid [][]int) int {
	rowLen := len(grid[0])
	colLen := len(grid)
	directions := [][]int{
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
		{-1, -1},
		{-1, 0},
		{-1, 1},
	}
	largest := 0
	for _, direction := range directions {
		row := x
		col := y
		product := 1
		for i := 0; i < 4; i++ {
			if (row >= 0 && row < rowLen) && (col >= 0 && col < colLen) {
				product *= grid[row][col]
			} else {
				break
			}
			row += direction[0]
			col += direction[1]
		}
		if product > largest {
			largest = product
		}
	}
	return largest
}

func parseInput(href string) [][]int {
	data, _ := ioutil.ReadFile(href)
	input := string(data)
	lineSplit := strings.Split(input, "\n")
	numberGrid := [][]int{}
	for _, line := range lineSplit {
		numbers := strings.Split(line, " ")
		newRow := []int{}
		for _, n := range numbers {
			val, _ := strconv.Atoi(n)
			newRow = append(newRow, val)
		}
		numberGrid = append(numberGrid, newRow)
	}
	return numberGrid
}
