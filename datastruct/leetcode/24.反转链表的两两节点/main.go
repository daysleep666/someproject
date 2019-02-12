package main

import "github.com/daysleep666/someproject/datastruct/singlelist"

func swap(_oneNode *singlelist.OneNode) *singlelist.OneNode {
	var (
		headNode  *singlelist.OneNode
		frontNode *singlelist.OneNode
		curNode   = _oneNode
		nextNode  *singlelist.OneNode
	)
	if curNode != nil && curNode.Next != nil {
		headNode = curNode
		nextNode = curNode.Next
		curNode.Next = nextNode.Next
		nextNode.Next = curNode
		headNode = nextNode

		frontNode = curNode
		curNode = curNode.Next

	} else {
		return curNode
	}

	for {
		if curNode == nil || curNode.Next == nil {
			return headNode
		}
		nextNode = curNode.Next
		frontNode.Next = nextNode
		curNode.Next = nextNode.Next
		nextNode.Next = curNode

		frontNode = curNode
		curNode = curNode.Next
	}
}

func main() {
	var headNode *singlelist.OneNode
	for i := 1; i < 10; i++ {
		headNode = singlelist.AddNode(headNode, i)
	}
	headNode = swap(headNode)
	singlelist.DisplayNode(headNode)

}
