package main

import "fmt"

func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for i, _ := range result {
		result[i] = make([]int, n)
	}

	i, j, k := 0, 0, 1

	for {
		o := k
		for ; j <= n-1 && i <= n-1; j++ {
			if result[i][j] != 0 {
				break
			}
			result[i][j] = k
			k++
		}
		j--
		i++

		for ; i <= n-1 && j >= 0; i++ {
			if result[i][j] != 0 {
				break
			}
			result[i][j] = k
			k++
		}
		i--
		j--

		for ; j >= 0 && i >= 0; j-- {
			if result[i][j] != 0 {
				break
			}
			result[i][j] = k
			k++
		}
		j++
		i--

		for ; i >= 0 && j <= n-1; i-- {
			if result[i][j] != 0 {
				break
			}
			result[i][j] = k
			k++
		}
		i++
		j++

		if o == k {
			break
		}
	}

	return result
}

func main() {
	fmt.Println(generateMatrix(2))
}
