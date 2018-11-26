package main

import (
	"fmt"
)

// 单链表

type OneNode struct {
	Data int
	Next *OneNode
}

func AddNode(_oneNode *OneNode, _newData int) *OneNode { // 传的是指针的拷贝
	fmt.Printf("%v\n", &_oneNode)
	if _oneNode == nil {
		_oneNode = new(OneNode)
		_oneNode.Data = _newData
		return _oneNode
	}

	_oneNode.Next = AddNode(_oneNode.Next, _newData)
	return _oneNode
}

func DisplayNode(_oneNode *OneNode) {
	if _oneNode == nil {
		return
	}
	fmt.Println(_oneNode.Data)
	DisplayNode(_oneNode.Next)
}

func DeleteNode(_oneNode *OneNode, _needDeleteData int) *OneNode {
	if _oneNode == nil {
		return nil
	}
	for _oneNode.Data == _needDeleteData {
		_oneNode = _oneNode.Next
		if _oneNode == nil {
			return nil
		}
	}
	_oneNode.Next = DeleteNode(_oneNode.Next, _needDeleteData)
	return _oneNode
}

func (oneNode *OneNode) AddNode(_newData int) {
	var tmpNode = oneNode
	for tmpNode.Next != nil {
		tmpNode = tmpNode.Next
	}

	tmpNode.Next = &OneNode{Data: _newData}
	oneNode = nil
	return
}

func (oneNode *OneNode) DisplayNode() {
	var tmpNode = oneNode
	for tmpNode != nil {
		fmt.Println(tmpNode.Data)
		tmpNode = tmpNode.Next
	}
}

func (oneNode *OneNode) DeleteNode(_needDeleteData int) {
	var tmpNode = oneNode
	var frontNode = tmpNode
	for tmpNode != nil {
		for tmpNode.Data == _needDeleteData {
			frontNode.Next = tmpNode.Next
			tmpNode = tmpNode.Next
			if tmpNode == nil {
				return
			}
		}
		frontNode = tmpNode
		tmpNode = tmpNode.Next
	}
}

func main() {
	// 递归
	// var headNode *OneNode = new(OneNode)
	// for i := 0; i < 1; i++ {
	// 	// 需要在外部赋值是因为传进函数里的是指针拷贝，换句话说，headNode的地址在函数内的和函数外是不一样的，所以在函数内的headNode=new(OneNode)是不会影响到函数外的headNode
	// 	headNode = AddNode(headNode, i)
	// 	fmt.Printf("%v\n", &headNode)
	// }
	// headNode = DeleteNode(headNode, 2)
	// DisplayNode(headNode)

	// 非递归
	var headNode *OneNode = new(OneNode)
	for i := 0; i < 10; i++ {
		headNode.AddNode(i)
	}
	headNode.DeleteNode(2)
	headNode.DisplayNode()
}
