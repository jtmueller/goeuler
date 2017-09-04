package main

import (
	"fmt"
	"math/big"
)

/*

2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2^1000?

*/

func run016() {
	fmt.Print("016: ")

	n := new(big.Int).Exp(big.NewInt(2), big.NewInt(1000), nil)
	ns := n.Text(10)

	var sum int
	for _, d := range ns {
		sum += int(d - '0')
	}

	fmt.Printf("Sum of all digits in 2^1000 = %d", sum)
}
