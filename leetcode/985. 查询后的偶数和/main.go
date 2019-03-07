package main

import "fmt"

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	result := make([]int, len(queries))
	origin := c(A)
	for i, v := range queries {
		value := v[0]
		index := v[1]
		if A[index]%2 == 0 {
			origin -= A[index]
		}
		A[index] += value
		if A[index]%2 == 0 {
			origin += A[index]
		}
		result[i] = origin
	}
	return result
}

func c(a []int) int {
	result := 0
	for _, v := range a {
		if v%2 == 0 {
			result += v
		}
	}
	return result
}

func main() {
	fmt.Println(sumEvenAfterQueries([]int{1, 2, 3, 4}, [][]int{[]int{1, 0}, []int{-3, 1}, []int{-4, 0}, []int{2, 3}}))
}
