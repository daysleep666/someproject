package main

import (
	"fmt"
	"math/rand"

	"github.com/daysleep666/someproject/datastruct/queue/queuesturct"
)

type BinaryTree struct {
	Data      int64
	Height    int64 // 树高
	LeftNode  *BinaryTree
	RightNode *BinaryTree
}

func getHeight(_b *BinaryTree) int64 {
	if _b != nil {
		return _b.Height
	}
	return 0
}

func max(_a, _b int64) int64 {
	if _a > _b {
		return _a
	}
	return _b
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

	// 判断左子树比右子树高2
	if getHeight(_node.LeftNode)-getHeight(_node.RightNode) == 2 { // 左高于右
		if getHeight(_node.LeftNode.LeftNode) > getHeight(_node.LeftNode.RightNode) { // ll
			return rr(_node)
		} else { // lr
			return lr(_node)
		}
	} else if getHeight(_node.LeftNode)-getHeight(_node.RightNode) == -2 {
		if getHeight(_node.RightNode.RightNode) > getHeight(_node.RightNode.LeftNode) { // rr
			return ll(_node)
		} else { // rl
			return rl(_node)
		}
	}

	_node.Height = max(getHeight(_node.LeftNode), getHeight(_node.RightNode)) + 1
	return _node
}

func rr(_node *BinaryTree) *BinaryTree {
	tmpNode := _node.LeftNode
	_node.LeftNode = tmpNode.RightNode
	tmpNode.RightNode = _node
	_node.Height = max(getHeight(_node.LeftNode), getHeight(_node.RightNode)) + 1
	tmpNode.Height = max(getHeight(tmpNode.LeftNode), getHeight(tmpNode.RightNode)) + 1
	return tmpNode
}

func rl(_node *BinaryTree) *BinaryTree {
	tmpNode := rr(_node.RightNode)
	_node.RightNode = tmpNode
	return ll(_node)
}

func ll(_node *BinaryTree) *BinaryTree {
	tmpNode := _node.RightNode
	_node.RightNode = tmpNode.LeftNode
	tmpNode.LeftNode = _node
	_node.Height = max(getHeight(_node.LeftNode), getHeight(_node.RightNode)) + 1
	tmpNode.Height = max(getHeight(tmpNode.LeftNode), getHeight(tmpNode.RightNode)) + 1
	return tmpNode
}

func lr(_node *BinaryTree) *BinaryTree {
	tmpNode := ll(_node.LeftNode)
	_node.LeftNode = tmpNode
	return rr(_node)
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
	curLast := _node
	nextLast := _node
	for q.Length() > 0 {
		value, _ := q.Pop()
		tmpNode := value.(*BinaryTree)
		if tmpNode == nil {
			fmt.Printf("*  ")
			continue
		}
		fmt.Printf("%v  ", tmpNode.Data)
		if tmpNode.LeftNode != nil {
			q.Push(tmpNode.LeftNode)
			nextLast = tmpNode.LeftNode
		}
		if tmpNode.RightNode != nil {
			q.Push(tmpNode.RightNode)
			nextLast = tmpNode.RightNode
		}
		if curLast == tmpNode {
			fmt.Println()
			curLast = nextLast
		}
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
