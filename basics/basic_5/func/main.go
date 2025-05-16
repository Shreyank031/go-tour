package main

import "fmt"

func test() {
	fmt.Println("it works")
}

func question(x int) {
	fmt.Println("no hope", x)
}

// pass a function to another function
func bob(myFunc func(x int) int) {
	fmt.Println(myFunc(8))
}

func main() {
	test()
	x := question
	x(88)

	timmy := func(ok int) int {
		fmt.Println("function inside a function, ", ok)
		return ok * (-2)
	}(80)
	//fmt.Println(timmy(80))
	fmt.Println(timmy)

	test1 := func(y int) int {
		return y * -1
	}

	bob(test1)

}
