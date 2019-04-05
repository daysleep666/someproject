package main

import "fmt"

func commonChars(A []string) []string {
	if len(A) == 0 {
		return nil
	}
	tmp := A[0]
	arr := make([]int, 26)
	reset(tmp, arr)

	for i := 1; i < len(A); i++ {
		tmp = ""
		for j, _ := range A[i] {
			if arr[A[i][j]-'a'] > 0 {
				tmp += string(A[i][j])
			}
			arr[A[i][j]-'a']--
		}
		reset(tmp, arr)
	}
	fmt.Println(tmp)
	result := make([]string, len(tmp))
	for i, _ := range tmp {
		result[i] = string(tmp[i])
	}
	return result
}

func reset(str string, arr []int) {
	for i, _ := range arr {
		arr[i] = 0
	}
	for i, _ := range str {
		arr[str[i]-'a']++
	}
}

func main() {
	fmt.Println(commonChars([]string{"bella", "label", "roller"}))
	fmt.Println(commonChars([]string{"bella"}))
}
