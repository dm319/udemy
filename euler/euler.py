import math

def pent(n):
    p = n*(3*n-1)/2
    return(p)

def pentagonal(p):
    a = 3.0/2.0
    b = -(1.0/2.0)
    c = -(p)

    x = (-b + math.sqrt(b*b-4*a*c)) / (2*a)

    return(x == int(x))

def isspecial(x,y):
    a = pent(x)
    b = pent(y)

    diff = abs(a - b)

    difftest = pentagonal(diff)
    sumtest = pentagonal(a+b)

    return(difftest and sumtest)

n = 2000

for i in range(1,n):
    for j in range(1,n):
        result = isspecial(j,i+j)
        if result:
            print("P%d(%d)\tP%d(%d)\n") % (j, pent(j), i, pent(i))

print("end")
