package main

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}
	m := make(map[int]bool)
	for _, v := range nums1 {
		m[v] = true
	}

	result := []int{}
	for _, v := range nums2 {
		if m[v] {
			result = append(result, v)
			m[v] = false
		}
	}
	return result
}

func main() {

}
