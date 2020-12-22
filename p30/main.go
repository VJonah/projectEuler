package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	powerMp := generatePowerMp(5)
	powerNs := findPowerNumbers(10000000, powerMp)
	fmt.Println(powerNs)
	fmt.Println(sumArr(powerNs))
}

func sumArr(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func findPowerNumbers(n int, powerMp map[int]int) []int {
	powerNmbrs := []int{}
	for i := 4000; i < n; i++ {
		sum := getDigitSum(i, powerMp)
		if sum == i {
			powerNmbrs = append(powerNmbrs, i)
		}
	}
	return powerNmbrs
}

func generatePowerMp(n int) map[int]int {
	powerMP := map[int]int{}
	for i := 0; i <= 9; i++ {
		power := int(math.Pow(float64(i), float64(n)))
		powerMP[i] = power
	}
	return powerMP
}

func getDigitSum(n int, powerMp map[int]int) int {
	digits := strconv.Itoa(n)
	sum := 0
	for _, digit := range digits {
		val, _ := strconv.Atoi(string(digit))
		sum += powerMp[val]
	}
	return sum
}
