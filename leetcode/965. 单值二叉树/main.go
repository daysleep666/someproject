package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	val := root.Val
	list := []*TreeNode{root}
	for len(list) != 0 {
		tmpList := make([]*TreeNode, 0)
		for _, v := range list {
			if v.Val != val {
				return false
			}
			if v.Left != nil {
				tmpList = append(tmpList, v.Left)
			}
			if v.Right != nil {
				tmpList = append(tmpList, v.Right)
			}
		}
		list = tmpList
	}
	return true
}

func main() {

}
