package main

import "fmt"

func combine(n int, k int) [][]int {
	if k == 0 || k > n {
		return [][]int{}
	}
	return d(1, n, k)
}

func d(from, to, k int) [][]int {
	result := make([][]int, 0)
	if k <= 1 {
		for i := from; i <= to; i++ {
			result = append(result, []int{i})
		}
		return result
	}

	for i := from; i < to; i++ {
		tmp := d(i+1, to, k-1)
		for _, v := range tmp {
			result = append(result, append([]int{i}, v...))
		}
	}
	return result
}

func main() {
	fmt.Println(combine(4, 2))
	fmt.Println(combine(2, 2))
	fmt.Println(combine(5, 2))
}
