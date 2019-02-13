package main

import "github.com/daysleep666/someproject/datastruct/singlelist"

func reverseKGroup(_oneNode *singlelist.OneNode, k int) *singlelist.OneNode {
	var (
		headNode  *singlelist.OneNode
		frontNode *singlelist.OneNode
		tmpNode       = _oneNode
		i         int = 1

		reverseHeadNode = _oneNode
		reverseLastNode = _oneNode
	)
	for {
		reverseLastNode = reverseLastNode.Next
		if i%k == 0 {

		}
	}
}

func reverse(_oneNode *singlelist.OneNode) (*singlelist.OneNode, *singlelist.OneNode) {
	var (
		headNode *singlelist.OneNode
		lastNode = _oneNode

		tmpNode   = _oneNode
		frontNode *singlelist.OneNode
	)
	for tmpNode != nil {
		headNode = tmpNode
		nextNode := tmpNode.Next
		tmpNode.Next = frontNode
		frontNode = tmpNode
		tmpNode = nextNode
	}

	return headNode, lastNode
}

func main() {
	// var headNode *singlelist.OneNode
	// for i := 1; i < 10; i++ {
	// 	headNode = singlelist.AddNode(headNode, i)
	// }
	// headNode = swap(headNode)
	// singlelist.DisplayNode(headNode)

}
