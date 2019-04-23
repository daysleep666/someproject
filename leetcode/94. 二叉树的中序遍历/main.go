package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	midOrder(root, &result)
	return result
}

func midOrder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	midOrder(root.Left, result)
	(*result) = append((*result), root.Val)
	midOrder(root.Right, result)
}

func main() {

}
