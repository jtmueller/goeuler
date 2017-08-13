package main

import (
	"fmt"
	"math"

	"github.com/willf/bitset"
)

/*

The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.

*/

func sieve(limit uint) *bitset.BitSet {
	primes := bitset.New(limit)
	var i, j uint

	// Mark everything prime to start
	for i = 2; i < limit; i++ {
		primes.Set(i)
	}

	// Only need to sieve up to sqrt(limit)
	imax := uint(math.Sqrt(float64(limit))) + 1

	for i = 2; i < imax; {
		j = i * i
		for j < limit {
			primes.Clear(j)
			j += i
		}

		// move i to next prime
		for {
			if i++; primes.Test(i) {
				break
			}
		}
	}

	return primes
}

func sum(nums *bitset.BitSet) uint {
	var x, i uint
	for ln := nums.Len(); i < ln; i++ {
		if nums.Test(i) {
			x += i
		}
	}
	return x
}

func run010() {
	fmt.Print("010: ")

	fmt.Printf("The sum of all the primes below two million is %d.", sum(sieve(2000000)))
}
