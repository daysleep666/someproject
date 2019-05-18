package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var head *BPLUSTREEENODE
	for i := 0; i <= 10; i++ {
		// for i := 5; i >= 0; i-- {
		n := rand.Intn(20)
		// for _, n := range []int{49, 40, 36, 39, 90, 50, 51} {
		head = Insert(head, n, 1)
		// var num int
		// head.Symbol(&num)
		// head.DisplayList()
	}
	var num int
	fmt.Println()
	head.Symbol(&num)
	head.Display()
	head.DisplayList()
}

//--------------------------

const M = 3

type BPLUSTREEENODE struct {
	Key      int               // 索引值
	Children []*BPLUSTREEENODE // 子树
	// leaf
	Value     int             // 真正的数据
	FrontNode *BPLUSTREEENODE // 指向相邻的叶子节点
	NextNode  *BPLUSTREEENODE // 指向相邻的叶子节点
	IsLeaf    bool            // 是否是叶子节点
	symbol    int
}

func NewNode() *BPLUSTREEENODE {
	return &BPLUSTREEENODE{Children: make([]*BPLUSTREEENODE, 0)}
}

func (n *BPLUSTREEENODE) InsertNode(key int) *BPLUSTREEENODE {
	newNode := NewNode()
	newNode.Key = key
	for i, v := range n.Children {
		if key < v.Key {
			tmp := make([]*BPLUSTREEENODE, 0)
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

func (n *BPLUSTREEENODE) InsertLeaf(key, value int) *BPLUSTREEENODE {
	newNode := NewNode()
	newNode.Key = key
	newNode.Value = value
	newNode.IsLeaf = true
	for i, v := range n.Children {
		if key < v.Key {
			tmp := make([]*BPLUSTREEENODE, 0)
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

func (n *BPLUSTREEENODE) IsFull() bool {
	return len(n.Children) > M
}

func Insert(head *BPLUSTREEENODE, key, value int) *BPLUSTREEENODE {
	if head == nil {
		head := NewNode()
		head.Key = key
		node := head.InsertNode(key)
		node.Value = value
		node.IsLeaf = true
		return head
	}

	parentNodeList := make([]*BPLUSTREEENODE, 0) // 记录每一层的父节点
	node := head
	isFind := false
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

	leafNode := node.InsertNode(key)
	leafNode.Value = value
	leafNode.IsLeaf = true

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
		newParentNode := parentNodeList[parentIndex]
		newIndex := (M + 1) / 2
		newChildren := node.Children[:newIndex]
		nnn := newParentNode.InsertNode(newChildren[newIndex-1].Key)
		nnn.Children = newChildren
		node.Children = node.Children[newIndex:]
	}

	newHead := NewNode()
	newHead.Key = node.Key
	newIndex := (M + 1) / 2
	newChildren := node.Children[:newIndex]
	nnn := NewNode()
	nnn.Key = newChildren[newIndex-1].Key
	nnn.Children = newChildren
	node.Children = node.Children[newIndex:]
	newHead.Children = append(newHead.Children, nnn, node)
	updateLastMax(newHead, key)
	return newHead
}

func updateLastMax(head *BPLUSTREEENODE, key int) {
	for head != nil && !head.IsLeaf {
		if key > head.Key {
			head.Key = key
		}
		head = head.Children[len(head.Children)-1]
	}
}

func (n *BPLUSTREEENODE) Symbol(num *int) {
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

func (n *BPLUSTREEENODE) Display() {
	fmt.Println("\n", "--------------------")
	list := make([]*BPLUSTREEENODE, 0)
	symbol := n.symbol
	for _, v := range n.Children {
		list = append(list, v)
	}
	for len(list) > 0 {
		tmpList := make([]*BPLUSTREEENODE, 0)
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

func (n *BPLUSTREEENODE) DisplayList() {
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
