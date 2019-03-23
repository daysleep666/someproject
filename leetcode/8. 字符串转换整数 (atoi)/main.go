package main

import (
	"fmt"
	"math"
)

func myAtoi(str string) int {
	sign := 0
	num := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '-' {
			if sign != 0 {
				return 0
			}
			sign = -1
			continue
		}
		if str[i] == '+' {
			if sign != 0 {
				return 0
			}
			sign = 1
			continue
		}
		if str[i] == ' ' && sign == 0 {
			continue
		}
		if n, is := number(str[i]); !is {
			if sign == -1 {
				return -num
			}
			return num
		} else {
			if num == 0 && n == 0 {
				return 0
			}
			if sign == 0 {
				sign = 1
			}
			num = num*10 + n
			if sign == 1 {
				if num >= math.MaxInt32 {
					return math.MaxInt32
				}
			}
			if sign == -1 {
				if -num < math.MinInt32 {
					return math.MinInt32
				}
			}
		}
	}
	if sign == -1 {
		return -num
	}
	return num
}

func number(a byte) (int, bool) {
	if a >= '0' && a <= '9' {
		return int(a - '0'), true
	}
	return 0, false
}

func main() {
	fmt.Println(-100 % 10)
	fmt.Println(myAtoi("1 123"))
	fmt.Println(myAtoi("-1 123"))

}
