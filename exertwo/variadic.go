package main

import "fmt"

func max (a ...int) (int,int) {
	var m int
	var n int
	for i, z := range a {
		fmt.Println("i and z =",i,z)
		if m < z {
			n, m = i, z
		}
	}
	return n, m
}

func main() {
	n, m := max(12,19,15,16,17)
	fmt.Println(n,m)
}


