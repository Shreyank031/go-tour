package main

import (
	"fmt"
	"runtime"
	"time"
)

func main () {
	fmt.Println("This machine runs: ")

	switch OS := runtime.GOOS; OS {
	case "windows":
	fmt.Println("OS windows i.e trash")
	case "linux":
	fmt.Println("OS linux")
	default:
	fmt.Printf("OS %s:\n", OS)
	}
	weekDay()
	time_bro()
}

func weekDay() {
	today := time.Now().Weekday()
	fmt.Println(today)
	fmt.Println(time.Saturday)
	fmt.Println(today+1)

	switch time.Tuesday {
	case today + 5:
	fmt.Println(time.Sunday)
	case today + 5:
	fmt.Println(time.Sunday)
	case today + 1:
	fmt.Println(time.Wednesday)
	case today + 2:
	fmt.Println(time.Thursday)
	case today + 3:
	fmt.Println(time.Friday)
	case today + 4:
	fmt.Println(time.Saturday)
	default:
	fmt.Println("It's Tuesday..Yes it is!")

	}
}

func time_bro() {
	fmt.Println(time.Now().Hour())

	//t:= time.Now().Hour()

	switch {
	case time.Now().Hour() < 12:
	fmt.Println("Gm bro!")
	case time.Now().Hour() > 12: 
	fmt.Printf("GoodAfterNoon I guess, it's %v O'clock\n", time.Now().Hour())
	default:
	fmt.Println("GN bruh!!")
	}

}
