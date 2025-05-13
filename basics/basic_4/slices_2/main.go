package main

import (
	"fmt"
)

func main () {
	a := make([]int, 8)
	printSlice("a", a)

	b := make([]int, 3, 8)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:4]
	printSlice("d", d)

}

func printSlice (str string, r []int) {
	fmt.Printf("%s size: %d capacity: %d %v\n", str, len(r), cap(r), r)
}
