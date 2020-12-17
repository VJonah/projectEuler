package main

import "fmt"

func main() {
	fmt.Println(bruteForceSearch())
}

func bruteForceSearch() int {
	longest := []int{13, 10}
	for i := 13; i < 1000000; i++ {
		length := iterateSequence(i)
		if length > longest[1] {
			longest[0] = i
			longest[1] = length
		}
	}
	return longest[0]
}

func iterateSequence(n int) int {
	length := 1
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		length++
	}
	return length
}
