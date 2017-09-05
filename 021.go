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

func sumFactors(number uint, primes *bitset.BitSet, numPrimes uint) uint {
	n, sum := number, uint(1)
	var j, p uint

	for i := uint(0); n > 1 && i < numPrimes && p*p <= n; i++ {
		p, ok := primes.NextSet(i)
		if !ok {
			break
		}
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
	const limit uint = 10000
	primes := sieve(uint(math.Sqrt(float64(limit)) + 1))
	numPrimes := primes.Count()
	var sumAmicable, fi uint

	for i := uint(2); i < limit; i++ {
		fi = sumFactors(i, primes, numPrimes)
		if fi > i && fi < limit {
			if fj := sumFactors(fi, primes, numPrimes); fj == i {
				sumAmicable += (i + fi)
			}
		}
	}

	fmt.Printf("The sum of all amicable numbers under %d is: %d.", limit, sumAmicable)
}
