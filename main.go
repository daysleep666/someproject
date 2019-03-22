package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	m["d"] = 4

	for k, v := range m {
		fmt.Printf("%v:%v,", k, v)
	}
}
