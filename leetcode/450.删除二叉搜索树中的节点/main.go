package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key { //找到其左子树的最大节点
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}

		tmpNode := root.Left
		if tmpNode.Right == nil { //
			tmpNode.Right = root.Right
			return tmpNode
		}

		var frontNode *TreeNode
		for tmpNode.Right != nil {
			frontNode, tmpNode = tmpNode, tmpNode.Right
		}

		frontNode.Right, tmpNode.Left, tmpNode.Right = tmpNode.Left, root.Left, root.Right
		return tmpNode
	}
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		root.Left = deleteNode(root.Left, key)
	}
	return root
}

func main() {

}
