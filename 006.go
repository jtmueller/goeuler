package main

import (
	"fmt"
)

/*

The sum of the squares of the first ten natural numbers is,
1**2 + 2**2 + ... + 10**2 = 385

The square of the sum of the first ten natural numbers is,
(1 + 2 + ... + 10)**2 = 55**2 = 3025

Hence the difference between the sum of the squares of the first ten natural
numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred
natural numbers and the square of the sum

*/

func sumSquares(limit int) int {
	var sum int
	for i := 1; i <= limit; i++ {
		sum += i * i
	}
	return sum
}

func squareSum(limit int) int {
	var sum int
	for i := 1; i <= limit; i++ {
		sum += i
	}
	return sum * sum
}

func run006() {
	fmt.Print("006: ")
	diff := squareSum(100) - sumSquares(100)
	fmt.Printf("The difference between the sum of the squares and the square of the sums of the first 100 numbers is: %d", diff)
}
