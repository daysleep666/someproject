package main

import "fmt"

func validMountainArray(A []int) bool {
	if len(A) < 2 || (len(A) >= 2 && A[0] > A[1]) {
		return false
	}
	down := false
	for i := 2; i < len(A); i++ {
		if A[i] == A[i-1] {
			return false
		} else if A[i] < A[i-1] {
			down = true
		} else {
			if down {
				return false
			}
		}
	}
	return down
}

func main() {
	fmt.Println(validMountainArray([]int{2, 1}))
	fmt.Println(validMountainArray([]int{2, 5, 5}))
	fmt.Println(validMountainArray([]int{0, 3, 2, 1}))
	fmt.Println(validMountainArray([]int{0, 1, 2, 3}))
}
