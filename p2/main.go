package main

import "fmt"

func main() {
	previous := 1
	secondPrev := 1
	next := 0
	sum := 0
	for next < 4000000 {
		next = previous + secondPrev
		if next%2 == 0 {
			sum += next
		}
		secondPrev = previous
		previous = next
	}
	fmt.Println(sum)
}
