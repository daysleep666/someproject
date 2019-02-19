package main

import "fmt"

func heap(_arr []int64, length, i int64) {
	for {
		// 当前节点，当前节点的左孩子，当前节点的右孩子， 选最大的放到当前节点
		maxPos := i // 当前最大的节点下标
		if 2*i+1 < length && _arr[maxPos] > _arr[2*i+1] {
			maxPos = 2*i + 1
		}
		if 2*i+2 < length && _arr[maxPos] > _arr[2*i+2] {
			maxPos = 2*i + 2
		}
		if maxPos == i { // 当前节点为最大节点 就不做任何
			return
		}
		_arr[i], _arr[maxPos] = _arr[maxPos], _arr[i]
		i = maxPos
	}
}

func buildHeap(_arr []int64) { // 将一个数组变为一个小顶堆
	// 只需要考虑非叶子节点
	length := int64(len(_arr))
	for i := (length - 1) / 2; i >= 0; i-- {
		heap(_arr, length, i)
	}
}

func HeapSort(_arr []int64) {
	// 1.堆化
	buildHeap(_arr)
	// 2.输出堆顶
	length := int64(len(_arr))
	lastIndex := length - 1
	for i := lastIndex; i > 0; i-- {
		// 将堆顶放和堆的最后一位交换
		_arr[0], _arr[i] = _arr[i], _arr[0]
		// 重新堆化
		heap(_arr, i, 0)
	}
}

func main() {
	arr := []int64{5, 4, 3, 2, 1, 10}
	HeapSort(arr)
	fmt.Println(arr)
}

//时间复杂度 O(nlogn)
