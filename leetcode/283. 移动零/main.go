package main

import "fmt"

func moveZeroes(nums []int) {
	ln := len(nums)

	j := 0
	for i := 0; i < ln; i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	fmt.Println(nums)
}

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
}
