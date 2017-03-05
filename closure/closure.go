package main

import "fmt"

func wrapper(y int) func() int {
	x := 0
	return func() int {
		x = x + y
		return x
	}
}

func main() {
	incrementtwo := wrapper(2)
	incrementthree := wrapper(3)
	fmt.Println(incrementtwo())
	fmt.Println(incrementtwo())
	fmt.Println(incrementthree())
	fmt.Println(incrementthree())
	fmt.Println(incrementtwo())
	fmt.Println(incrementthree())
}
