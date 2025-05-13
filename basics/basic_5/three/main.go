package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var x string = fmt.Sprintf("Hello, Go World")
	fmt.Println(x)

	scanner := bufio.NewScanner(os.Stdin) //read from standard input.
	fmt.Print("Type something bruh: ")
	scanner.Scan()          //waits for the user to type and press Enter.
	input := scanner.Text() //gets the typed line as a string and stores it in input.
	fmt.Printf("You typed: %q\n", input)

	fmt.Printf("Type the year you were born: ")
	scanner.Scan()
	input_bro, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		fmt.Println("Oops! Invalid year 😅")
		return
	}
	fmt.Printf("You will be %d years old at the end of 2025\n", 2025-input_bro)

	// var inputt string
	// fmt.Scanln(&inputt)
}
