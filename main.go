// -*- tab-width: 2 -*-

package main

import (
	"fmt"
	"strings"
)

var digits [8192]uint8

const (
	numDigits = 8192
	numTerms  = 3400
	ten       = 10
)

var onetwothrees = []string{
	"0",
	"1",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
}

func divide(n int) {
	var val int

	val = 0
	for i := range numDigits {
		val = val*ten + int(digits[i])
		if val < n {
			digits[i] = 0

			continue
		}

		digits[i] = uint8(val / n) //nolint:gosec
		val %= n
	}
}

func main() {
	digits[0] = 1

	for loop := numTerms; loop > 0; loop-- {
		divide(loop)

		digits[0]++
	}

	digitStrs := make([]string, numDigits)
	for loop2 := range numDigits {
		digitStrs[loop2] = onetwothrees[digits[loop2]]
	}

	fmt.Println(strings.Join(digitStrs, ""))
}
