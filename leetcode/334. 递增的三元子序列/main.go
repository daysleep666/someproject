package main

import "fmt"

func increasingTriplet(nums []int) bool {
	ln := len(nums)
	if ln < 3 {
		return false
	}
	a, b, min := nums[0], nums[1], nums[0]
	for i := 1; i < ln; i++ {
		if a >= b {
			a = b
			b = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		} else if nums[i] > min && nums[i] <= b {
			a = min
			b = nums[i]
		}
		if is(a, b, nums[i]) {
			fmt.Println(a, b, nums[i])
			return true
		}

	}
	return false
}

func is(a, b, c int) bool {
	if a < b && b < c {
		return true
	}
	return false
}

func main() {
	// fmt.Println(increasingTriplet([]int{1, 2, 3, 4, 5}))
	// fmt.Println(increasingTriplet([]int{5, 4, 3, 2, 1}))
	// fmt.Println(increasingTriplet([]int{5, 6, 0, 8, 7}))
	fmt.Println(increasingTriplet([]int{5, 0, 0, 10, 0, 100}))
}
