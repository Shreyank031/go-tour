package main //every file with .go extension needs to have a package.

import (
	"fmt" //import standard libraray
	"math/cmplx"
)

var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 - 1
	z complex128 = cmplx.Sqrt(3-4i)
)

func main () {
	fmt.Println("Hello, World!")

	fmt.Printf("Type: %T Value: %v:\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v:\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value %v:\n", z,z)

	fmt.Println(loop(122, 88))

}

func loop (a, b int) (int,int) {

	var sum int
	summ := 2

	for i:=0; i< 11; i++ {
		sum+= (i*a)/b
	}

	for summ<1000 {
		summ += sum
	}


	return sum, summ

}
