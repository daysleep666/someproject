package main

import "fmt"

func reverseWords(s string) string {
	var words []string
	var str string
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if str != "" {
				words = append(words, str)
				str = ""
			}
		} else {
			str += string(s[i])
		}
	}
	if str != "" {
		words = append(words, str)
		str = ""
	}

	for i := len(words) - 1; i >= 0; i-- {
		str += string(words[i]) + " "
	}
	if len(str) > 0 {
		return str[:len(str)-1]
	}
	return ""
}

func main() {
	// fmt.Println(reverseWords("the sky is blue"))
	// fmt.Println(reverseWords("  hello world!  "))
	// fmt.Println(reverseWords("a good   example"))
	fmt.Println(reverseWords(""))
}
