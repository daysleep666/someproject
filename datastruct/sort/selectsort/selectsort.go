package selectsort

func SelectSort(_arr []int64) []int64 {
	for i := 0; i < len(_arr)-1; i++ {
		for j := i + 1; j < len(_arr); j++ {
			if _arr[i] > _arr[j] {
				_arr[i], _arr[j] = _arr[j], _arr[i]
			}
		}
	}
	return _arr
}

// 时间复杂度
// 最好时间复杂度 O(n^2)
// 最坏时间复杂度 O(n^2)
// 平均时间复杂度 O(n^2)

// 空间复杂度 O(1)
// 原地排序

// 不是稳定排序
// [3] (3) 1 2
// 1 (3) [3] 2
// 两个3的顺序变了
