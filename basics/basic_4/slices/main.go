
package main

import (
	"fmt"
)

func main() {
	// ✅ Basic Slice Operations
	// Slice of strings
	slice := []string{"Jemmy", "Oliver", "Mickey", "Crimson"}
	fmt.Println(slice) // Output: [Jemmy Oliver Mickey Crimson]

	// Slicing from index 1 (removes "Jemmy")
	slice = slice[1:]
	fmt.Println(slice) // Output: [Oliver Mickey Crimson]

	// Slicing up to index 2 (keeps first two elements)
	slice = slice[:2]
	fmt.Println(slice) // Output: [Oliver Mickey]

	// Slicing from index 1 (removes "Oliver")
	slice = slice[1:]
	fmt.Println(slice) // Output: [Mickey]

	// Calling function to demonstrate slice of struct
	arrayStruct()
}

func arrayStruct() {
	// ✅ Slice of Structs
	as := []struct {
		i int
		j bool
	}{
		{1, true},
		{2, false},
		{3, false},
		{4, true},
	}
	fmt.Println(as) // Output: [{1 true} {2 false} {3 false} {4 true}]

	// ✅ Slice of Integers
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(s) // Output: len=8 capacity=8 [1 2 3 4 5 6 7 8]

	// Slicing to length 0 (empty slice, but retains capacity)
	s = s[:0]
	printSlice(s) // Output: len=0 capacity=8 []

	// Extending the slice to length 4 (reuses underlying array)
	s = s[:4]
	printSlice(s) // Output: len=4 capacity=8 [1 2 3 4]

	// Slicing from index 2 to the end (removes first two elements)
	s = s[2:]
	printSlice(s) // Output: len=2 capacity=6 [3 4]

	// ✅ Declaring an empty slice (nil slice)
	var arr []int
	printSlice(arr) // Output: len=0 capacity=0 []

	// Check if the slice is nil (default zero value of a slice is nil)
	if arr == nil {
		fmt.Println("nil") // Output: nil
	}
}

func printSlice(s []int) {
	// Prints the length, capacity, and elements of a slice
	fmt.Printf("len=%d capacity=%d %v\n", len(s), cap(s), s)
}

