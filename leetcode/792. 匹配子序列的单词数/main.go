package main

import "fmt"

func numMatchingSubseq(S string, words []string) int {
	nums := make([][]string, 26, 26)
	for i, _ := range words {
		nums[words[i][0]-'a'] = append(nums[words[i][0]-'a'], words[i])
	}

	result := 0
	for i, _ := range S {
		tmp := S[i] - 'a'
		if len(nums[tmp]) == 0 {
			continue
		}

		tmpNums := nums[tmp]
		nums[tmp] = nums[tmp][:0]
		for _, v := range tmpNums {
			if len(v) == 1 {
				result++
			} else {
				v = v[1:]
				nums[v[0]-'a'] = append(nums[v[0]-'a'], v)
			}
		}
	}
	return result
}

func main() {
	fmt.Println(numMatchingSubseq("abcde", []string{"a", "bb", "acd", "ace"}))
	fmt.Println(numMatchingSubseq("dsahjpjauf", []string{"ahjpjau", "ja", "ahbwzgqnuk", "tnmlanowax"}))
}
