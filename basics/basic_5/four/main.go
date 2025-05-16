package main

import (
	"flag"
	"fmt"
)

func changeFirst(slice []int) {
	slice[0] = 999
}

func main() {
	name := flag.String("name", "Guest", "Your Name bruh")
	age := flag.Int("age", 99, "Your age bruh")
	debug := flag.Bool("debug", false, "Enable Debug")

	// Parse the flags from command line
	flag.Parse()

	fmt.Printf("Hello, %s!\n", *name)
	fmt.Printf("You are %d years old\n", *age)
	fmt.Printf("Debug mode: %v\n", *debug)

	fmt.Printf("\n----------\n\n")

	//mutable and immutable data typesane
	var x []int = []int{1, 2, 3}
	y := x
	y[0] = 999
	fmt.Println(x, y)

	var k [3]int = [3]int{1, 2, 3}
	l := k
	l[0] = 999
	fmt.Println(k, l)

	a := map[string]int{
		"Hello": 1,
		"World": 2,
		"Go":    3,
	}
	b := a
	b["Go"] = 999

	fmt.Println(a, b)

	s := []int{1, 2, 3}
	fmt.Println(s)
	changeFirst(s)
	fmt.Println(s)

}
