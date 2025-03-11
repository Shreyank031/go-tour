package main

import (
	"fmt"
	"math/rand"
	"math"
)

var look, book bool

func main () {
	fmt.Println("Random number between 1 to 100 is, ", rand.Intn(100))
	fmt.Println(math.Pi)

	q :=  add(98, 77)
	fmt.Println(q)

	fmt.Println(stringg("Bob", "has a"))

	j,k := swap("Hello", "World")
	fmt.Println(j, k)

	z1,z2 := 72, 44
	fmt.Println(split(z1,z2))

	var number int
	var string_bro string
	fmt.Println(look, book, number, string_bro, "yes_sir")
}

func split (z1 int, z2 int) (n1 int, n2 int) {
	n1 = z1 - 88 * 2
	n2 = z2 - z1 / 8
	return 
}

func add (a int, b int) int {

	k := a + b
	//return a+b
	return k
}

func stringg (c, d string) (string,string,string) {
	e := "Cute cat"
	return c,d,e
}

func swap (x, y string) (string, string) {
	return y, x
}
