package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	return d(root, 101)
}

func d(root *TreeNode, min int) int {
	if root.Left != nil {
		node := root.Left
		for node.Right != nil {
			node = node.Right
		}
		tmp := root.Val - node.Val
		if tmp < min {
			min = tmp
		}
		tmp = d(root.Left, min)
		if tmp < min {
			min = tmp
		}
	}
	if root.Right != nil {
		node := root.Right
		for node.Left != nil {
			node = node.Left
		}
		tmp := node.Val - root.Val
		if tmp < min {
			min = tmp
		}
		tmp = d(root.Right, min)
		if tmp < min {
			min = tmp
		}
	}
	return min
}

func main() {

}
