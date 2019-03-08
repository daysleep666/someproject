package main

import "fmt"

var m map[string]bool

func wordBreak(s string, wordDict []string) bool {
	m = make(map[string]bool)
	return d(s, wordDict)
}

func d(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}
	for i := 1; i < len(s)+1; i++ {
		word := s[:i]
		fmt.Println(word)
		if find(word, wordDict) {
			if v, isExist := m[s[i:]]; !isExist {
				if d(s[i:], wordDict) {
					m[s[i:]] = true
					return true
				} else {
					m[s[i:]] = false
				}
			} else {
				if v {
					return true
				}
			}
		}
	}
	return false
}

func find(s string, wordDict []string) bool {
	if len(s) == 0 {
		return false
	}
	for _, v := range wordDict {
		if s == v {
			return true
		}
	}
	return false
}

func main() {
	// fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	// fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	// fmt.Println(wordBreak("bb", []string{"a", "b", "bbb", "bbbb"}))
	// fmt.Println(wordBreak("aaaaaaa", []string{"aaaa", "aaa"}))
	// fmt.Println(wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}))
}
