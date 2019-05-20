package main

import "fmt"

func frequencySort(s string) string {
	arr := make([]int, 128)
	for i, _ := range s {
		arr[int(s[i])]++
	}
	arrArr := make([][]int, len(s)+1)
	for i, v := range arr {
		if v > 0 {
			arrArr[v] = append(arrArr[v], i)
		}
	}

	fmt.Println(arrArr)
	result := ""
	for i := len(s); i > 0; i-- {
		for _, v := range arrArr[i] {
			for c := 0; c < i; c++ {
				result += string(byte(v))
			}
		}
	}
	return result
}

func main() {
	fmt.Println(frequencySort("ee"))
}
