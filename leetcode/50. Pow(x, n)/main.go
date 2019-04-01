package main

import "fmt"

func myPow(x float64, n int) float64 {
	curn := abs(n)
	v := d(x, curn)
	if n < 0 {
		v = 1 / v
	}
	return v
}

func d(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else if n == 1 {
		return x
	}
	v := d(x*x, n/2)
	if n%2 != 0 {
		v = v * x
	}
	return v
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	fmt.Println(myPow(2, 10))
	fmt.Println(myPow(2.1, 3))
	fmt.Println(myPow(2, -2))
	fmt.Println(myPow(-1, 2147483647))
}
