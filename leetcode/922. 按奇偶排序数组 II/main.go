package main

import "fmt"

func sortArrayByParityII(A []int) []int {
	for i, j := 0, 1; i < len(A) && j < len(A); {
		for i < len(A) && A[i]%2 == 0 {
			i = i + 2
		}
		for j < len(A) && A[j]%2 != 0 {
			j = j + 2
		}
		if i < len(A) && j < len(A) {
			A[i], A[j] = A[j], A[i]
		}
	}
	return A
}

func main() {
	fmt.Println(sortArrayByParityII([]int{4, 2, 5, 7}))
}
