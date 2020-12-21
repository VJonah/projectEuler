package main

import (
	"math"
	"strconv"
)

func main() {
	largeFib(1000)
	//fmt.Println(addLargeNumbers([]int{8}, []int{1, 3}))
}

func largeFib(n int) int {
	idx := 1
	prev := []int{1}
	cur := []int{1}
	nxt := []int{}
	for len(cur) != n {
		nxt = addLargeNumbers(prev, cur)
		prev = cur
		cur = nxt
		idx++
	}
	return idx
}

func addLargeNumbers(smaller, larger []int) []int {
	newNumber := []int{}
	carry := 0
	lenDiff := len(larger) - len(smaller)
	for i := len(larger) - 1; i > -1; i-- {
		sum := carry
		lrgVal := larger[i]
		smllVal := 0
		if i-lenDiff > -1 {
			smllVal = smaller[i-lenDiff]
		}
		sum += lrgVal + smllVal
		if sum >= 10 {
			carry = int(math.Floor(float64(sum) / 10.0))
			sum = sum % 10
		} else {
			carry = 0
		}
		newNumber = append([]int{sum}, newNumber...)
	}
	if carry > 0 {
		digits := strconv.Itoa(carry)
		for i := len(digits) - 1; i > -1; i-- {
			digit := digits[i]
			val, _ := strconv.Atoi(string(digit))
			newNumber = append([]int{val}, newNumber...)
		}
	}
	return newNumber
}
