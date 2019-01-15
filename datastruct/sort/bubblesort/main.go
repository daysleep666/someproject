package main

import "fmt"

// 从小到大
func BubbleSort(_arr []int) []int {
	length := len(_arr)
	for i := 0; i < length-1; i++ {
		noMove := false
		for j := 1; j < length; j++ {
			if _arr[j-1] > _arr[j] {
				_arr[j-1], _arr[j] = _arr[j], _arr[j-1]
				noMove = true
			}
		}
		if !noMove {
			// 这次排序没有移动，说明已经是正确的顺序了
			return _arr
		}
	}
	return _arr
}

func main() {
	result := BubbleSort([]int{3, 2, 4, 1, 5})
	fmt.Println(result)
}

// 最好时间复杂度O(n) 最坏时间复杂度是O(n^2)
// 原地排序 不需要额外空间 空间复杂度O(1)
// 是稳定排序 相同的两个数不会交换数据

// 平均复杂度计算
// 3, 2, 4, 1, 5
// 有序度是6:  		(3,4) (3,5) (2,4) (2,5) (4,5) (1,5)
// 逆有序度是4:  	(3,2) (3,1) (2,1) (4,1)
// 满有序度是10

// 对于排序来说，是比较和移动，每移动一次，逆有序度-1，有序度+1。
// 有序度从0到n*(n+1)/2。平均有序度是n*(n+1)/4。
// 因此 平均复杂度是 O(n^2)
