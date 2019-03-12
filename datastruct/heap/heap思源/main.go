package main

// 判断一个数组 [lo,hi] range 是不是最小堆
func CheckMin(data []int, lo, hi int) bool {
	if lo >= hi {
		return true
	}

	for cur := lo; cur <= (hi-lo-1)/2; cur++ {
		l := (cur-lo)*2 + 1
		if !LE(data, cur, l) {
			return false
		}
		r := l + 1
		if r > hi {
			break
		}
		if !LE(data, cur, r) {
			return false
		}
	}
	return true
}

// 判断一个数组 [lo,hi] range 是不是最大堆
func CheckMax(data []int, lo, hi int) bool {
	if lo >= hi {
		return true
	}

	for cur := lo; cur <= (hi-lo-1)/2; cur++ {
		l := (cur-lo)*2 + 1
		if !GE(data, cur, l) {
			return false
		}
		r := l + 1
		if r > hi {
			break
		}
		if !GE(data, cur, r) {
			return false
		}
	}
	return true
}
func LE(data []int, a, b int) bool {
	return data[a] <= data[b]
}
func GE(data []int, a, b int) bool {
	return data[a] >= data[b]
}

func Swap(data []int, a, b int) {
	data[a], data[b] = data[b], data[a]
}
func Heapify(data []int, lo, hi int) {
	lastIdxParent := (hi - lo - 1) / 2
	for cur := lastIdxParent; cur >= lo; cur-- {
		DownMin(data, lo, hi, cur)
	}
}
func DownMin(data []int, lo, hi int, cur int) {
	i := cur
	for {
		l := 2*(i-lo) + 1
		if l > hi {
			break
		}
		if r := l + 1; r <= hi && LE(data, r, l) {
			l = r
		}
		if LE(data, i, l) {
			break
		}
		Swap(data, i, l)
		i = l
	}
}
func Sort(data []int, lo, hi int) {
	HeapifyMax(data, lo, hi)
	for cur := hi; cur > lo; cur-- {
		Swap(data, lo, cur)
		DownMax(data, lo, cur-1, lo)
	}
}
func HeapifyMax(data []int, lo, hi int) {
	lastIdxParent := (hi - lo - 1) / 2
	for cur := lastIdxParent; (cur - lo) >= 0; cur-- {
		DownMax(data, lo, hi, cur)
	}
}
func DownMax(data []int, lo, hi int, cur int) {
	i := cur
	for {
		l := (i-lo)*2 + 1
		if l > hi {
			break
		}
		if r := l + 1; r <= hi && GE(data, r, l) {
			l = r
		}
		if GE(data, i, l) {
			break
		}
		Swap(data, i, l)
		i = l
	}
}

func main() {}
