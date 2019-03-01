package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	var new = make([]int, 0, m+n)
	i, j := 0, 0
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			new = append(new, nums1[i])
			i++
		} else {
			new = append(new, nums2[j])
			j++
		}
	}
	for ; i < m; i++ {
		new = append(new, nums1[i])
	}
	for ; j < n; j++ {
		new = append(new, nums2[j])
	}

	if (m+n)%2 == 0 {

		return (float64(new[(m+n)/2]) + float64(new[(m+n)/2-1])) / 2
	}
	return float64(float64(new[(m+n)/2]))
}

func main() {
	num1 := []int{2}
	num2 := []int{}
	v := findMedianSortedArrays(num1, num2)
	fmt.Println(v)
}
