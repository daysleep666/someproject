package main

import "fmt"

func judgeCircle(moves string) bool {
	if len(moves)%2 != 0 {
		return false
	} else if len(moves) == 0 {
		return true
	}

	x, y := 0, 0
	for i := 0; i < len(moves); i++ {
		if moves[i] == 'L' {
			x--
		} else if moves[i] == 'R' {
			x++
		} else if moves[i] == 'U' {
			y++
		} else {
			y--
		}
	}
	return x == 0 && y == 0
}

func main() {
	fmt.Println(judgeCircle("RLUURDDDLU"))
	fmt.Println(judgeCircle("LL"))
	// fmt.Println(judgeCircle("LR"))
	// fmt.Println(judgeCircle("LUDR"))
}
