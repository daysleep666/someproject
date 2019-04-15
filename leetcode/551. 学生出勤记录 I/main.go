package main

func checkRecord(s string) bool {
	a, l := 0, 0
	for i, _ := range s {
		if s[i] == 'A' {
			a++
			if a > 1 {
				return false
			}
			l = 0
		} else if s[i] == 'L' {
			l++
			if l > 2 {
				return false
			}
		} else {
			l = 0
		}
	}
	return true
}

func main() {

}
