package main

import (
	"fmt"
	"strconv"
)

func main() {
	sum := 0
	for i := 1; i <= 1000; i++ {
		sum += convertToWordsAndCount(i)
	}
	fmt.Println(sum)
}

func convertToWordsAndCount(n int) int {
	wordMap := map[string]string{
		"1":    "one",
		"2":    "two",
		"3":    "three",
		"4":    "four",
		"5":    "five",
		"6":    "six",
		"7":    "seven",
		"8":    "eight",
		"9":    "nine",
		"10":   "ten",
		"11":   "eleven",
		"12":   "twelve",
		"13":   "thirteen",
		"14":   "fourteen",
		"15":   "fifteen",
		"16":   "sixteen",
		"17":   "seventeen",
		"18":   "eighteen",
		"19":   "nineteen",
		"20":   "twenty",
		"30":   "thirty",
		"40":   "forty",
		"50":   "fifty",
		"60":   "sixty",
		"70":   "seventy",
		"80":   "eighty",
		"90":   "ninety",
		"100":  "hundred",
		"1000": "thousand",
	}
	digits := strconv.Itoa(n)
	words := ""
	switch {
	case n > 999:
		words = "onethousand"
	case n > 99:
		if wordMap[digits] == "" {
			words = wordMap[string(digits[0])] + wordMap["100"]
			if digits[1:] != "00" {
				words += "and"
			}
			if wordMap[digits[1:]] == "" {
				tens := string(digits[1]) + "0"
				unit := string(digits[2])
				words += wordMap[tens] + wordMap[unit]
			} else {
				words += wordMap[digits[1:]]
			}
		} else {
			if digits == "100" {
				words = "onehundred"
			} else {
				words = wordMap[digits]
			}
		}
	default:
		if wordMap[digits] == "" {
			tens := string(digits[0]) + "0"
			unit := string(digits[1])
			words = wordMap[tens] + wordMap[unit]
		} else {
			words = wordMap[digits]
		}
	}
	fmt.Println(words)
	return len(words)
}
