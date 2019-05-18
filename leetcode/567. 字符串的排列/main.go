package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	m := make([]int, 26)
	for i, _ := range s1 {
		m[s1[i]-'a']++
	}
	k := len(s1)

	for i, j := 0, 0; i < len(s2); {
		if m[s2[i]-'a'] > 0 {
			m[s2[i]-'a']--
			k--
			i++
		} else {
			if i == j {
				i++
				j++
				continue
			}
			m[s2[j]-'a']++
			k++
			j++
		}
		if k == 0 {
			return true
		}
	}
	if k == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(checkInclusion("adc", "dcda"))
	fmt.Println(checkInclusion("ab", "eidbaooo"))
	fmt.Println(checkInclusion("ab", "eidboaoo"))
}
