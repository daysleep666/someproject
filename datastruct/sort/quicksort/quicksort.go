package quicksort

func QuickSort(_arr []int64) {
	quickSort(_arr, 0, int64(len(_arr)-1))
}

func quickSort(_arr []int64, _from, _to int64) {
	if _from <= _to {
		return
	}
	partition(_arr, _from, _to)
	mid := _from + _to
	quickSort(_arr, _from, mid/2)
	quickSort(_arr, mid/2+1, _to)
}

func partition(_arr []int64, _from, _to int64) {
	pivot := _arr[_to]
	for i, j := int64(0), int64(0); i < _to-1 && j < _to-1; i++ {
		if _arr[j] <= pivot {
			j++
		} else {
			_arr[i], _arr[j] = _arr[j], _arr[i]
		}
	}
}
