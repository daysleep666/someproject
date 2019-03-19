package main

import (
	"fmt"
)

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	l, w := len(matrix), len(matrix[0])
	result := make([]int, l*w)
	i, j, di, dj := 0, 0, 0, 1
	for index := 0; index < l*w; index++ {
		result[index] = matrix[i][j]
		matrix[i][j] = 0
		if matrix[gety(i, di, l)][gety(j, dj, w)] == 0 {
			di, dj = dj, -di
		}
		i, j = i+di, j+dj
	}
	return result
}

func gety(i, d, l int) int {
	if i+d < 0 {
		return 0
	}
	return (i + d) % l
}

func main() {
	fmt.Println(spiralOrder([][]int{[]int{1, 2, 3, 4}, []int{5, 6, 7, 8}, []int{9, 10, 11, 12}}))
	// fmt.Println(spiralOrder([][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}))
	// fmt.Println(spiralOrder([][]int{[]int{1}, []int{4}, []int{7}}))
}
