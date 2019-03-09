package main

import "fmt"

func reverseString(s []byte) {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
	fmt.Println(s)
}

func main() {
	reverseString([]byte{'h', 'e', 'l', 'l', 'o'})
}
