package main

import "fmt"

func main () {
	for i := 0; i <= 100; i++ {
		if i % 2 == 0 {
			fmt.Printf("%d",i)
			if i < 100 {
				fmt.Printf(", ")
			} else {
				fmt.Printf("\n")
			}
		}
	}
}
