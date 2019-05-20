package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var head *BPlusNode
	for i := 0; i < 10; i++ {
		n := rand.Intn(100)
		head = Insert(head, n, i)
	}
	var num int
	fmt.Println()
	head.Symbol(&num)
	head.Display()
	head.DisplayList()
}

//--------------------------

const M = 3
const newIndex = (M + 1) / 2 // 拆分的时候以这个值为准

type BPlusNode struct {
	Key      int          // 索引值
	Children []*BPlusNode // 子树
	// leaf
	Value     int        // 真正的数据
	FrontNode *BPlusNode // 指向相邻的叶子节点
	NextNode  *BPlusNode // 指向相邻的叶子节点
	IsLeaf    bool       // 是否是叶子节点
	symbol    int
}

func NewNode() *BPlusNode {
	return &BPlusNode{Children: make([]*BPlusNode, 0)}
}

func (n *BPlusNode) InsertNodeWithChildren(key int, children []*BPlusNode) *BPlusNode {
	newNode := NewNode()
	newNode.Key = key
	newNode.Children = children
	for i, v := range n.Children {
		if key < v.Key {
			tmp := make([]*BPlusNode, 0)
			tmp = append(tmp, n.Children[:i]...)
			tmp = append(tmp, newNode)
			tmp = append(tmp, n.Children[i:]...)
			n.Children = tmp
			return newNode
		}
	}
	n.Children = append(n.Children, newNode)
	return newNode
}

func (n *BPlusNode) InsertLeaf(key, value int) *BPlusNode {
	newNode := NewNode()
	newNode.Key = key
	newNode.Value = value
	newNode.IsLeaf = true
	for i, v := range n.Children {
		if key < v.Key {
			tmp := make([]*BPlusNode, 0)
			tmp = append(tmp, n.Children[:i]...)
			tmp = append(tmp, newNode)
			tmp = append(tmp, n.Children[i:]...)
			newNode.FrontNode = n.Children[i].FrontNode
			newNode.NextNode = n.Children[i]
			if n.Children[i].FrontNode != nil {
				n.Children[i].FrontNode.NextNode = newNode
			}
			n.Children[i].FrontNode = newNode

			n.Children = tmp
			return newNode
		}
	}
	if len(n.Children) > 0 {
		newNode.NextNode = n.Children[len(n.Children)-1].NextNode
		if newNode.NextNode != nil {
			newNode.NextNode.FrontNode = newNode
		}
		n.Children[len(n.Children)-1].NextNode = newNode
		newNode.FrontNode = n.Children[len(n.Children)-1]
	}
	n.Children = append(n.Children, newNode)
	return newNode
}

func (n *BPlusNode) IsFull() bool {
	return len(n.Children) > M
}

func Insert(head *BPlusNode, key, value int) *BPlusNode {
	if head == nil {
		head := NewNode()
		head.InsertLeaf(key, value)
		return head
	}

	parentNodeList := make([]*BPlusNode, 0) // 记录每一层的父节点
	node := head
	isFind := false

	// 找到插入的位置
	for !isFind {
		parentNodeList = append(parentNodeList, node)
		var i, l int = 0, len(node.Children)
		for ; i < len(node.Children); i++ {
			child := node.Children[i]
			if child.IsLeaf { // 应该插在这个节点下面
				isFind = true
				break
			}
			if key <= child.Key { // 往下接着找
				node = child
				break
			}
		}
		if i == l {
			node = node.Children[l-1]
		}
	}

	// 插入叶子节点
	node.InsertLeaf(key, value)

	// 如果 Children超过规定的M，就需要拆分节点
	parentIndex := len(parentNodeList) - 1
	for {
		node = parentNodeList[parentIndex]
		if !node.IsFull() {
			updateLastMax(head, key)
			return head
		}
		parentIndex--
		if parentIndex < 0 {
			break
		}

		// 假设M=3，将[1,2,3,4]拆分为[1,2],[3,4]
		newParentNode := parentNodeList[parentIndex]
		newChildren := node.Children[:newIndex]
		newParentNode.InsertNodeWithChildren(newChildren[newIndex-1].Key, newChildren)
		node.Children = node.Children[newIndex:]
	}

	newHead := NewNode()
	newHead.Key = node.Key
	newChildren := node.Children[:newIndex]
	newHead.InsertNodeWithChildren(newChildren[newIndex-1].Key, newChildren)

	node.Children = node.Children[newIndex:]
	newHead.InsertNodeWithChildren(node.Children[len(node.Children)-1].Key, node.Children)
	updateLastMax(newHead, key)
	return newHead
}

func updateLastMax(head *BPlusNode, key int) {
	for head != nil && !head.IsLeaf {
		if key > head.Key {
			head.Key = key
		}
		head = head.Children[len(head.Children)-1]
	}
}

func (n *BPlusNode) Symbol(num *int) {
	if len(n.Children) != 0 {
		(*num)++
	}
	for _, v := range n.Children {
		v.symbol = (*num)
	}
	for _, v := range n.Children {
		v.Symbol(num)
	}
}

func (n *BPlusNode) Display() {
	fmt.Println("\n", "--------------------")
	list := make([]*BPlusNode, 0)
	symbol := n.symbol
	for _, v := range n.Children {
		list = append(list, v)
	}
	for len(list) > 0 {
		tmpList := make([]*BPlusNode, 0)
		for _, v := range list {
			if v.symbol != symbol {
				symbol = v.symbol
				fmt.Printf("}   { ")
			}
			if v.IsLeaf {
				fmt.Printf("%v[%v] ", v.Key, v.Value)
			} else {
				fmt.Printf("%v ", v.Key)
			}
			tmpList = append(tmpList, v.Children...)
		}
		list = tmpList
		fmt.Printf("}")
		fmt.Println()
	}
	fmt.Println("\n", "--------------------")
}

func (n *BPlusNode) DisplayList() {
	fmt.Println("\n", "--------------------")
	node := n
	for !node.IsLeaf {
		node = node.Children[0]
	}
	lastNode := node
	for node != nil {
		fmt.Printf("%v->", node.Key)
		lastNode = node
		node = node.NextNode
	}
	fmt.Println()
	node = lastNode
	for node != nil {
		fmt.Printf("%v->", node.Key)
		node = node.FrontNode
	}
	fmt.Println("\n", "--------------------")

}
