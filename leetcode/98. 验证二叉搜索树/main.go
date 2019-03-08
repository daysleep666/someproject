package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	arr := d(root)
	fmt.Println(arr)
	cur := arr[0]
	for i := 1; i < len(arr); i++ {
		if cur >= arr[i] {
			return false
		}
		cur = arr[i]
	}
	return true
}

func d(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	arr := make([]int, 0)
	arr = append(arr, d(root.Left)...)
	arr = append(arr, root.Val)
	arr = append(arr, d(root.Right)...)
	return arr
}

func isValidBST(root *TreeNode) bool {
	return d2(root, math.MinInt64, math.MaxInt64)
}

func d2(root *TreeNode, min, max int) bool { // 对于左子树来说我永远比传来的值小，对于右子树来说我永远比传来的值大
	if root == nil {
		return true
	}
	return min < root.Val && root.Val < max && d2(root.Left, min, root.Val) && d2(root.Right, root.Val, max)
}

func a(arr []int) {
	arr[0] = 1
	arr = append(arr, 2)
}

func main() {
	arr := []int{2}
	a(arr)
	fmt.Println(arr)
}
