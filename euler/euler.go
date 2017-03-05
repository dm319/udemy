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

This a problem on pentagonal numbers.  These numbers are like square numbers, except they represent the concept in pentagons.  The series starts 1,5,12,22...  and is given by the formula in pent().  As it is a quadratic I use function pentagonal(), which solves the square, to check whether a number is pentagonal or not.  The exercise asked you to find a pair of pentagonal numbers where the difference and the sum were also pentagonals.  The function isspecial() checks to see if the difference and the sum are pentagonal, in which case it returns true.

At the end I use a loop within a loop to loop through pentagonal numbers, incrementing the difference between them using the outer loop.

I needed a function that would find the absolute value of an integer - hence the abs() function.  Solving the quadratic I had to convert everything to float64.  I wonder if there is a more elegant way to do this.

*/
