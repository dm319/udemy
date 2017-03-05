package main

import "fmt"

func hello() {
	fmt.Print("Hello ")
}

func world() {
	fmt.Print("World")
}

func again() {
	fmt.Println(" Again")
}

func main() {
	defer again()
	defer world()
	hello()
}

