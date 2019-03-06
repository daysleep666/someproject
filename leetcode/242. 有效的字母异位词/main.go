package main

import "fmt"

func isAnagram(s string, t string) bool {
	ls := len(s)
	lt := len(t)
	if ls != lt {
		return false
	}
	m := make([]int, 26)
	for i := 0; i < ls; i++ {
		m[s[i]-'a']++
	}
	for i := 0; i < ls; i++ {
		m[t[i]-'a']--
		if m[t[i]-'a'] < 0 {
			return false
		}
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("aacc", "ccaa"))
}
