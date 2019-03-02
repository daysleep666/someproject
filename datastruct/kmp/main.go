package main

import "fmt"

func kmp(mainStr, subStr string) int {
	lenMain := len(mainStr)
	lenSub := len(subStr)
	pFind := p(subStr)
	for i := 0; i < lenMain; {
		j := 0
		for j < lenSub && i < lenMain {
			if j == -1 || mainStr[i] == subStr[j] {
				i++
				j++
			} else {
				j = pFind[j]
			}
		}
		if j == lenSub {
			return i - lenSub
		}
	}
	return -1
}

func p(str string) []int {
	arr := make([]int, len(str))
	i, j := -1, 0
	for j < len(str) {
		if i == -1 {
			arr[j] = i
			i++
			j++
		} else if str[i] == str[j] {
			arr[j] = arr[i]
			i++
			j++
		} else {
			arr[j] = arr[i]
			j++
		}
	}
	return arr
}

func main() {
	fmt.Println(kmp("abcabdabeabf", "sa"))
}
