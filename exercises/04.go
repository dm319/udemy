package main

import "fmt"

func main () {
	var small int
	var large int
	fmt.Printf("%s","Please enter a number: ")
	fmt.Scan(&small)
	fmt.Printf("%s","Please enter a larger number :")
	fmt.Scan(&large)
	answer := large % small
	fmt.Printf("%d\n",answer)
}
