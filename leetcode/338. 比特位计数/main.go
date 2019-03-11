package main

import "fmt"

func countBits(num int) []int {
	dp := make([]int, num+1)
	for i := 1; i <= num; i++ {
		dp[i] = dp[i>>1] + (i & 1) // i右移一位加上不知道是不是存在的1(i&1)
	}
	return dp
}

func main() {
	fmt.Println(countBits(2))
	fmt.Println(countBits(5))
	fmt.Println(countBits(8))
}
