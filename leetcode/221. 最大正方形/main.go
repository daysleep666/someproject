package main

import "fmt"

func maximalSquare(matrix [][]byte) int {
	max := 0
	var v1, v2, v3, v4 int
	for i, _ := range matrix {
		for j, _ := range matrix[i] {
			if i == 0 || j == 0 {
				if matrix[i][j] == '1' && max == 0 {
					max = 1
				}
			} else {
				v1 = int(matrix[i][j] - '0')
				v2 = int(matrix[i][j-1] - '0')
				v3 = int(matrix[i-1][j] - '0')
				v4 = int(matrix[i-1][j-1] - '0')
				if v1 > 0 && v2 > 0 && v3 > 0 && v4 > 0 {
					v1 = min(v2, v3, v4) + 1
					matrix[i][j] = byte(v1) + '0'
					if max < v1 {
						max = v1
					}

				}
			}
		}
	}

	return max * max
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	}
	return c
}
func main() {
	fmt.Println(maximalSquare([][]byte{
		[]byte{'1', '1', '1', '0', '0'},
		[]byte{'1', '1', '1', '0', '0'},
		[]byte{'1', '1', '1', '1', '1'},
		[]byte{'0', '1', '1', '1', '1'},
		[]byte{'0', '1', '1', '1', '1'},
		[]byte{'0', '1', '1', '1', '1'},
	}))

	fmt.Println(maximalSquare([][]byte{
		[]byte{'1', '0', '1', '0', '0'},
		[]byte{'1', '0', '1', '1', '1'},
		[]byte{'1', '1', '1', '1', '1'},
		[]byte{'1', '0', '0', '1', '0'},
	}))

	fmt.Println(maximalSquare([][]byte{
		[]byte{'0', '0', '0', '1'},
		[]byte{'1', '1', '0', '1'},
		[]byte{'1', '1', '1', '1'},
		[]byte{'0', '1', '1', '1'},
		[]byte{'0', '1', '1', '1'},
	}))
}
