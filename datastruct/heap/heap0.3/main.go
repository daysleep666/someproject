package main

func heap(nums []int, l int) {
	for i := l - 1; i > 0; i-- {
		down(nums, i)
	}
}

func down(nums []int, from int) {
	for from > 0 {
		parent := (from - 1) / 2
		if nums[from] < nums[parent] {
			nums[from], nums[parent] = nums[parent], nums[from]
		} else {
			return
		}
		from = parent
	}
}
