package main

import (
	"fmt"
)

func subsets(nums []int) [][]int {
	return append([][]int{[]int{}}, d(nums)...)
}

func d(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	} else if len(nums) == 1 {
		return [][]int{[]int{nums[0]}}
	}
	sub := make([][]int, 0)
	for i := 0; i < len(nums)+1; i++ {
		tmpResult := nums[:i]

		if len(tmpResult) > 0 {
			sub = append(sub, tmpResult)
		}
		if i+1 <= len(nums) {
			tmpReturn := d(nums[i+1:])
			for _, tr := range tmpReturn {
				sub = append(sub, make([]int, 0, len(tmpResult)+len(tr)))
				sub[len(sub)-1] = append(sub[len(sub)-1], tmpResult...)
				sub[len(sub)-1] = append(sub[len(sub)-1], tr...)
			}
		}
	}
	return sub
}

func main() {
	a := []int{1, 2, 3}
	b := []int{}
	b = append(b, a...)
	a[0] = 999
	fmt.Println(a, b)
	// fmt.Println(subsets([]int{5, 2, 3, 4, 1}))
}
