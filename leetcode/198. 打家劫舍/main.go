package main

import "fmt"

func rob(nums []int) int {
	ln := len(nums)
	if ln == 0 {
		return 0
	} else if ln == 1 {
		return nums[0]
	}
	dp := make([]int, ln+2)
	for i := 2; i < ln+2; i++ {
		if i-3 < 0 {
			dp[i] = dp[i-2] + nums[i-2]
			continue
		}
		dp[i] = max(dp[i-2]+nums[i-2], dp[i-3]+nums[i-3])
	}
	return max(dp[ln], dp[ln+1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
	fmt.Println(rob([]int{}))
	fmt.Println(rob([]int{2}))
	fmt.Println(rob([]int{2, 5}))
	fmt.Println(rob([]int{6, 5}))
}
