package main

import (
	"github.com/daysleep666/someproject/datastruct/singlelist"
)

// 反转一个单链表。

// 示例:

// 输入: 1->2->3->4->5->NULL
// 输出: 5->4->3->2->1->NULL
// 进阶:
// 你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

// 时间复杂度 O(n)
// 空间复杂度 O(1)
func reverseList(_oneNode *singlelist.OneNode) *singlelist.OneNode {
	if _oneNode == nil {
		return nil
	}
	var (
		// 用来缓存前一个节点
		preNode *singlelist.OneNode
		// 指向当前节点
		curNode = _oneNode
		// 缓存当前节点的下一个节点
		nextNode *singlelist.OneNode
	)

	for curNode != nil {
		// 将当前节点的下一个节点保存下
		nextNode = curNode.Next
		// 反转当前节点指向
		curNode.Next = preNode
		// 将当前节点作为下一个节点
		preNode = curNode
		// 将指针指向下一个节点
		curNode = nextNode
	}
	// 如果当前节点是nil，那上一个节点必然是链表最后一个节点
	return preNode
}

func reverseListDIGUI(_oneNode, _preNode *singlelist.OneNode) *singlelist.OneNode {
	if _oneNode == nil {
		return nil
	}
	var nextNode = _oneNode.Next
	_oneNode.Next = _preNode
	var tmpNode = reverseListDIGUI(nextNode, _oneNode)
	if tmpNode == nil {
		return _oneNode
	}
	return tmpNode
}

func main() {
	var headNode *singlelist.OneNode
	for i := 0; i < 10; i++ {
		headNode = singlelist.AddNode(headNode, i)
	}
	headNode = reverseList(headNode)
	headNode = reverseListDIGUI(headNode, nil)
	singlelist.DisplayNode(headNode)

}
