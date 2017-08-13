package main

import (
	"fmt"
)

/*

Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down, there are exactly 6 routes to the bottom right corner.

How many such routes are there through a 20×20 grid?

*/

func run015() {
	fmt.Print("015: ")

	const gridSize int = 20
	paths := 1

	for i := 0; i < gridSize; i++ {
		paths *= (2 * gridSize) - i
		paths /= i + 1
	}

	fmt.Printf("In a 20x20 grid there are %d possible paths.", paths)
}
