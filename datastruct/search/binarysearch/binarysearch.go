package binarysearch

func BinarySearch(_arr []int64, _value int64) int64 {
	return binarySearch(_arr, _value, 0, int64(len(_arr)-1))
}

func binarySearch(_arr []int64, _value int64, _low, _high int64) int64 {
	if _low > _high {
		return -1
	}
	mid := (_low + (_high - _low)) >> 1 // 不直接写 (high+low)/2 是为了防止溢出
	if _value == _arr[mid] {
		return mid
	} else if _value > _arr[mid] {
		return binarySearch(_arr, _value, mid+1, _high)
	}
	return binarySearch(_arr, _value, _low, mid-1)
}

// 时间复杂度 O(logn)

// 空间复杂度 O(1)
