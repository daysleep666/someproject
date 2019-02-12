package problem

// 求一个数的平方根

func F求一个数的平方根(_value int64) {
	value := float64(_value)
	f求一个数的平方根(value, 0, value)
}

func f求一个数的平方根(_value, _low, _high float64) float64 {
	if _low > _high {
		return -1
	}
	mid := (_low + (_high - _low)) / 2
	if mid*mid == _value {
		return mid
	} else if mid*mid < _value {
		return f求一个数的平方根(_value, mid, _high)
	}
	return f求一个数的平方根(_value, _low, mid)
}
