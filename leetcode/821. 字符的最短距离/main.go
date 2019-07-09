package main

import "fmt"

func shortestToChar(S string, C byte) []int {
	var pos int = -1
	var result = make([]int, len(S))
	var t int
	for i := 0; i < len(S); i++ {
		if S[i] == C {
			result[i] = 0
			pos = i
			t = 1
			var tmp int = 1
			for j := i - 1; j >= 0; j-- {
				if S[j] != C && (result[j] > tmp || result[j] == 0) {
					result[j] = tmp
				} else {
					break
				}
				tmp++
			}
		} else {
			if pos == -1 {
				continue
			} else {
				result[i] = t
				t++
			}
		}
	}
	return result
}

func main() {
	fmt.Println(shortestToChar("loveleetcode", 'e'))
}
