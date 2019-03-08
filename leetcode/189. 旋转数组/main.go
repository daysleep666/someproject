package main

import "fmt"

func rotate1(nums []int, k int) {
	ln := len(nums)
	for i := 0; i < k; i++ {
		last := nums[ln-1]
		for j := ln - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = last
	}
	fmt.Println(nums)
}

func rotate2(nums []int, k int) {
	ln := len(nums)
	arr := make([]int, ln)
	for i, v := range nums {
		newI := (i + k) % ln
		arr[newI] = v
	}
	copy(nums, arr)
	fmt.Println(nums, "!")
}

func rotate(nums []int, k int) {
	ln := len(nums)
	k = k % ln
	reverse(nums, 0, ln-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, ln-1)
	fmt.Println(nums)
}

func reverse(nums []int, from, to int) {
	for from < to {
		nums[from], nums[to] = nums[to], nums[from]
		from++
		to--
	}
}

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	rotate([]int{-1, -100, 3, 99}, 2)
	rotate([]int{-1}, 2)
}
