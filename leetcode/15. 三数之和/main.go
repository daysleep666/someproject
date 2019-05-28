package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)

	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if m[nums[i]] {
			continue
		}
		m[nums[i]] = true
		tmp := find(nums, nums[i], i)
		result = append(result, tmp...)
		// break
	}
	return result
}

func find(num []int, dst int, remain int) [][]int {
	var result [][]int
	for i, j := remain+1, len(num)-1; i < j; {
		if i > remain+1 && num[i] == num[i-1] {
			i++
			continue
		}
		if j < len(num)-1 && num[j] == num[j+1] {
			j--
			continue
		}
		if num[i]+num[j]+dst == 0 {
			a, b, c := so(num[i], num[j], dst)
			// fmt.Println(a,b,c)
			result = append(result, []int{a, b, c})
			i++
		} else if num[i]+num[j]+dst > 0 {
			j--
		} else {
			i++
		}
	}
	return result
}

func so(a, b, c int) (int, int, int) {
	if a > b {
		a, b = b, a
	}
	if a > c {
		a, c = c, a
	}
	if b > c {
		b, c = c, b
	}
	return a, b, c
}

func main() {
	// fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	// fmt.Println(threeSum([]int{-1, 0, 1}))
	// fmt.Println(threeSum([]int{0, 0, 0, 0}))
	// fmt.Println(threeSum([]int{}))
	// fmt.Println(threeSum([]int{-2, 0, 1, 1, 2}))
	fmt.Println(threeSum([]int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0}))

}
