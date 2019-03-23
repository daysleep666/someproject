package main

import "fmt"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	arr := make([]int, 0)
	for i := len(num1) - 1; i >= 0; i-- {
		index := len(num1) - i - 1
		jinwei := 0
		for j := len(num2) - 1; j >= 0; j-- {
			a := convertToNum(num1[i])
			b := convertToNum(num2[j])
			mul := a*b + jinwei
			jinwei = mul / 10
			v := mul % 10
			fmt.Println(arr, index, v)
			add(&arr, index, v)
			index++
		}
		if jinwei > 0 {
			add(&arr, index, jinwei)
		}

	}
	result := ""
	for j := len(arr) - 1; j >= 0; j-- {
		result = fmt.Sprintf("%v%v", result, arr[j])
	}
	fmt.Println(result)
	return result
}

func add(arr *[]int, index, v int) {
	if index >= len((*arr)) {
		(*arr) = append((*arr), v)
		return
	}
	jinwei := 0
	s := (*arr)[index] + v
	jinwei = s / 10
	(*arr)[index] = s % 10
	if jinwei == 0 {
		return
	}
	add(arr, index+1, jinwei)
}

func convertToNum(num byte) int {
	return int(num - '0')
}

func main() {
	multiply("0", "456")
}

//----------------
func multiply1(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	zero := ""
	result := ""
	for i := len(num1) - 1; i >= 0; i-- {
		tmpresult := zero
		jinwei := 0
		for j := len(num2) - 1; j >= 0; j-- {
			a := convertToNum(num1[i])
			b := convertToNum(num2[j])
			mul := a*b + jinwei
			jinwei = mul / 10
			tmpresult = fmt.Sprintf("%v%v", mul%10, tmpresult)
		}
		if jinwei > 0 {
			tmpresult = fmt.Sprintf("%v%v", jinwei, tmpresult)
		}
		fmt.Println("sum =", tmpresult)
		zero += "0"
		result = sum(result, tmpresult)
	}
	fmt.Println(result)
	return result
}

func sum(num1 string, num2 string) string {
	result := ""
	jinwei := 0
	i, j := len(num1)-1, len(num2)-1
	for ; i >= 0 && j >= 0; i, j = i-1, j-1 {
		a := convertToNum(num1[i])
		b := convertToNum(num2[j])
		mul := a + b + jinwei
		jinwei = mul / 10
		result = fmt.Sprintf("%v%v", mul%10, result)
	}

	for ; i >= 0; i-- {
		a := convertToNum(num1[i])
		mul := a + jinwei
		jinwei = mul / 10
		result = fmt.Sprintf("%v%v", mul%10, result)
	}
	for ; j >= 0; j-- {
		a := convertToNum(num2[j])
		mul := a + jinwei
		jinwei = mul / 10
		result = fmt.Sprintf("%v%v", mul%10, result)
	}
	if jinwei > 0 {
		result = fmt.Sprintf("%v%v", jinwei, result)
	}
	return result
}
