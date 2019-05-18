package main

import "fmt"

func isPerfectSquare(num int) bool {
	for i := 1; i <= num; i++ {
		if i*i == num {
			return true
		} else if i*i > num {
			return false
		}
	}
	return false
}

func main() {
	fmt.Println(isPerfectSquare(1))
	fmt.Println(isPerfectSquare(2))
	fmt.Println(isPerfectSquare(3))
	fmt.Println(isPerfectSquare(4))
	fmt.Println(isPerfectSquare(5))
	fmt.Println(isPerfectSquare(6))
}
