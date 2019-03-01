package problem

import "fmt"

type myHeap []int

func (h *myHeap) Push(_n int) {
	curIndex := len(*h)
	*h = append(*h, _n)

	for i := curIndex; i >= 0; {
		if (i-1)/2 >= 0 && (*h)[i] < (*h)[(i-1)/2] {
			(*h)[i], (*h)[(i-1)/2] = (*h)[(i-1)/2], (*h)[i]
			i = (i - 1) / 2
		} else if (i-2)/2 >= 0 && (*h)[i] < (*h)[(i-2)/2] {
			(*h)[i], (*h)[(i-2)/2] = (*h)[(i-2)/2], (*h)[i]
			i = (i - 2) / 2
		} else {
			break
		}
	}

	curIndex++
}

func (h *myHeap) Pop() int {
	if len(*h) <= 1 {
		*h = (*h)[:0]
		return 0
	}
	n := (*h)[0]
	*h = (*h)[1:]
	return n
}

func Partition(_num []int) float64 {
	length := len(_num)
	if length == 0 {
		return 0
	}
	var midNum int
	if length%2 != 0 {
		midNum = length / 2
	} else {
		midNum = (length - 1) / 2
	}

	if length%2 != 0 {
		return float64(compute(_num, midNum))
	}
	return (float64(compute(_num, midNum)) + float64(compute(_num, midNum+1))) / 2
}

func compute(_num []int, midNum int) int {
	length := len(_num)
	from := 0
	to := length - 1
	for {
		mid := partition(_num, from, to)
		if mid == midNum {
			break
		} else if mid > midNum {
			to = mid - 1
		} else {
			from = mid + 1
		}
	}
	return _num[midNum]
}

func partition(_num []int, _from, _to int) int {
	i, j := _from, _from
	midValue := _num[_to]
	for ; i < _to; i++ {
		if _num[i] <= midValue {
			_num[i], _num[j] = _num[j], _num[i]
			j++
		}
	}
	_num[j], _num[_to] = _num[_to], _num[j]
	return j
}

func 第k大(_num []int, _k int) int {
	length := len(_num)
	k := _k - 1
	if length <= k {
		return 0
	}
	from, to := 0, length-1
	for {
		mid := partition(_num, from, to)
		if mid == k {
			return _num[mid]
		} else if mid > k {
			to = mid - 1
		} else {
			from = mid + 1
		}
	}
}

func findMid(_num []int) float64 {
	return Partition(_num)
}

func Test() {
	num := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// fmt.Println(第k大(num, 3))
	fmt.Println(Partition(num))
}
