package main

import "fmt"

func magicalString(n int) int {
	o1 := make([]int, 2)
	o2 := make([]int, 1)
	o1[0] = 1
	o1[1] = 2
	o2[0] = 1
	last := 1
	if n == 0 {
		return 0
	}
	count := 1

	for i := 1; i < n; i++ {
		if o1[i] == 1 {
			count++
			if last == 1 {
				o2 = append(o2, 2)
				last = 2
			} else {
				o2 = append(o2, 1)
				last = 1
			}
		} else {
			if last == 1 {
				o2 = append(o2, 2, 2)

				last = 2
			} else {
				o2 = append(o2, 1, 1)

				last = 1
			}
		}
		o1 = o2
	}

	return count
}

func main() {
	fmt.Println(magicalString(6))
}
