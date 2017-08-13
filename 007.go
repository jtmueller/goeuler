package main

import (
	"fmt"
)

/*

By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13,
we can see that the 6th prime is 13.

What is the 10,001st prime number?

*/

func nthPrime(n uint) uint {
	var i, count uint
	limit := n * 12
	primes := sieve(limit)

	for i = 2; i < limit; i++ {
		if primes.Test(i) {
			count++
		}
		if count == n {
			return i
		}
	}

	return 0
}

func run007() {
	fmt.Print("007: ")
	prime := nthPrime(10001)
	fmt.Printf("10,001st prime: %d", prime)
}
