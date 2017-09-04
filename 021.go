package main

import (
	"fmt"
	"math"
)

/*

Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.

For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284.
The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.

Evaluate the sum of all the amicable numbers under 10000.

*/

func sumFactors(n int) int {
	sqrt := int(math.Sqrt(float64(n)))
	sum := 1

	// If the number is a perfect square, count the sqrt once in the
	// sum of factors.
	if n == sqrt*sqrt {
		sum += sqrt
		sqrt--
	}

	for i := 2; i <= sqrt; i++ {
		if n%i == 0 {
			sum += (i + (n / i))
		}
	}

	return sum
}

func run021() {
	fmt.Print("021: ")
	const limit = 10000
	var sumAmicable, fi int

	for i := 2; i < limit; i++ {
		fi = sumFactors(i)
		if fi > i && fi < limit {
			if fj := sumFactors(fi); fj == i {
				sumAmicable += (i + fi)
			}
		}
	}

	fmt.Printf("The sum of all amicable numbers under %d is: %d.", limit, sumAmicable)
}
