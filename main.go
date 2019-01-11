package main

import "fmt"

func main() {
	originPrice := 100
	price := int64(float64(originPrice) * 1.0)
	fmt.Println(price)

}

var onePage int64 = 15

func queryAllWaitBuyTrade(_page int64) ([]int, int64, int64) {
	nums := make([]int, 100, 100)
	size := int64(len(nums))
	if size == 0 {
		return []int{}, 0, 0
	}
	maxPage := size / 15
	page := _page
	if page > maxPage {
		page = maxPage
	}
	min := page * onePage
	max := (page + 1) * onePage
	if max > size {
		max = size
	}

	return nums[min:max], page, size
}
