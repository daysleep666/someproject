package main

import "fmt"

func main() {
	a := []interface{}{1, 2, 3, 4}
	test(a...)
}

func test(a ...interface{}) {
	fmt.Println(len(a))
}
