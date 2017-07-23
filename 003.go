package main

import (
	"fmt"
)

/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/

func isPrime(n int) bool {
	switch {
	case n <= 1:
		return false
	case n == 2 || n == 3:
		return true
	case n%2 == 0 || n%3 == 0:
		return false
	default:
		i, w := 5, 2
		for i*i <= n {
			if n%i == 0 {
				return false
			}
			i += w
			w = 6 - w
		}
		return true
	}
}

func largestPrimeFactor(n uint64) uint64 {
	x := n
	var fact uint64

	for i := uint64(2); i*i <= x; {
		switch {
		case x%i == 0:
			x = x / i
			fact = i
		case i == 2:
			i = 3
		default:
			i += 2
		}
	}
	if x > fact {
		// the remainder is a prime number
		fact = x
	}

	return fact
}

func run003() {
	fmt.Print("003: ")
	fmt.Println(largestPrimeFactor(600851475143))
}
