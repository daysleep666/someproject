package main

import (
	"fmt"
	"strconv"
)

func maximumSwap(num int) int {
	tmpBytes := []byte(fmt.Sprintf("%v", num))
	for i := 0; i < len(tmpBytes)-1; i++ {
		if tmpBytes[i] == '9' {
			continue
		}
		// 找到当前字符右边最大的数
		maxPos := findMax(tmpBytes[i+1:]) + i + 1
		// fmt.Println(string(tmpBytes[maxPos]))
		if tmpBytes[i] < tmpBytes[maxPos] {
			// fmt.Println(string(tmpBytes[i]), string(tmpBytes[maxPos]))
			// 交换
			tmpBytes[i], tmpBytes[maxPos] = tmpBytes[maxPos], tmpBytes[i]
			n, _ := strconv.Atoi(string(tmpBytes))
			return n
		}
	}

	return num
}

func findMax(bs []byte) int {
	index := 0
	max := bs[0]
	for i, v := range bs {
		if v >= max {
			max = v
			index = i
			// fmt.Println("---", string(max), index)
		}
	}
	return index
}

func main() {
	// fmt.Println(maximumSwap(98368))
	// fmt.Println(maximumSwap(215))
	// fmt.Println(maximumSwap(9908))
	fmt.Println(maximumSwap(1993))
}
