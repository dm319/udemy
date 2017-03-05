package main

import "fmt"

func main() {

	sum := 0

	for i := 0; i < 10; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum = sum + i
			fmt.Printf("%d ", sum)
		}
	}
}
