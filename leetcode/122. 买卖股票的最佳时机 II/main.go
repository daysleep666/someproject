package main

func maxProfit(prices []int) int {
	result := 0
	dp := make([]int, len(prices)) //
	for k := 0; k < len(prices); k++ {
		dp[k] = 0 // 从这天开始的
		min := prices[k]
		max := 0
		for i := k + 1; i < len(prices); i++ {
			if prices[i] > min {
				dp[i] = prices[i] - min
			} else if prices[i] < min {
				min = prices[i]
			}
			if dp[i] > max {
				max = dp[i]
				k = i + 1
			}
			result += max
		}
	}
	return result
}

func main() {

}
