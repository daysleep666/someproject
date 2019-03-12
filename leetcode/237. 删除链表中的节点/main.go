package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	tmpNode := node
	frontNode := node
	for tmpNode != nil && tmpNode.Next != nil {
		tmpNode.Val = tmpNode.Next.Val
		frontNode = tmpNode
		tmpNode = tmpNode.Next
	}
	frontNode.Next = nil
}

func main() {

}
