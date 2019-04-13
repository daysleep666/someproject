package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	list := make([]*TreeNode, 1)
	list[0] = root
	for len(list) > 0 {
		result = append(result, []int{})
		tmpList := make([]*TreeNode, 0)
		for _, v := range list {
			if v.Left != nil {
				tmpList = append(tmpList, v.Left)
			}
			if v.Right != nil {
				tmpList = append(tmpList, v.Right)
			}
			result[len(result)-1] = append(result[len(result)-1], v.Val)
		}
		list = tmpList
	}
	return result
}

func main() {

}
