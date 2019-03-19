package main

import "fmt"

func getPermutation(n int, k int) string {
	// length := compute(n)
	nums := make([]byte, n, n)
	for i := 0; i < n; i++ {
		nums[i] = byte(i+1) + '0'
	}
	strs := d(nums)
	fmt.Println(strs)
	return ""
}

func d(nums []byte) []string {
	if len(nums) == 2 {
		return []string{string(append([]byte{}, nums[0], nums[1])), string(append([]byte{}, nums[1], nums[0]))}
	}
	result := make([]string, 0, compute(len(nums)))
	for i := 0; i < len(nums); i++ {
		cur := string(nums[i])
		new := append([]byte{}, nums[:i]...)
		new = append(new, nums[i+1:]...)
		tmp := d(new)
		for _, v := range tmp {
			result = append(result, cur+v)
		}
	}
	return result
}

func compute(n int) int {
	mul := 1
	for n > 0 {
		mul *= n
		n--
	}
	return mul
}

func main() {
	getPermutation(4, 1)
}
