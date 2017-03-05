package main

import "fmt"

func main () {
	var name string
	fmt.Printf("%s","Please enter your name: ")
	fmt.Scan(&name)
	fmt.Printf("%s%s\n","Hello ",name)
}
