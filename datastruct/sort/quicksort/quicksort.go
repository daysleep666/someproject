package quicksort

func QuickSort(_arr []int64) {
	quickSort(_arr, 0, int64(len(_arr)-1))
}

func quickSort(_arr []int64, _from, _to int64) {
	if _from >= _to {
		return
	}
	mid := partition(_arr, _from, _to)
	quickSort(_arr, _from, mid-1)
	quickSort(_arr, mid+1, _to)
}

func partition(_arr []int64, _from, _to int64) int64 {
	pivot := _arr[_to]
	i, j := _from, _from
	for ; j < _to; j++ {
		if _arr[j] < pivot {
			_arr[i], _arr[j] = _arr[j], _arr[i]
			i++
		}
	}
	_arr[_to], _arr[i] = _arr[i], _arr[_to]
	return i
}
