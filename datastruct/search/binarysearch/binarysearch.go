package binarysearch

func BinarySearch(_arr []int64, _value int64) int64 {
	return binarySearch(_arr, _value, 0, int64(len(_arr)-1))
}

func binarySearch(_arr []int64, _value int64, _low, _high int64) int64 {
	if _low > _high {
		return -1
	}
	mid := _low + (_high-_low)>>1 // 不直接写 (high+low)/2 是为了防止溢出
	if _value == _arr[mid] {
		return mid
	} else if _value > _arr[mid] {
		return binarySearch(_arr, _value, mid+1, _high)
	}
	return binarySearch(_arr, _value, _low, mid-1)
}

// 时间复杂度 O(logn)

// 空间复杂度 O(1)

// 变体
// 第一个等于值的

func BinarySearch第一个等于值的(_arr []int64, _value int64) int64 {
	var (
		length       = int64(len(_arr))
		low    int64 = 0
		high   int64 = length - 1
	)

	for low <= high {
		mid := low + (high-low)>>1
		if _arr[mid] >= _value {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	if low < length && _arr[low] == _value {
		return low
	}
	return -1
}

func BinarySearch最后一个等于值的(_arr []int64, _value int64) int64 {
	var (
		length       = int64(len(_arr))
		low    int64 = 0
		high   int64 = length - 1
	)

	for low <= high {
		mid := low + (high-low)>>1
		if _arr[mid] > _value {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	if high >= 0 && _arr[high] == _value {
		return high
	}
	return -1
}

func BinarySearch第一个大于等于给定值的(_arr []int64, _value int64) int64 {
	var (
		length       = int64(len(_arr))
		low    int64 = 0
		high   int64 = length - 1
	)

	for low <= high {
		mid := low + (high-low)>>1
		if _arr[mid] >= _value {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	if low < length && _arr[low] >= _value {
		return low
	}
	return -1
}

func BinarySearch循环有序数组(_arr []int64, _value int64) int64 {
	if len(_arr) == 0 {
		return -1
	}
	var (
		length       = int64(len(_arr))
		low    int64 = 0
		high   int64 = length - 1
	)

	for low <= high {
		first := _arr[low]
		mid := low + (high-low)>>1
		if _value == _arr[mid] {
			return mid
		} else if _value < _arr[mid] && _value >= first {
			high = mid - 1
		} else if _value < _arr[mid] && _value < first {
			low = mid + 1
		} else if _value > _arr[mid] && _value >= first {
			low = mid + 1
		} else if _value > _arr[mid] && _value < first {
			high = mid - 1

		}
	}
	return -1
}
