package main

import "fmt"

func minSteps(n int) int {
	if n == 1 {
		return 0
	}
	return copyAndPaste(1, 1, n, 1)
}

func copyAndPaste(num int, copy int, final int, step int) int {
	if num == final {
		return step
	} else if num > final {
		return 9999
	}

	c := 0
	if num > copy {
		c = copyAndPaste(num, num, final, step+1)
	}
	p := copyAndPaste(num+copy, copy, final, step+1)

	if c < p && c != 0 {
		return c
	}
	return p
}

func minSteps2(n int) int {
	return di(n)
}

func di(n int) int {
	if n == 1 {
		return 0
	} else if n == 2 {
		return 2
	} else if n == 3 {
		return 3
	}
	for i := n - 1; i > 0; i-- {
		if n%i == 0 {
			return n/i + di(i)
		}
	}
	return 0
}

func minSteps3(n int) int {
	step := 0
	for {
		if n == 1 {
			return step
		} else if n == 2 {
			return step + 2
		} else if n == 3 {
			return step + 3
		}
		for i := n - 1; i > 0; i-- {
			if n%i == 0 {
				step += n / i
				n = i
				break
			}
		}
	}
}

func main() {
	fmt.Println(minSteps3(8))
}
