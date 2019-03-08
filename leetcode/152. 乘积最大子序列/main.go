package main

import "fmt"

func maxProduct1(nums []int) int {
	ln := len(nums)
	if ln == 0 {
		return 0
	}
	max := nums[0]
	for k := 0; k < ln; k++ {
		cur := nums[k]
		if max < cur {
			max = cur
		}
		for i := k + 1; i < ln; i++ {
			cur *= nums[i]
			if max < cur {
				max = cur
			}
		}
	}
	return max
}

func maxProduct(nums []int) int {
	ln := len(nums)
	if ln == 0 {
		return 0
	}
	res, max, min := nums[0], nums[0], nums[0]
	for i := 1; i < ln; i++ {
		if nums[i] < 0 {
			max, min = min, max
		}
		max = getmax(max*nums[i], nums[i])
		min = getmin(min*nums[i], nums[i])
		res = getmax(max, res)
	}
	return res
}

func getmax(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func getmin(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	fmt.Println(maxProduct([]int{-2, 0, -1}))
	fmt.Println(maxProduct([]int{1, -2, 3, -4}))
	fmt.Println(maxProduct([]int{0, 2}))
	fmt.Println(maxProduct([]int{-2}))
}
