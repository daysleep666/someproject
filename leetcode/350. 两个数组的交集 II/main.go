package main

import (
	"fmt"
	"sort"
)

func intersect1(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	l1, l2 := len(nums1), len(nums2)
	result := make([]int, 0)
	for i, j := 0, 0; i < l1 && j < l2; {
		if nums1[i] == nums2[j] {
			result = append(result, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return result
}

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for _, v := range nums1 {
		m[v]++
	}
	result := make([]int, 0)
	for _, v := range nums2 {
		if m[v] > 0 {
			result = append(result, v)
			m[v]--
		}
	}
	return result
}

func main() {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2}))
}
