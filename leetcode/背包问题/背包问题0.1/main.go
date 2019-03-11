package main

import "fmt"

// Problem Description:
//    有 n 个重量和价值分别为Wi,Vi的物品，现从这些物品中挑选出总量不超过 W 的物品，求所有方案中价值总和的最大值。
// Input:
// 输入包含多组测试用例，每一例的开头为两位整数 n、W（1<=n<=10000,1<=W<=1000）
// ，接下来有 n 行，每一行有两位整数 Wi、Vi（1<=Wi<=10000,1<=Vi<=100）。

var maxValue int
var mw int
var length int

func Bag1(weight []int, value []int, maxWeight int) { // 暴力法
	mw = maxWeight
	length = len(weight)
	d(weight, value, 0, 0, 0)
	fmt.Println(maxValue)
}

func d(weight []int, value []int, i, curWeight int, curValue int) {
	if curValue > maxValue {
		maxValue = curValue
		return
	}

	if i >= length {
		return
	}

	// 不放自己
	d(weight, value, i+1, curWeight, curValue)

	if curWeight+weight[i] > mw { // 放了自己进去不能超过最大重量
		return
	}
	// 放自己进去
	d(weight, value, i+1, curWeight+weight[i], curValue+value[i])
}

func Bag2(weight []int, value []int, maxWeight int) {
	ln := len(weight)
	state := make([][]int, ln) // state[第x次决策][此时重量] = 第x决策的完的总价值
	for i, _ := range state {
		state[i] = make([]int, maxWeight+1)
		for j, _ := range state[i] {
			state[i][j] = -1
		}
	}

	state[0][0] = 0                // 第1次决策，不放1进去
	state[0][weight[0]] = value[0] // 第1次决策，放1进去

	for i := 1; i < ln; i++ {
		for j := 0; j < maxWeight+1; j++ {
			if state[i-1][j] != -1 {
				state[i][j] = state[i-1][j] // 这是不放当前物品的结果
				if j+weight[i] <= maxWeight {
					state[i][j+weight[i]] = state[i-1][j] + value[i] // 这是放当前物品的结果
				}
			}
		}
	}

	maxValue := 0
	for i := 0; i < maxWeight+1; i++ {
		if state[ln-1][i] > maxValue {
			maxValue = state[ln-1][i]
		}
	}
	fmt.Println(maxValue, state)
}

func main() {
	Bag1([]int{1, 2, 3, 4}, []int{1, 2, 3, 1}, 6)
	Bag2([]int{1, 2, 3, 4}, []int{1, 2, 3, 1}, 6)
}
