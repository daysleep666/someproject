package main

import "fmt"

type TrieTree struct {
	Char  rune
	Nodes map[rune]*TrieTree
}

func NewTrieTreeNode(_char rune) *TrieTree {
	return &TrieTree{Char: _char, Nodes: make(map[rune]*TrieTree)}
}

func Add(_tt *TrieTree, _rs []rune) *TrieTree {
	if len(_rs) == 0 {
		return nil
	}
	var char = _rs[0]
	if _tt == nil {
		_tt = NewTrieTreeNode(char)
	}
	_tt.Nodes[char] = Add(_tt.Nodes[char], _rs[1:])
	return _tt
}

func Find(_tt *TrieTree, _rs []rune) bool {
	if len(_rs) == 0 {
		return true
	}
	if _tt == nil {
		return false
	}
	var char = _rs[0]
	node, isExist := _tt.Nodes[char]
	if !isExist {
		return false
	}
	return Find(node, _rs[1:])
}

func main() {
	head := &TrieTree{Nodes: make(map[rune]*TrieTree)}
	head = Add(head, []rune("hello world"))
	head = Add(head, []rune("wewqexzcxzc"))
	head = Add(head, []rune("fjksadneqw"))
	fmt.Println(Find(head, []rune("wewqexzcxzc")))
}
