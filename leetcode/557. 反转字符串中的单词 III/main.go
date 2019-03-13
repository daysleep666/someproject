package main

import "fmt"

func reverseWords(s string) string {
	low := 0
	tmp := []rune(s)
	for i, v := range s {
		if v == ' ' {
			reverse(tmp, low, i-1)
			low = i + 1
		}
	}
	reverse(tmp, low, len(s)-1)
	return string(tmp)
}

func reverse(str []rune, low, high int) {
	for i, j := low, high; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
}

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}
