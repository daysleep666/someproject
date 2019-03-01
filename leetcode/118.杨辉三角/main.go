package main

import "fmt"

func generate(numRows int) [][]int {
	arr := make([][]int, numRows, numRows)
	for i := 0; i < numRows; i++ {
		arr[i] = make([]int, 0, numRows)
		for j := 0; j < i+1; j++ {
			value := 1
			if i > 0 && j > 0 && j < i {
				value = arr[i-1][j-1] + arr[i-1][j]
			}
			arr[i] = append(arr[i], value)
		}
	}
	return arr
}

func main() {
	fmt.Println(generate(5))
}
