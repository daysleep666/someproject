package main

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := make([][]int, 0)
	front := intervals[0][0]
	last := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= last && intervals[i][1] > last { // s合并
			last = intervals[i][1]
		} else if intervals[i][0] <= last && intervals[i][1] <= last {

		} else {
			result = append(result, []int{front, last})
			front = intervals[i][0]
			last = intervals[i][1]
		}
	}
	result = append(result, []int{front, last})
	return result
}

func main() {

}
