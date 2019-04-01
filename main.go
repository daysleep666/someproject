package main

import "fmt"

func main() {
	for i := 0; i < 120; i++ {
		a(i)
	}
}

func a(count int) {
	a := make([]int, 10)

	if len(a) < count {
		fmt.Println("large")
		return
	}
	fmt.Println(a[:count])
}
