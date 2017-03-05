package main

import (
	"fmt"
	"math"
)

func pent(n int) int {
	p := n * (3*n - 1) / 2
	return (p)
}

func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}

func pentagonal(p int) bool {
	a := 3.0 / 2.0
	b := -(1.0 / 2.0)
	c := -float64(p)

	x := (-b + math.Sqrt(b*b-4*a*c)) / (2 * a)

	if x == float64(int(x)) {
		return true
	} else {
		return false
	}
}

func isspecial(x, y int) bool {
	a := pent(x)
	b := pent(y)

	diff := abs(a - b)

	difftest := pentagonal(diff)
	sumtest := pentagonal(a + b)

	if difftest && sumtest {
		return true
	} else {
		return false
	}
}

func main() {
	n := 2000 //10000 max

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			a := j
			b := a + i
			result := isspecial(a, b)
			if result {
				fmt.Printf("P%d(%d)\t P%d(%d)\n", a, b, pent(a), pent(b))
				fmt.Printf("difference:%d\n",abs(pent(a)-pent(b)))
			}
		}
	}
	print("end\n")
}

/*

Pentagonal numbers are like square numbers but with pentagons.  The exercise on project euler asked if you could find the pair of pentagonal numbers where the difference was also pentagonal, and so was the sum.

The function pent() generates the pentagonal number series, i.e. P2 is 5 (the second pentagonal number).

Function pentagonal() checks whether the number is pentagonal or not.  I had to convert to float64 - not sure if this is very elegant, maybe their is a better way, but I needed the square root.

Function isspecial() takes two numbers which it converts to the pentagonal series, then checks the difference and sum to see if they are pentagonal, if they are it returns true.

the main function has a loop within a loop to generate numbers 1-2000 with differences of 1-2000 - sending these to isspecial, and only printing when it finds a match.

*/
