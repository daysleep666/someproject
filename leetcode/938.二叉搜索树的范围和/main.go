package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	sum := 0
	if root == nil {
		return sum
	}
	if root.Val >= L && root.Val <= R {
		sum += root.Val
	}

	if root.Val <= R {
		sum += rangeSumBST(root.Right, L, R)
	}
	if root.Val >= L {
		sum += rangeSumBST(root.Left, L, R)
	}
	return sum
}

func main() {

}
