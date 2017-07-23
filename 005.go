package main

import (
	"fmt"
	"math"
)

/*

2520 is the smallest number that can be divided by each of the numbers
from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of
the numbers from 1 to 20?

*/

func generatePrimes(upperLimit int) []int {
	primes := []int{2}
	for i := 3; i < upperLimit; i += 2 {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func run005() {
	fmt.Print("005: ")

	divisorMax, result := 20, 1
	ps := generatePrimes(divisorMax)

	for _, p := range ps {
		a := math.Floor(math.Log(float64(divisorMax)) / math.Log(float64(p)))
		result = result * int(math.Pow(float64(p), a))
	}

	fmt.Printf("Smallest divisible by all numbers in the range 1-20 is %d.", result)
}
