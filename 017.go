package main

import (
	"fmt"
)

/*

If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.
If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?

*/

func run017() {
	fmt.Print("017: ")

	var letterCount int
	for i := 1; i <= 1000; i++ {
		letterCount += WordLetterCount(i)
	}

	fmt.Printf("The number of letters in the English version of 1-1000 is %d.", letterCount)
}
