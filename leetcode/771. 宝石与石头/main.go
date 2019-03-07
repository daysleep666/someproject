package main

import "fmt"

func numJewelsInStones(J string, S string) int {
	result := 0
	arr := make([]int, 128)
	for _, v := range J {
		arr[v]++
	}
	for _, v := range S {
		if arr[v] > 0 {
			result++
		}
	}

	return result
}

func main() {
	fmt.Println(numJewelsInStones("aA", "z"))
}
