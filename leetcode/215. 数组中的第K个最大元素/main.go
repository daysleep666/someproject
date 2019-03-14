package main

func partition(nums []int, low, high int) int {
	last := nums[high]
	j := low
	for i := low; i < high; i++ {
		if nums[i] <= last {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	nums[j], nums[high] = nums[high], nums[j]
	return j
}

func findKthLargest(nums []int, k int) int {
	k--
	low, high := 0, len(nums)-1
	for {
		mid := partition(nums, low, high)
		if mid == k {
			return nums[k]
		} else if mid > k {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

}

func main() {

}
