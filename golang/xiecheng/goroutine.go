package main

import (
	"fmt"
	"runtime"
)

func say(k int) {

	fmt.Println(k)
}

func main() {
	runtime.GOMAXPROCS(1)

	for i := 0; i < 100; i++ {
		go say(i)
	}

	for {

	}
}
