package main

import (
	"fmt"
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	i, j, result := 0, len(people)-1, 0
	for i <= j {
		if people[i]+people[j] <= limit {
			i, j = i+1, j-1
		} else {
			j = j - 1
		}
		result++
	}
	return result
}

func main() {
	fmt.Println(numRescueBoats([]int{1, 2}, 3))
	fmt.Println(numRescueBoats([]int{3, 5, 3, 4}, 5))
	fmt.Println(numRescueBoats([]int{}, 5))
}
