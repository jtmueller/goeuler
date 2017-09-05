package main

import (
	"fmt"
	"math/big"
)

/*

The Fibonacci sequence is defined by the recurrence relation:

    Fn = Fn−1 + Fn−2, where F1 = 1 and F2 = 1.

Hence the first 12 terms will be:

    F1 = 1
    F2 = 1
    F3 = 2
    F4 = 3
    F5 = 5
    F6 = 8
    F7 = 13
    F8 = 21
    F9 = 34
    F10 = 55
    F11 = 89
    F12 = 144

The 12th term, F12, is the first term to contain three digits.

What is the index of the first term in the Fibonacci sequence
to contain 1000 digits?

*/

func run025() {
	fmt.Print("025: ")

	i, cnt := 0, 2
	limit := new(big.Int).Exp(big.NewInt(10), big.NewInt(999), nil)
	fib := [3]*big.Int{}

	fib[0] = big.NewInt(1)
	fib[2] = big.NewInt(1)

	for fib[i].Cmp(limit) < 1 {
		i = (i + 1) % 3
		cnt++
		fib[i] = new(big.Int).Add(fib[(i+1)%3], fib[(i+2)%3])
	}

	fmt.Printf("Index of the first 1000-digit fibonacci number: %d", cnt)
}
