package main

import "fmt"

func hasAlternatingBits2(n int) bool {
	result := n&1 == 1
	for n != 0 {
		n = n >> 1
		if (n&1 == 1) == result {
			return false
		}
		result = !result
	}
	return true
}

func hasAlternatingBits(n int) bool {
	for n != 0 {
		r1 := n & 1
		n = n >> 1
		if n == 0 {
			break
		}
		if r1 == n&1 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(hasAlternatingBits(5))
	fmt.Println(hasAlternatingBits(7))
	fmt.Println(hasAlternatingBits(11))
	fmt.Println(hasAlternatingBits(10))
	fmt.Println(hasAlternatingBits(0))
	fmt.Println(hasAlternatingBits(1))
	fmt.Println(hasAlternatingBits(2))
	fmt.Println(hasAlternatingBits(3))
	fmt.Println(hasAlternatingBits(4))
}
