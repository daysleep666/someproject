package main

import "fmt"

func findLength(A []int, B []int) int {
	max, tmp := 0, 0
	dp := make([]int, len(B))

	for i := 0; i < len(A); i++ {
		for j := len(B) - 1; j >= 0; j-- {
			if A[i] == B[j] {
				tmp = 0
				if j-1 >= 0 {
					tmp = dp[j-1]
				}
				dp[j] = tmp + 1
				if dp[j] > max {
					max = dp[j]
				}
			} else {
				dp[j] = 0
			}
		}
	}

	return max
}

func main() {
	a := []int{1, 2, 3, 2, 1}
	b := []int{3, 2, 1, 4, 7}
	fmt.Println(findLength(b, a))
}
