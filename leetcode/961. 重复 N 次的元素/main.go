package main

import "sort"

func repeatedNTimes1(A []int) int {
	n := len(A) / 2
	arr := make([]int, 10000)
	for _, v := range A {
		arr[v]++
	}
	for i, v := range arr {
		if v == n {
			return i
		}
	}
	return 0
}

func repeatedNTimes(A []int) int {
	sort.Ints(A)
	n := (len(A) - 1) / 2
	if A[0] == A[n] {
		return A[n]
	}
	return A[n+1]
}

func main() {

}
