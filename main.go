package main

import (
	"fmt"
	"sort"
)

func bs(data []int, target int) int {
	from, to, mid := 0, len(data)-1, 0
	for from < to {
		mid = (from + to) / 2
		if data[mid] == target {
			to--
		} else if data[mid] < target {
			from = mid + 1
		} else {
			to = mid - 1
		}
	}

	k := from
	if data[k] < target {
		k++
	}
	data = append(data, 0)
	ls := len(data)
	for i := ls - 2; i >= k; i-- {
		data[i+1] = data[i]
	}
	data[k] = target
	fmt.Println(data, k)
	return k
}

func bs2(data []int, target int) int {
	from, to, mid := 0, len(data)-1, 0
	for from < to {
		mid = (from + to) / 2
		if data[mid] == target {
			from++
		} else if data[mid] < target {
			from = mid + 1
		} else {
			to = mid - 1
		}
	}

	k := from
	if data[from] != target {
		k = -1
	}
	fmt.Println(k)
	return k
}

func main() {
	sort.SearchInts()
	(bs2([]int{1, 1, 2, 2}, 4))
	(bs2([]int{1, 1, 3, 3}, 0))
	(bs2([]int{1, 1, 2, 2}, 2))
	(bs2([]int{1, 1, 3, 3}, 2))
	(bs2([]int{1, 3, 5}, 4))
	(bs2([]int{1, 3, 3, 5}, 4))
}
