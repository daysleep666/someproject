package main

import "fmt"

func titleToNumber(s string) int {
	result := 0
	temp := 1
	for i := len(s) - 1; i >= 0; i-- {
		v := getReal(s[i])
		result += v * temp
		temp *= 26
	}
	return result
}

func getReal(a byte) int {
	return int(a-'A') + 1
}

func main() {
	fmt.Println(titleToNumber("A"))
	fmt.Println(titleToNumber("AB"))
	fmt.Println(titleToNumber("ZY"))
}
