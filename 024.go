package main

import (
	"fmt"
	"strconv"
)

/*

A permutation is an ordered arrangement of objects. For example, 3124
is one possible permutation of the digits 1, 2, 3 and 4. If all of the
permutations are listed numerically or alphabetically, we call it
lexicographic order. The lexicographic permutations of 0, 1 and 2 are:

012   021   102   120   201   210

What is the millionth lexicographic permutation of
the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?

*/

func fact(i int) int {
	if i <= 0 {
		return 0
	}

	p := 1
	for j := 1; j <= i; j++ {
		p *= j
	}

	return p
}

func run024() {
	fmt.Print("024: ")

	digits := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	remain := 1000000 - 1
	var permNum string

	for i := 1; i < 10 && remain > 0; i++ {
		f := fact(10 - i)
		j := remain / f
		remain = remain % f
		permNum += strconv.Itoa(digits[j])
		digits = append(digits[:j], digits[j+1:]...)
	}

	for i := 0; i < len(digits); i++ {
		permNum += strconv.Itoa(digits[i])
	}

	fmt.Printf("THe millionth permutation is: %s", permNum)
}
