package main

import (
	"fmt"
)

func main() {
	// Declaring two integer variables
	var i, j int = 72, 88

	// Printing initial values of i and j
	fmt.Println(i, j) // Output: 72 88

	// Declaring pointer variables that store the memory addresses of i and j
	p, r := &i, &j
	fmt.Println(p, r) // Prints the memory addresses of i and j

	// Dereferencing pointers and modifying the values stored at those addresses
	*p = 44 // Modifies i (since p points to i)
	*r = 11 // Modifies j (since r points to j)

	// Printing modified values of i and j
	fmt.Println(i, j) // Output: 44 11

	// Changing p to point to j instead of i
	p = &j
	*p = *p / 3 // j = j / 3 (integer division)

	// Printing final values of i and j
	fmt.Println(i, j) // Output: 44 3
}

