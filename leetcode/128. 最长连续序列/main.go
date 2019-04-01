package main

import "fmt"

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	//找出max
	max, min := nums[0], nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	offset := -min

	// 创建足够大的数组
	arr := make([]int, offset+max+1)

	// 计数
	for _, v := range nums {
		arr[v+offset]++
	}

	maxLength, curLength := 0, 0
	// 遍历
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			curLength++
			if curLength > maxLength {
				maxLength = curLength
			}
		} else {
			curLength = 0
		}
	}
	return maxLength
}

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutive([]int{1}))
	fmt.Println(longestConsecutive([]int{0, -1}))
	fmt.Println(longestConsecutive([]int{0, -3}))
}
