package main

import "fmt"

func main() {
	var arr [8]string
	arr[0] = "shreyank"
	arr[3] = "bob"
	arr[7] = "Timmy"
	var num [8]int
	num[0] = 99
	num[1] = 89
	num[2] = 79
	num[7] = 59
	fmt.Println(arr)
	fmt.Println(num)

	array := []int{}
	array = append(array, 22, 33, 55, 88, 99, 102, 200)
	fmt.Println(array)

	sum := 0
	for i := range array {
		sum += array[i]
	}
	fmt.Println(sum)

	var slice []int = []int{1, 2, 3, 4, 5, 9}
	fmt.Println(slice)
	fmt.Println(cap(slice))
	slice = append(slice, 10, 11, 12)
	fmt.Println(slice)
	fmt.Println(cap(slice))
	slice_slice := slice[2:4]
	fmt.Println(slice_slice)
	fmt.Println(cap(slice_slice))

	a_make := make([]int, 8)
	b_make := make([]string, 8)
	fmt.Println(a_make)
	fmt.Println(b_make)

}
