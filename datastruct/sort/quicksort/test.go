package quicksort

import "fmt"

// 题目：在数组中找出第n小的数据
func Test() {
	arr := []int64{3, 4, 1, 2, 5}
	v := find(arr, 0, int64(len(arr)-1), 1)
	fmt.Println(v)
}

func find(_arr []int64, _from, _to, _n int64) int64 {
	if _from >= _to {
		return _arr[_from]
	}

	var v int64
	mid := partitionFind(_arr, _from, _to)
	if mid+1 == _n {
		return _arr[mid]
	} else if mid+1 > _n {
		v = find(_arr, _from, mid-1, _n)
	} else {
		v = find(_arr, mid+1, _to, _n)
	}
	return v
}

func partitionFind(_arr []int64, _from, _to int64) int64 {
	var (
		i         = _from
		j         = _from
		partition = _arr[_to]
	)

	for ; i < _to; i++ {
		if _arr[i] <= partition {
			_arr[i], _arr[j] = _arr[j], _arr[i]
			j++
		}
	}
	_arr[_to], _arr[j] = _arr[j], _arr[_to]
	return j
}
