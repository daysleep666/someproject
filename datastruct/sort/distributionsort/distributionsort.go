package distributionsort

// 基数排序

func DistributionSort(_arr []int64) {
	var (
		max int64
	)

	for _, v := range _arr {
		if v > max {
			max = v
		}
	}

	for i := int64(1); i <= max; i *= 10 {
		bubble := make([][]int64, 10)
		for _, v := range _arr { // n次
			tmpV := v / i
			index := tmpV % 10
			bubble[index] = append(bubble[index], v)
		}
		var k int64
		for _, arr := range bubble { // b次
			for _, v := range arr {
				_arr[k] = v
				k++
			}
		}
	}
}

// 时间复杂度 O(n*log(r)m)

// 空间复杂度 O(1)

// 是稳定的排序算法
