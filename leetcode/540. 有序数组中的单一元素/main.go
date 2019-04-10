package main

func singleNonDuplicate(nums []int) int {
	for i := 1; i < len(nums); i = i + 2 {
		if nums[i] != nums[i-1] {
			return nums[i-1]
		}
	}
	return nums[len(nums)-1]
}

func main() {

}
