package main

import (
	"fmt"
	"math"

	"github.com/willf/bitset"
)

/*

Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.

For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284.
The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.

Evaluate the sum of all the amicable numbers under 10000.

*/

func nextPrime(n int, primes *bitset.BitSet) int {
	for i, ln := uint(n), primes.Len(); i < ln; i++ {
		if primes.Test(i) {
			return int(i)
		}
	}

	return 0
}

func sumFactors(number int, primes *bitset.BitSet) int {
	n, sum := number, 1
	var j, p int
	numPrimes := int(primes.Count())

	for i := 0; p*p <= n && n > 1 && i < numPrimes; i++ {
		p = nextPrime(i, primes)
		if n%p == 0 {
			j = p * p
			n /= p
			for n%p == 0 {
				j *= p
				n /= p
			}
			sum *= (j - 1) / (p - 1)
		}
	}

	// a prime factor larger than the square root remains
	if n > 1 {
		sum *= n + 1
	}

	return sum - number
}

func run021() {
	fmt.Print("021: ")
	const limit = 10000
	primes := sieve(uint(math.Sqrt(float64(limit)) + 1))
	var sumAmicable, fi int

	for i := 2; i < limit; i++ {
		fi = sumFactors(i, primes)
		if fi > i && fi < limit {
			if fj := sumFactors(fi, primes); fj == i {
				sumAmicable += (i + fi)
			}
		}
	}

	fmt.Printf("The sum of all amicable numbers under %d is: %d.", limit, sumAmicable)
}
