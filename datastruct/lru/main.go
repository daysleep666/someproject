package main

import "fmt"

// lru 最近最少使用

type OneNode struct {
	Data int
	Next *OneNode
}

const MAXNODESIZE int = 10

var nodeSize int

func AddNode(_oneNode *OneNode, _newData int) *OneNode { // 传的是指针的拷贝
	fmt.Printf("%v\n", &_oneNode)
	if _oneNode == nil {
		_oneNode = new(OneNode)
		_oneNode.Data = _newData
		nodeSize++
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
	// 1.找不到
	// 1.1空间未满，插入到头节点
	if _oneNode == nil {
		return AddNode(_oneNode, _viewValue)
	}

	var tmpNode = _oneNode
	for tmpNode.Data != _viewValue {
		tmpNode = tmpNode.Next
		if tmpNode == nil {
			break
		}
	}

	if tmpNode == nil && nodeSize < MAXNODESIZE {
		return AddNode(_oneNode, _viewValue)
	}

	// 1.2空间已满，删除最后一个节点，并插入到头节点

	// 2.找到，就将这个值从当前位置移到头节点
	return nil
}

func main() {

}
