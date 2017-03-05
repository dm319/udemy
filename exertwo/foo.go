package main

import "fmt"

func foo(b... int) {
	var sum int
	if true {
		for _,j := range b {
			sum =+ j
		}
	} else {
		sum = 0
	}
	fmt.Println("The sum is: ", sum)
}

func main() {
	foo(1,2)
	foo(1,2,3)
	aSlice := []int{1,2,3,4}
	fmt.Printf("%T\n", aSlice)
	foo(aSlice...)
	foo()
}

