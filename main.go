package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n := scanner.Text()
	ni, _ := strconv.Atoi(n)
	if ni%2 == 0 && ni != 2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
