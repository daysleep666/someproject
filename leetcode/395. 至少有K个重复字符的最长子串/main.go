package main

func longestSubstring(s string, k int) int {
	m := make([]int, 128)
	for i, _ := range s {
		m[int(s[i])]++
	}

	for i, _ := range s {
		if m[int(s[i])] < k {
			a := longestSubstring(s[:i], k)
			b := longestSubstring(s[i+1:], k)
			if a > b {
				return a
			} else {
				return b
			}
		}
	}
	return len(s)
}

func main() {

}
