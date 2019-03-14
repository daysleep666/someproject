package main

import "fmt"

func isPowerOfTwo(n int) bool {
	return n&(n-1) == 0 && n != 0
}

func main() {
	fmt.Println(isPowerOfTwo(2))
	fmt.Println(isPowerOfTwo(3))
	fmt.Println(isPowerOfTwo(4))
	fmt.Println(isPowerOfTwo(5))
	fmt.Println(isPowerOfTwo(8))
	fmt.Println(isPowerOfTwo(0))
	fmt.Println(isPowerOfTwo(-4))
}
