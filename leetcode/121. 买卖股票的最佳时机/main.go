package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	ln := len(prices)
	dp := make([]int, ln) // 这一天卖出时的最大收益
	if ln == 0 {
		return 0
	}
	min := prices[0]
	maxDpi := 0
	for i := 1; i < ln; i++ {
		dp[i] = prices[i] - min
		if prices[i] < min {
			min = prices[i]
		}
		if dp[i] > maxDpi {
			maxDpi = dp[i]
		}
	}
	return maxDpi
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
