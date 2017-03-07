pent <- function(n) {
	p <- n*(3*n-1)/2
	return(p)
}

pentagonal <- function(p) {
	a <- 3/2
	b <- -(1/2)
	c <- -p

	x <- (-b + sqrt(b*b-4*a*c))/(2*a)

	return(x == floor(x))
}

isspecial <- function(x,y) {
	a <- pent(x)
	b <- pent(y)

	diff <- abs(a-b)

	difftest <- pentagonal(diff)
	sumtest <- pentagonal(a+b)

	return(difftest & sumtest)
}

n <- 2000

for (i in 1:n) {
	for (j in 1:n) {
		result <- isspecial(j,i+j)
		if(result) {
			print(cat(c("P",j,"(",pent(j),"), P",i,"(",pent(i),")"),sep = ""))
		}
	}
}

print("end")


