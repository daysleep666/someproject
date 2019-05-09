package main

import "fmt"

func maxAreaOfIsland(grid [][]int) int {
	max := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				count := compute(grid, i, j)
				if count > max {
					max = count
				}
			}
		}
	}
	// fmt.Println(grid)
	return max
}

func compute(grid [][]int, i, j int) int {
	conut := 0
	if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) && grid[i][j] == 1 {
		conut++
		grid[i][j] = 0
		conut += compute(grid, i, j+1)
		conut += compute(grid, i, j-1)
		conut += compute(grid, i+1, j)
		conut += compute(grid, i-1, j)
	}
	return conut
}

func main() {
	max := maxAreaOfIsland([][]int{[]int{1, 1, 0, 0, 0},
		[]int{1, 1, 0, 0, 0},
		[]int{1, 1, 0, 0, 0},
		[]int{1, 1, 0, 0, 0},
	})
	fmt.Println(max)
}
