package main

import (
	"fmt"
)

func main() {
	// Defining a struct named 'Vertex' with two fields: x and y
	type Vertex struct {
		x int
		y int
	}

	// Struct literal: Creating an instance of Vertex with values {1, 2}
	fmt.Println(Vertex{1, 2}) // Output: {1 2}

	// Accessing and modifying struct fields
	v := Vertex{x: 7, y: 8}
	v.y = 0 // Modifying the value of field 'y'
	fmt.Println(v) // Output: {7 0}

	// Pointer to struct
	k := &v  // 'k' stores the memory address of 'v'
	k.x = 1e9 // Modifying 'x' through the pointer
	// Since 'k' points to 'v', changes reflect directly on 'v'
	fmt.Printf("%v\n%v\n", v, k) 
	// Output: {1000000000 0}
	// Output: &{1000000000 0}

	// Calling shoot() to demonstrate array usage
	shoot()
}

func shoot() {
	// Array declaration with a fixed size of 2
	var a [2]string
	a[0] = "Hello"  // Assigning value to index 0
	a[1] = "World"  // Assigning value to index 1
	fmt.Println(a) // Output: [Hello World]

	// Array with initialization
	primes := [5]int{1, 2, 3, 4, 5}
	fmt.Println(primes) // Output: [1 2 3 4 5]

	// Array is a fixed-size data structure in Go
	// Unlike slices, the size is part of the type
}

