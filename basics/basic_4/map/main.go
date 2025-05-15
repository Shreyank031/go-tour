package main

import "fmt"

func main() {

	var mp map[string]int = map[string]int{
		"apple":        8,
		"banana":       9,
		"coconut":      1,
		"dragon fruit": 11,
	}
	mp["banana"] = 2222
	delete(mp, "coconut")
	fmt.Println(mp)

	fmt.Println(mp["apple"])
	fmt.Println(mp["dragon fruit"])

	val, ok := mp["dragon fruit"]
	val_brush, ok_bruh := mp["coconut"]
	fmt.Println(val, ok, val_brush, ok_bruh)

	ap := map[int]string{
		3: "carrot",
		2: "cucumber",
		1: "radish",
	}
	fmt.Println(ap)

	value, boolean := ap[3]
	fmt.Println(value, boolean)
	value_bruh, boolean_bruh := ap[888]
	fmt.Println(value_bruh, boolean_bruh)

}
