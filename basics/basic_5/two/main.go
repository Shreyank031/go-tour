package main

import "fmt"

type Student struct {
	name   string
	age    int
	grades []int
}

func (s1 *Student) setValues(a string, b int, c []int) {
	s1.age = b
	s1.grades = c
	s1.name = a
}

func (avg Student) studentAverage() float32 {
	sum := 0
	for _, value := range avg.grades {
		sum += value
	}
	return float32(sum) / float32(len(avg.grades))
}

func (max Student) studentMax() int {
	currentMax := 0

	for _, value := range max.grades {
		if currentMax < value {
			currentMax = value
			//test
			//fmt.Println(currentMax)
		}
	}
	return currentMax
}

func main() {
	fmt.Println("Hello, Go World!")

	var s Student = Student{"shrey", 25, []int{78, 68, 49, 79}}
	std := Student{"Bob", 28, []int{30, 40, 59, 88, 92}}
	fmt.Println("Max number std: ", std.studentMax())
	fmt.Println(s)
	fmt.Println(s.age)
	fmt.Println(s.grades)
	fmt.Println("The max number: ", s.studentMax())

	fmt.Println(s.studentAverage())
	s.setValues("shreyank", 790, []int{40, 70, 80, 99})
	fmt.Println(s)

	for i := range 5 {
		println(i)
	}

	fmt.Println(s.studentAverage())
	fmt.Println("The max number: ", s.studentMax())

}
