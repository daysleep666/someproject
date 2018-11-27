package main

import "fmt"

// 双向链表

type OneNode struct {
	Data     int
	Previous *OneNode
	Next     *OneNode
}

func AddNode(_oneNode *OneNode, _newData int) *OneNode {
	if _oneNode == nil {
		return &OneNode{
			Data:     _newData,
			Previous: nil,
			Next:     nil,
		}
	}
	var tmpNode = _oneNode
	var nextNode = AddNode(tmpNode.Next, _newData)
	tmpNode.Next = nextNode
	nextNode.Previous = tmpNode
	return tmpNode
}

func DeleteNode(_oneNode *OneNode, _newData int) *OneNode {
	if _oneNode == nil {
		return nil
	}
	var tmpNode = _oneNode
	for tmpNode.Data == _newData {
		tmpNode = tmpNode.Next
		if tmpNode == nil {
			return nil
		}
	}
	var nextNode = DeleteNode(tmpNode.Next, _newData)
	tmpNode.Next = nextNode
	if nextNode != nil {
		nextNode.Previous = tmpNode
	}
	return tmpNode
}

func DisplayNode(_oneNode *OneNode) {
	if _oneNode == nil {
		return
	}
	fmt.Printf("cur=%v", _oneNode.Data)
	if _oneNode.Previous != nil {
		fmt.Printf(", previous=%v", _oneNode.Previous.Data)
	}
	if _oneNode.Next != nil {
		fmt.Printf(", next=%v", _oneNode.Next.Data)
	}
	fmt.Println()
	DisplayNode(_oneNode.Next)
}

func main() {
	var headNode *OneNode
	for i := 0; i < 10; i++ {
		headNode = AddNode(headNode, i)
	}
	headNode = DeleteNode(headNode, 2)
	DisplayNode(headNode)
}
