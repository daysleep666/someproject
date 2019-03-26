package main

import "fmt"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	start, end, cur := nums[0], nums[0], nums[0]
	result := make([]string, 0)
	for i := 1; i < len(nums); i++ {
		v := nums[i]
		if v == cur+1 {
			cur = cur + 1
			end = v
		} else {
			str := ""
			if start == end {
				str = fmt.Sprintf("%v", start)
			} else {
				str = fmt.Sprintf("%v->%v", start, end)
			}
			result = append(result, str)
			start, end, cur = v, v, v
		}
	}
	str := ""
	if start == end {
		str = fmt.Sprintf("%v", start)
	} else {
		str = fmt.Sprintf("%v->%v", start, end)
	}
	result = append(result, str)
	return result
}

func main() {
	fmt.Println(summaryRanges([]int{1, 2, 3, 4}))
	fmt.Println(summaryRanges([]int{}))
	fmt.Println(summaryRanges([]int{1, 3}))
	fmt.Println(summaryRanges([]int{1, 2, 3, 5}))
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
}
