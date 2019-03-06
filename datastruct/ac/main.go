package main

import "fmt"

type Node struct {
	Data       rune
	ChildNode  map[rune]*Node
	FailedNode *Node
}

func newNode() *Node {
	return &Node{ChildNode: make(map[rune]*Node)}
}

func (n *Node) Push(str string) {
	node := n
	tmpNode := n
	isExist := false
	for _, v := range str {
		tmpNode, isExist = node.ChildNode[v]
		if !isExist {
			tmpNode = newNode()
			tmpNode.Data = v
			node.ChildNode[v] = tmpNode
		}
		node = tmpNode
	}
}

func (n *Node) BuildAcAutoMation() {

}

func (n *Node) Display() {
	for _, v := range n.ChildNode {
		fmt.Println(string(v.Data))
		v.Display()
	}
}

func main() {
	n := newNode()
	n.Push("hello")
	n.Push("hi")
	n.Display()
}
