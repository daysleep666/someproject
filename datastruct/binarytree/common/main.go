package main

import (
	"fmt"
	"math/rand"

	"github.com/daysleep666/someproject/datastruct/queue/queuesturct"
)

type BinaryTree struct {
	Data      int64
	LeftNode  *BinaryTree
	RightNode *BinaryTree
}

func Insert(_node *BinaryTree, _data int64) *BinaryTree {
	if _node == nil {
		_node = &BinaryTree{Data: _data}
		return _node
	}
	if _data < _node.Data {
		_node.LeftNode = Insert(_node.LeftNode, _data)
	} else {
		_node.RightNode = Insert(_node.RightNode, _data)
	}
	return _node
}

func PreOrder(_node *BinaryTree) { // 中左右
	if _node == nil {
		return
	}
	fmt.Println(_node.Data)
	PreOrder(_node.LeftNode)
	PreOrder(_node.RightNode)
}

func MidOrder(_node *BinaryTree) { // 左中右
	if _node == nil {
		return
	}
	PreOrder(_node.LeftNode)
	fmt.Println(_node.Data)
	PreOrder(_node.RightNode)
}

func PosrtOrder(_node *BinaryTree) { // 左中右
	if _node == nil {
		return
	}
	PreOrder(_node.LeftNode)
	PreOrder(_node.RightNode)
	fmt.Println(_node.Data)
}

func Order层序(_node *BinaryTree) { // 层序
	if _node == nil {
		return
	}
	q := queuesturct.NewQueue()
	q.Push(_node)
	for q.Length() > 0 {
		value, _ := q.Pop()
		tmpNode := value.(*BinaryTree)
		if tmpNode == nil {
			fmt.Printf("*  ")
			continue
		}
		fmt.Printf("%v  ", tmpNode.Data)
		q.Push(tmpNode.LeftNode, tmpNode.RightNode)
	}

}

func main() {
	var head *BinaryTree
	for i := int64(0); i < 10; i++ {
		data := rand.Int63n(100)
		head = Insert(head, data)
	}
	Order层序(head)
}
