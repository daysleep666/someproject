package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var max int

func maxPathSum(root *TreeNode) int {
	max = root.Val
	d(root)
	return max
}

func d(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := d(root.Left)
	r := d(root.Right)

	// 作为主干
	mainV := root.Val
	mainV = getMax(mainV, mainV+l)
	mainV = getMax(mainV, mainV+r)
	max = getMax(mainV, max)

	// 作为分支
	return getMax(root.Val, root.Val+l, root.Val+r)
}

func getMax(nums ...int) int {
	m := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > m {
			m = nums[i]
		}
	}
	return m
}

func main() {

}
