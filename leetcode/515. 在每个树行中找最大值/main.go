package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	list := make([]*TreeNode, 1)
	list[0] = root
	for len(list) != 0 {
		max := list[0].Val
		tmp := make([]*TreeNode, 0)
		for _, v := range list {
			if v.Left != nil {
				tmp = append(tmp, v.Left)
			}
			if v.Right != nil {
				tmp = append(tmp, v.Right)
			}
			if v.Val > max {
				max = v.Val
			}
		}
		result = append(result, max)
		list = tmp
	}
	return result
}

func main() {

}
