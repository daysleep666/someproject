package main

import (
	"fmt"
	"strings"
)

func numUniqueEmails(emails []string) int {
	m := make(map[string]bool, 0)
	for _, email := range emails {
		strs := strings.Split(email, "@")
		var tmp string
		for i, _ := range strs[0] {
			if strs[0][i] == '.' {
				continue
			} else if strs[0][i] == '+' {
				break
			} else {
				tmp += string(strs[0][i])
			}
		}
		m[tmp+"@"+strs[1]] = true
	}
	fmt.Println(m)
	return len(m)
}

func main() {
	fmt.Println(numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}))
}
