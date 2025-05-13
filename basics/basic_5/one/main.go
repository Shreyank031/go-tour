package main

import (
	"fmt"
)

type Employee struct {
	first_name, last_name string
	age, salary           int
}

func main() {
	emp1 := Employee{
		first_name: "James",
		last_name:  "Kava",
		age:        22,
		salary:     10,
	}

	fmt.Println("emp1: ", emp1)

	emp2 := Employee{
		"David", "Tutu", 70, 20,
	}
	fmt.Println("emp2: ", emp2)

	// anonymous struct
	emp3 := struct {
		f_name, l_name string
		age, salary    int
	}{
		f_name: "Bob",
		l_name: "Jackson",
		age:    18,
		salary: 1,
	}
	fmt.Println("emp3", emp3)

	fmt.Println("f_name: ", emp3.f_name)
	fmt.Println("l_name: ", emp3.l_name)
	fmt.Printf("age: $%d", emp3.age)
	emp3.age = 17
	fmt.Println("New age", emp3.age)

}
