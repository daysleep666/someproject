package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	result := new(int)
	d(root, &k, result)
	return *result
}

func d(root *TreeNode, k, result *int) {
	if root == nil {
		return
	}
	d(root.Left, k, result)
	(*k)--
	if (*k) == 0 {
		(*result) = root.Val
		return
	}
	d(root.Right, k, result)
	return
}

func main() {
}
