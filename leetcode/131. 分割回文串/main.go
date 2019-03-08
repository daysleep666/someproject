package main

import "fmt"

var m map[string][][]string

func partition(s string) [][]string {
	m = make(map[string][][]string)
	return d(s)
}

func d(s string) [][]string {
	if len(s) == 0 {
		return [][]string{[]string{""}}
	}
	arr := make([][]string, 0)
	for i := 1; i < len(s)+1; i++ {
		v := s[:i]
		if isP(v) {
			tmp, isExist := m[s[i:]]
			if !isExist {
				tmp = d(s[i:])
				m[s[i:]] = tmp
			}
			for _, t := range tmp {
				arr = append(arr, []string{})
				k := len(arr) - 1
				arr[k] = append(arr[k], v)
				for _, tt := range t {
					if tt != "" {
						arr[k] = append(arr[k], tt)
					}
				}
			}
		}
	}
	return arr
}

func isP(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	fmt.Println(partition("cdd"))
}
