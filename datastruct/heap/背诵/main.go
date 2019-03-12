package main

import (
	"fmt"
)

func LE(data []int, a, b int) bool {
	return data[a] <= data[b]
}

func CheckMin(data []int, lo, hi int) bool {
	if lo >= hi {
		return true
	}

	for cur := lo; cur <= (hi-lo-1)/2; cur++ {
		l := (cur-lo)*2 + 1
		if !LE(data, cur, l) {
			return false
		}
		r := l + 1
		if r > hi {
			break
		}
		if !LE(data, cur, r) {
			return false
		}
	}
	return true
}

// 小顶堆
func heapify(nums []int, start, end int) {
	for i := end / 2; i >= 0; i-- {
		downHeap(nums, i, end)
	}
}

func downHeap(nums []int, start int, end int) {
	for {
		childIndex := start
		if start*2+1 < end && nums[start] > nums[start*2+1] {
			childIndex = start*2 + 1
		}
		if start*2+2 < end && nums[childIndex] > nums[start*2+2] {
			childIndex = start*2 + 2
		}
		if start != childIndex {
			nums[start], nums[childIndex] = nums[childIndex], nums[start]
			start = childIndex
		} else {
			break
		}
	}
}

func heapSort(nums []int) { // 小顶堆
	heapify(nums, 0, len(nums)-1)

	for i := len(nums) - 1; i >= 0; i-- { // 循环将堆顶元素移到数组的最后。  由大到小排序
		nums[0], nums[i] = nums[i], nums[0]
		heapify(nums, 0, i)
	}

}

func main() {
	// n := 5
	// arr := make([]int, 0)
	// for i := n; i >= 0; i-- {
	// 	arr = append(arr, rand.Intn(10))
	// }
	arr := []int{5, 4, 3, 2, 1, 10}
	// heapify(arr, 0, len(arr)-1)
	heapSort(arr)
	fmt.Println(arr)
}

// func isPowerOfTwo(x uintptr) bool {
// 	return x&(x-1) == 0
// }
