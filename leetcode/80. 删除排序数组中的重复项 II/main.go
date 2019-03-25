package main

import "fmt"

func removeDuplicates1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	cur, count := nums[0], 1
	for i := 1; i < len(nums); {
		if nums[i] == cur {
			if count == 2 { // 这个不能要
				nums = append(nums[:i], nums[i+1:]...)
				continue
			}
			count++
		} else {
			cur = nums[i]
			count = 1
		}
		i++
	}
	return len(nums)
}

func removeDuplicates(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if i < 2 || nums[i] != nums[count-2] {
			nums[count] = nums[i]
			count++
		}
	}
	return count
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	count := removeDuplicates(nums)
	fmt.Println(nums, count)
}
