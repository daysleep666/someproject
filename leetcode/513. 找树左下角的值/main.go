package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	var maxLevel, maxVal int = 0, root.Val
	d(root, 0, &maxLevel, &maxVal)
	return maxVal
}

func d(root *TreeNode, curLevel int, maxLevel, maxVal *int) {
	if root == nil {
		return
	}
	if curLevel > (*maxLevel) {
		(*maxLevel) = curLevel
		(*maxVal) = root.Val
	}
	d(root.Left, curLevel+1, maxLevel, maxVal)
	d(root.Right, curLevel+1, maxLevel, maxVal)
}

func main() {

}
