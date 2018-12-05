package main

import (
	"fmt"
	"math/rand"
	"time"
)

type OneNode struct {
	Data      int
	Next      *OneNode   // 当前层的下一个节点
	OtherSelf []*OneNode // 其他层的自己
}

type SkipList struct {
	Level    int      // 最高层数
	HeadNode *OneNode // 头节点
}

func InsertNode(_oneNode *OneNode, _newData int) *OneNode {
	if _oneNode == nil {
		_oneNode = new(OneNode)
		_oneNode.Data = _newData
		return _oneNode
	}

	if _oneNode.Data <= _newData {
		_oneNode.Next = InsertNode(_oneNode.Next, _newData)

	} else {
		tmpNode := new(OneNode)
		tmpNode.Data = _newData
		tmpNode.Next = _oneNode
		_oneNode = tmpNode
	}

	return _oneNode
}

func DisplayNode(_oneNode *OneNode) {
	if _oneNode == nil {
		return
	}
	fmt.Println(_oneNode.Data)
	DisplayNode(_oneNode.Next)
}

func NeedToImprove() bool {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	randValue := r.Intn(2)
	return randValue == 0
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

func main() {
	var headNode *OneNode
	for i := 0; i < 20; i++ {
		headNode = InsertNode(headNode, rand.Intn(100))
	}
	DisplayNode(headNode)
}
