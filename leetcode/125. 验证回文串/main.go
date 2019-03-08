package main

import "fmt"

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {

		for i < len(s)-1 && !isValid(s[i]) {
			i++
		}
		for j > 0 && !isValid(s[j]) {
			j--
		}

		if i >= j {
			break
		}

		if toLow(s[i]) != toLow(s[j]) {
			fmt.Println(i, j)
			return false
		}
		i++
		j--
	}
	return true
}

func isValid(a byte) bool {
	if ('a' <= a && a <= 'z') || ('A' <= a && a <= 'Z') || ('0' <= a && a <= '9') {
		return true
	}
	return false
}

func toLow(a byte) byte {
	if 'a' <= a && a <= 'z' {
		return a - 32
	}
	return a
}

func main() {
	fmt.Println(isPalindrome(".,"))
}
