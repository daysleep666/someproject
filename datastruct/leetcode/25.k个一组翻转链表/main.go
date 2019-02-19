package main

import (
	"github.com/daysleep666/someproject/datastruct/singlelist"
)

func reverseKGroup(head *singlelist.OneNode, k int) *singlelist.OneNode {
	var (
		headNode     *singlelist.OneNode
		frontNode    *singlelist.OneNode
		bigFrontNode *singlelist.OneNode
		i            int = 0

		reverseHeadNode = head
		reverseLastNode = head
	)
	for reverseLastNode != nil {
		i++
		frontNode = reverseLastNode.Next
		if i%k == 0 { // 需要翻转了
			if headNode == nil {
				headNode = reverseLastNode
			}
			continueNode := reverseHeadNode
			for reverseHeadNode != reverseLastNode {
				tmpNode := reverseHeadNode.Next
				reverseHeadNode.Next = frontNode
				frontNode = reverseHeadNode
				reverseHeadNode = tmpNode
			}
			reverseLastNode.Next = frontNode
			if bigFrontNode != nil {
				bigFrontNode.Next = reverseHeadNode
			}
			bigFrontNode = continueNode
			reverseLastNode = continueNode.Next
			reverseHeadNode = continueNode.Next
			continue
		}
		reverseLastNode = reverseLastNode.Next
	}
	if headNode == nil {
		return head
	}
	return headNode
}

func main() {
	var headNode *singlelist.OneNode
	for i := 1; i <= 4; i++ {
		headNode = singlelist.AddNode(headNode, i)
	}
	headNode = reverseKGroup递归(headNode, 2)
	singlelist.DisplayNode(headNode)
}

func reverseKGroup递归(head *singlelist.OneNode, k int) *singlelist.OneNode {
	tmpNode := head
	var frontNode *singlelist.OneNode
	for i := 0; i < k; i++ { // 检查够不够k长度
		if tmpNode == nil {
			return head
		}
		tmpNode = tmpNode.Next
		frontNode = tmpNode
	}
	tmpNode = head
	for i := 0; i < k; i++ {
		nextNode := tmpNode.Next
		tmpNode.Next = frontNode
		frontNode = tmpNode
		tmpNode = nextNode
	}
	head.Next = reverseKGroup递归(head.Next, k)
	return frontNode
}
