package main

import "fmt"

func minIncrementForUnique(A []int) int {
	m := make([]int, 40000+len(A))
	for _, v := range A {
		m[v]++
	}
	move := 0
	for _, v := range A {
		for {
			if m[v] <= 1 {
				break
			}
			tmp := m[v] - 1
			move += tmp
			m[v] = 1
			m[v+1] = m[v+1] + tmp
			v++
		}
	}
	return move
}

func main() {
	fmt.Println(minIncrementForUnique([]int{1, 2, 2}))
	fmt.Println(minIncrementForUnique([]int{3, 2, 1, 2, 1, 7}))
}
