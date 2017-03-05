package main

import "fmt"

func divide(i int) (int, bool) {
	a := i / 2
	var b bool
	if i%2 == 0 {
		b = true
	} else {
		b = false
	}
	return a, b
}

func main() {
	var get int
	fmt.Scanln(&get)
	fmt.Println(divide(get))
}
