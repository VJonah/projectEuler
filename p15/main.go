package main

import (
	"fmt"
)

func main() {
	fmt.Println(countLattice([]int{20, 20}, map[string]int{}))
}

func countLattice(coordinate []int, square map[string]int) int {
	x := coordinate[0]
	y := coordinate[1]
	if x == 0 || y == 0 {
		return 1
	}
	left := []int{x - 1, y}
	up := []int{x, y - 1}
	leftC := fmt.Sprintf("%d,%d", x-1, y)
	upC := fmt.Sprintf("%d,%d", x, y-1)
	if square[leftC] == 0 {
		square[leftC] = countLattice(left, square)
	}
	if square[upC] == 0 {
		square[upC] = countLattice(up, square)
	}
	return square[leftC] + square[upC]
}
