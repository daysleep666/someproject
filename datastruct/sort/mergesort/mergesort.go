package mergesort

// 由大到小排

func MergeSort(_arr []int64) []int64 {
	if len(_arr) == 1 {
		return _arr
	} else if len(_arr) == 2 {
		if _arr[0] > _arr[1] {
			return []int64{_arr[0], _arr[1]}
		} else {
			return []int64{_arr[1], _arr[0]}
		}
	}
	arr1 := MergeSort(_arr[:len(_arr)/2])
	arr2 := MergeSort(_arr[len(_arr)/2:])
	return merge(arr1, arr2)
}

func merge(_arr1, _arr2 []int64) []int64 {
	tmpArr := make([]int64, len(_arr1)+len(_arr2))
	i, j := 0, 0
	for i < len(_arr1) && j < len(_arr2) {
		if _arr1[i] >= _arr2[j] {
			tmpArr[i+j] = _arr1[i]
			i++
		} else {
			tmpArr[i+j] = _arr2[j]
			j++
		}
	}
	if i < len(_arr1) {
		tmpArr[i+j] = _arr1[i]
		i++
	}
	if j < len(_arr2) {
		tmpArr[i+j] = _arr2[j]
		j++
	}
	return tmpArr
}
