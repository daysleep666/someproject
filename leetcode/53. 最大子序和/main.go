package main

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var cur, max int = nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if cur+nums[i] > nums[i] {
			cur = cur + nums[i]
		} else {
			cur = nums[i]
		}
		if cur > max {
			max = cur
		}
	}
	return max
}

func main() {

}
