package main

import "fmt"

func findCircleNum(M [][]int) int {
	var cnt int
	var list []int
	for k := 0; k < len(M); k++ {
		if M[k][k] == 0 {
			continue
		}
		cnt++
		list = append(list, k)
		for len(list) != 0 {
			var tmp []int
			for n := 0; n < len(list); n++ {
				i := list[n]
				M[i][i] = 0
				for j := 0; j < len(M[i]); j++ {
					if M[i][j] == 1 && M[j][j] == 1 {
						tmp = append(tmp, j)
					}
				}
			}
			list = tmp
		}
	}
	return cnt
}

func main() {
	fmt.Println(findCircleNum([][]int{[]int{1}}))
}
