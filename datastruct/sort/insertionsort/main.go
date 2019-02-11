package insertionsort

// 插入排序 由小到大
func InsertionSort(_arr []int64) []int64 {
	length := len(_arr)
	for i := 0; i < length-1; i++ {
		k := i + 1
		tmp := _arr[k]
		for ; k > 0; k-- {
			if _arr[k-1] > _arr[k] {
				_arr[k] = _arr[k-1]
			} else {
				break
			}
		}
		_arr[k] = tmp
	}
	return _arr
}

// 时间复杂度
// 最好时间复杂度 O(n)
// 最坏时间复杂度 O(n^2)
// 平均时间复杂度 O(n^2)
// 		计算方式: 最小有序度0，最大有序度(1+n)*n/2	【有序度等于交换次数】
//				  平均有序度是 (0 + (1+n)*n/2) / 2 = O(n^2)

// 空间复杂度 O(1)
// 原地排序

// 是稳定排序
