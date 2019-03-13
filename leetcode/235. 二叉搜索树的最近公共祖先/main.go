package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	tmp := root
	for tmp != nil {
		if v := compare(tmp, p, q); v == 0 {
			return tmp
		} else if v < 0 { // 在左边
			tmp = tmp.Left
		} else { // 在右边
			tmp = tmp.Right
		}
	}
	return tmp
}

func compare(v, p, q *TreeNode) int {
	if p.Val > v.Val && q.Val > v.Val {
		return 1
	} else if p.Val < v.Val && q.Val < v.Val {
		return -1
	}
	return 0
}

func main() {

}
