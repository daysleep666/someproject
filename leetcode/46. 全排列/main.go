package main

import "fmt"

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	} else if len(nums) == 1 {
		return [][]int{[]int{nums[0]}}
	}

	result := make([][]int, 0)
	for i := 1; i < len(nums); i++ {
		cur := nums[:i]
		if i < len(nums) {
			tmpReturn := permute(nums[i:])
			for _, v := range tmpReturn {
				for j := 0; j < len(v)+1; j++ {
					tmp := make([]int, 0, len(v)+len(cur))
					tmp = append(tmp, v[:j]...)
					tmp = append(tmp, cur...)
					tmp = append(tmp, v[j:]...)
					if !find(result, tmp) {
						result = append(result, tmp)
					}
				}
			}
		}
	}
	return result
}

func find(result [][]int, tmp []int) bool {
	for _, v := range result {
		i := 0
		for _, vv := range v {
			if vv != tmp[i] {
				break
			}
			i++
		}
		if i == len(tmp) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(permute([]int{1, 2}))
}
