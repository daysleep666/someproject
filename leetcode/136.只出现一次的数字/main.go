package main

import "fmt"

func singleNumber(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}

func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))
}
