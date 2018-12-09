package main

import (
	"fmt"
)

// 哈希函数 time33
// 解决冲突 链地址法
// resize

const (
	MAXLOADINGFACTOR = 0.75 // 如果装填因子大于0.75，就需要对哈希表进行扩容
)

// 单链表
type oneNode struct {
	Key   string
	Value int
	Next  *oneNode
}

func (on *oneNode) AddNode(_key string, _newData int) {
	var tmpNode = on
	for tmpNode.Next != nil {
		tmpNode = tmpNode.Next
	}

	tmpNode.Next = &oneNode{Key: _key, Value: _newData}
	on = nil
	return
}

func (on *oneNode) DeleteNode(_key string) {
	var tmpNode = on
	var frontNode = tmpNode
	for tmpNode != nil {
		for tmpNode.Key == _key {
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

func (on *oneNode) Find(_key string) (int, bool) {
	var tmpNode = on
	for tmpNode != nil {
		for tmpNode.Key == _key {
			return tmpNode.Value, true
		}
		tmpNode = tmpNode.Next
	}
	return 0, false
}

type myMap struct {
	HashTable []*oneNode
	MaxSize   uint // 最大长度
	CurSize   uint // 当前大小
}

func NewMap(_maxSize uint) *myMap {
	if _maxSize == 0 {
		panic("maxsize must more than zero")
	}

	var realMaxSize = uint(float32(_maxSize) / MAXLOADINGFACTOR)

	return &myMap{
		HashTable: make([]*oneNode, realMaxSize),
		MaxSize:   realMaxSize,
	}
}

func (m *myMap) hash(_key string) int {
	var hash int
	for _, v := range _key {
		hash += int(v) * 33
	}
	return (hash & 0x7FFFFFFF) % int(m.MaxSize)
}

func (m *myMap) GetValue(_key string) (int, bool) {
	hash := m.hash(_key)
	node := m.HashTable[hash]
	if node == nil {
		return 0, false
	}
	return node.Find(_key)
}

func (m *myMap) SetValue(_key string, _value int) {
	hash := m.hash(_key)
	node := m.HashTable[hash]
	for node != nil {
		if node.Key == _key {
			node.Value = _value
			return
		}
	}
	m.HashTable[hash] = new(oneNode)
	m.HashTable[hash].AddNode(_key, _value)
}

func (m *myMap) DelValue(_key string) {
	hash := m.hash(_key)
	node := m.HashTable[hash]
	if node == nil {
		return
	}
	node.DeleteNode(_key)
}

func main() {
	var m = NewMap(10)
	for i := 0; i < 11; i++ {
		m.SetValue(fmt.Sprintf("%v", i), i)
	}
	for i := 0; i < 11; i++ {
		fmt.Println(m.GetValue(fmt.Sprintf("%v", i)))
	}
}
