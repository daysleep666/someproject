package main

import "fmt"

// lru 最近最少使用

type OneNode struct {
	Data int
	Next *OneNode
}

const MAXNODESIZE int = 10

var nodeSize int

func InsertToHeadNode(_oneNode *OneNode, _newData int) *OneNode { // 传的是指针的拷贝
	tmpNode := &OneNode{
		Data: _newData,
		Next: _oneNode,
	}
	nodeSize++
	return tmpNode
}

func DisplayNode(_oneNode *OneNode) {
	if _oneNode == nil {
		return
	}
	fmt.Println(_oneNode.Data)
	DisplayNode(_oneNode.Next)
}

func DeleteLastNode(_oneNode *OneNode) *OneNode {
	if _oneNode == nil {
		return nil
	}
	var (
		tmpNode   = _oneNode
		frontNode *OneNode
	)
	for tmpNode.Next != nil {
		frontNode = tmpNode
		tmpNode = tmpNode.Next
	}
	if frontNode == nil {
		return nil
	}
	frontNode.Next = nil
	nodeSize--
	return _oneNode
}

func DeleteNode(_oneNode *OneNode, _needDeleteData int) *OneNode {
	if _oneNode == nil {
		return nil
	}
	for _oneNode.Data == _needDeleteData {
		_oneNode = _oneNode.Next
		nodeSize--
		if _oneNode == nil {
			return nil
		}
	}
	_oneNode.Next = DeleteNode(_oneNode.Next, _needDeleteData)
	return _oneNode
}

func ViewNode(_oneNode *OneNode, _viewValue int) *OneNode {
	// 先在链表里找这个值
	var (
		tmpNode  = _oneNode
		headNode = _oneNode
	)
	if tmpNode != nil {
		for tmpNode.Data != _viewValue {
			tmpNode = tmpNode.Next
			if tmpNode == nil {
				break
			}
		}
	}

	if tmpNode == nil {
		// 1.找不到
		if nodeSize < MAXNODESIZE {
			// 1.1空间未满，插入到头节点
			return InsertToHeadNode(headNode, _viewValue)
		}
		// 1.2空间已满，删除最后一个节点，并插入到头节点
		headNode = DeleteLastNode(headNode)
		return InsertToHeadNode(headNode, _viewValue)
	}

	// 2.找到，就将这个值从当前位置移到头节点
	headNode = DeleteNode(headNode, tmpNode.Data)
	return InsertToHeadNode(headNode, _viewValue)
}

func main() {
	var headNode *OneNode
	for i := 0; i < 10; i++ {
		headNode = ViewNode(headNode, i)
	}
	headNode = ViewNode(headNode, 2)
	headNode = ViewNode(headNode, 11)
	headNode = ViewNode(headNode, 7)
	DisplayNode(headNode)
}
