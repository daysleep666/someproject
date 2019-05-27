package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	var list []*TreeNode
	var isLeft bool = true
	list = append(list, root)
	for len(list) != 0 {
		var tmpRes []int
		var tmp []*TreeNode
		for i := 0; i < len(list); i++ {
			v := list[i]
			tmpRes = append(tmpRes, v.Val)
			if isLeft {
				if v.Left != nil {
					tmp = append(tmp, v.Left)
				}
				if v.Right != nil {
					tmp = append(tmp, v.Right)
				}
			} else {
				if v.Right != nil {
					tmp = append(tmp, v.Right)
				}
				if v.Left != nil {
					tmp = append(tmp, v.Left)
				}
			}
		}
		if isLeft {
			revers(list)
		}
		isLeft = !isLeft
		list = tmp
		result = append(result, tmpRes)
	}
	// fmt.Println(result)
	return result
}

func revers(list []*TreeNode) {
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
}

func main() {

}
