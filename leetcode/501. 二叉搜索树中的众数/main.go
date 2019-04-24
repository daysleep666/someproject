package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//中序遍历

func findMode(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	var max, front, count int
	d(root, &front, &count, &max, &result)
	return result
}

func d(root *TreeNode, front *int, count *int, max *int, arr *[]int) {
	if root == nil {
		return
	}
	d(root.Left, front, count, max, arr)

	if root.Val == (*front) {
		(*count)++
	} else {
		(*front) = root.Val
		(*count) = 1
	}

	if (*count) > (*max) {
		(*arr) = []int{root.Val}
		(*max) = (*count)
	} else if (*count) == (*max) {
		(*arr) = append((*arr), root.Val)
	}
	d(root.Right, front, count, max, arr)
}

func main() {

}
