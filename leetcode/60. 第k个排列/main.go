package main

import "fmt"

func getPermutation(n int, k int) string {
	nums := make([]byte, n, n)
	for i := 0; i < n; i++ {
		nums[i] = byte(i+1) + '0'
	}

	result := ""
	for i := n - 1; i >= 0; i-- {
		index := k / compute(i)
		fmt.Println(k, compute(i), index)
		result += string(nums[index])
		nums = append(nums[:index], nums[index+1:]...)
		k = k - compute(i)
	}

	fmt.Println(string(nums))
	return string(nums)
}

func compute(n int) int {
	mul := 1
	for n > 0 {
		mul *= n
		n--
	}
	return mul
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

func main() {
	// getPermutation(4, 9)
	getPermutation(4, 5)
	// getPermutation(4, 6)
	// getPermutation(4, 24)
	// getPermutation(3, 1)
	// getPermutation(3, 3)
	// getPermutation(3, 4)
}
