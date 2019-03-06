package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	a := 0
	d(root, &a)
	return root
}

func d(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	d(root.Right, sum)
	(*sum) += root.Val
	root.Val = (*sum)
	d(root.Left, sum)
}

func main() {

}
