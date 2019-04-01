package main

import (
	"fmt"
	"math"
)

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}
	dp := make([]int, len(triangle)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = math.MaxInt64
	}
	dp[0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += getMin(dp, j)
			fmt.Println(triangle[i][j])
		}
		for j := 0; j < len(triangle[i]); j++ {
			dp[j] = triangle[i][j]
		}
	}
	min := dp[0]
	for i := 0; i < len(dp); i++ {
		if dp[i] < min {
			min = dp[i]
		}
	}
	return min
}

func getMin(arr []int, i int) int {
	min := arr[i]
	if i-1 >= 0 && min > arr[i-1] {
		min = arr[i-1]
	}
	return min
}

func main() {

}
