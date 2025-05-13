package main

import "fmt"

func main() {
	var x int = 55

	switch x {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 5, 55:
		fmt.Println("five")
		fmt.Println("fifty_five")

	default:
		fmt.Println("I give up")
	}

	switch {
	case x < 5:
		fmt.Println("it's 5_yoo")

	case x == 6, x > 55:
		fmt.Println("you got me: 5")
	default:
		fmt.Println("No hope, hope")
	}

}
