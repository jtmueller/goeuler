package main

import (
	"fmt"
)

/*

A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
a**2 + b**2 = c**2

For example, 3**2 + 4**2 = 9 + 16 = 25 = 5**2.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.

*/

func run009() {
	fmt.Print("009: ")

	a, b, c := 0, 0, 0
	s := 1000
	found := false

	for a = 1; a < s/3; a++ {
		for b = a; b < s/2; b++ {
			c = s - a - b
			if a*a+b*b == c*c {
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	fmt.Printf("Triple is %d, %d, %d. Product: %d", a, b, c, a*b*c)
}
