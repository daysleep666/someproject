package main

import (
	"fmt"
	"math/rand"
	"time"
)

type OneNode struct {
	Data      int
	NextNodes []*OneNode // 每一层的下一个节点
}

type SkipList struct {
	Level    int      // 最高层数
	HeadNode *OneNode // 头节点
}

func NewSkipList() *SkipList {
	return &SkipList{
		Level:    1,
		HeadNode: &OneNode{NextNodes: make([]*OneNode, 2)},
	}
}

// 由小到大
func (skipList *SkipList) InsertNode(_newData int) {
	var (
		level       = skipList.Level
		frontNode   = skipList.HeadNode
		tmpNode     *OneNode
		updateNodes = make([]*OneNode, skipList.Level+1) // 记录
	)

	// 先找到插入的位置， 记录每层向下移动的点
	for i := level; i >= 0; i-- {
		tmpNode = frontNode.NextNodes[i]
		for tmpNode != nil && tmpNode.Data < _newData { // 在本层找到符合条件的点
			frontNode = tmpNode
			tmpNode = tmpNode.NextNodes[i]
		}
		updateNodes[i] = frontNode
	}

	// 计算需要提升几层
	var willImproveLevel int
	for i := 0; i <= level; i++ {
		if !needToImprove() {
			// 不需要提升
			break
		}
		willImproveLevel = i
	}

	// 产生了新的一层
	if willImproveLevel >= level {
		skipList.Level++
		skipList.HeadNode.NextNodes = append(skipList.HeadNode.NextNodes, nil)
	}

	// 生成新的节点
	var newNode = &OneNode{Data: _newData, NextNodes: make([]*OneNode, willImproveLevel+1)}

	// 将新节点插入回每层
	for i := 0; i <= willImproveLevel; i++ {
		newNode.NextNodes[i] = updateNodes[i].NextNodes[i]
		updateNodes[i].NextNodes[i] = newNode
	}
}

func (skipList *SkipList) Delete(_value int) {

}

func (skipList *SkipList) Find(_value int) bool {
	var (
		tmpNode   = skipList.HeadNode
		frontNode = skipList.HeadNode
	)

	for i := skipList.Level - 1; i >= 0; i-- {
		tmpNode = frontNode.NextNodes[i]
		for tmpNode != nil {
			fmt.Println("here ", tmpNode.Data)
			if tmpNode.Data == _value {
				return true
			}
			if tmpNode.Data > _value {
				break
			}
			frontNode = tmpNode
			tmpNode = tmpNode.NextNodes[i]
		}
	}
	return false
}

func (skipList *SkipList) Display() {
	var (
		tmpNode = skipList.HeadNode
	)

	for i := skipList.Level - 1; i >= 0; i-- {
		fmt.Printf("Level is %v   :", i)
		tmpNode = skipList.HeadNode.NextNodes[i]
		for tmpNode != nil {
			fmt.Printf(" %v", tmpNode.Data)
			tmpNode = tmpNode.NextNodes[i]
		}
		fmt.Println()
	}
}

func needToImprove() bool { // 50%几率 是否提升
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	randValue := r.Intn(2)
	return randValue == 0
}

func main() {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	var skipListNode = NewSkipList()
	for i := 0; i < 10; i++ {
		skipListNode.InsertNode(r.Intn(100))
	}
	skipListNode.InsertNode(77)

	skipListNode.Display()

	fmt.Println(skipListNode.Find(77))
}
