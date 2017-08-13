package main

import (
	"fmt"
)

/*

The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:
13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1

It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.

*/

func collatz(n int) int {
	if n%2 == 0 {
		return n / 2
	}
	return 3*n + 1
}

func run014() {
	fmt.Print("014: ")

	target := 1000000
	cache := make([]int, target+1)
	cache[1] = 1

	var maxLen, start int
	for i := 2; i < target; i++ {
		var n, k int
		for n = i; n > 1 && n >= i; {
			n = collatz(n)
			k++
		}
		ln := k + cache[n]
		cache[i] = ln

		if ln > maxLen {
			maxLen = ln
			start = i
		}
	}

	fmt.Printf("Collatz starting number with longest chain: %d", start)
}
