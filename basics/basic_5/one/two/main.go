package main

import "fmt"

type Employee struct {
	f_name, l_name string
	age            int
}

func main() {

	e1 := Employee{f_name: "rachel", l_name: "chas", age: 19}
	e2 := &e1
	fmt.Println(e2.age)
	e2.age = 20

	fmt.Println("before: ", e1)
	fmt.Println("address: ", e2)
	fmt.Println(e2.age)

}
