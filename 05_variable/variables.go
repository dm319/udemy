package main

import "fmt"

var a int
var b string
var c float64
var d bool

func main() {
	a = 23
	b = "golang"
	c = 4.17
	d = true

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)

}
