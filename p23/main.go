package main

import "fmt"

func main() {
	//fmt.Println(getAbundantNmbrs(28124))
	abundant := getAbundantNmbrs(30)
	fmt.Println(abundant)
	nonA := getNonAbundandSumNmbrs(abundant, 30)
	fmt.Println(nonA)

}

func getNonAbundandSumNmbrs(abundant []int, n int) int {
	sum := 0
	notAbundant := []int{}
	for i := 1; i < n; i++ {
		abundantN := abundant[0]
		isAbundantSum := true
		for j := 1; abundantN <= i && j < len(abundant) - 1; j++ {
			difference := j - abundantN
			if contains(abundant, difference) {
				isAbundantSum = false
				break
			}
			abundantN = abundant[j]
		}
		if isAbundantSum {
			notAbundant = append(notAbundant, i)
			sum += i
		}
	}
	fmt.Println(notAbundant)
	return sum
}

func contains(arr []int, val int) bool {
	for _, n := range arr {
		if n == val {
			return true
		}
	}
	return false
}

func getAbundantNmbrs(n int) []int {
	abundant := []int{}
	for i := 12; i < n; i++ {
		if sumArr(getDivisors(i)) > i {
			abundant = append(abundant, i)
		}
	}
	return abundant
}

func sumArr(arr []int) int {
	sum := 0
	for _, n := range arr {
		sum += n
	}
	return sum
}

func getDivisors(n int) []int {
	divisors := []int{1}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors
}
