package main

import "fmt"

func diStringMatch(S string) []int {
	min, max := 0, len(S)
	result := make([]int, len(S)+1)
	if S[0] == 'I' {
		result[0] = min
		min++
	} else {
		result[0] = max
		max--
	}
	i := 0
	for ; min < max; i++ {
		if (S[i] == 'I' && S[i+1] == 'D') || (S[i] == 'D' && S[i+1] == 'D') {
			result[i+1] = max
			max--
		} else {
			result[i+1] = min
			min++
		}
	}
	result[i+1] = min
	return result
}

func main() {
	fmt.Println(diStringMatch("IDID"))
	fmt.Println(diStringMatch("III"))
	fmt.Println(diStringMatch("DDI"))
	fmt.Println(diStringMatch("DDD"))
	fmt.Println(diStringMatch("D"))
}
