package main

import (
	"fmt"
	"strconv"
)

/*
A palindromic number reads the same both ways. The largest palindrome made
from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/

func reverse(s string) string {
	r := []rune(s)
	ln := len(r)
	for i, j := 0, ln-1; i < ln/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func makePalindrome(firstHalf int) int {
	fh := strconv.Itoa(firstHalf)
	sh := reverse(fh)
	out, _ := strconv.Atoi(fh + sh)
	return out
}

func run004() {
	fmt.Printf("004: ")

	found := false
	firstHalf, palin := 998, 0
	var factors [2]int

	for !found {
		firstHalf--
		palin = makePalindrome(firstHalf)

		for i := 999; i*i > palin; i-- {
			if palin/i > 999 {
				break
			}

			if palin%i == 0 {
				found = true
				factors[0] = palin / i
				factors[1] = i
				break
			}
		}
	}
	fmt.Printf("Found the palindrome %d, which is made from %d * %d.", palin, factors[0], factors[1])
}
