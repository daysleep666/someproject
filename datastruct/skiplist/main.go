package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	MAXLEVEL           = 32
	LIFTINGPROBABILITY = 50 // 50%
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
		HeadNode: &OneNode{NextNodes: make([]*OneNode, MAXLEVEL)},
	}
}

// 由小到大
func (skipList *SkipList) InsertNode(_newData int) {
	var (
		frontNode   = skipList.HeadNode
		tmpNode     *OneNode
		updateNodes = make([]*OneNode, MAXLEVEL) // 记录
	)

	// 先找到插入的位置， 记录每层向下移动的点
	for i := skipList.Level; i >= 0; i-- {
		tmpNode = frontNode.NextNodes[i]
		for tmpNode != nil && tmpNode.Data <= _newData { // 在本层找到符合条件的点
			if tmpNode.Data == _newData { // 不允许重复
				return
			}
			frontNode = tmpNode
			tmpNode = tmpNode.NextNodes[i]
		}
		updateNodes[i] = frontNode
	}

	// 计算需要提升几层
	var willImproveLevel = needToImprove()

	// 产生了新层
	for willImproveLevel >= skipList.Level {
		updateNodes[skipList.Level] = skipList.HeadNode
		skipList.Level++
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
	var (
		tmpNode   = skipList.HeadNode
		frontNode *OneNode
	)
	for i := skipList.Level - 1; i >= 0; i-- {
		tmpNode = skipList.HeadNode.NextNodes[i]
		frontNode = nil
		for tmpNode != nil {
			for tmpNode.Data == _value {
				if frontNode == nil {
					skipList.Level--
					break
				}
				frontNode.NextNodes[i] = tmpNode.NextNodes[i]
				tmpNode = tmpNode.NextNodes[i]
				if tmpNode == nil {
					break
				}
			}
			if tmpNode == nil || tmpNode.Data > _value {
				break
			}
			frontNode = tmpNode
			tmpNode = tmpNode.NextNodes[i]
		}
	}
}

func (skipList *SkipList) Find(_value int) bool {
	var (
		tmpNode   = skipList.HeadNode
		frontNode = skipList.HeadNode
	)
	fmt.Printf("搜索路径: ")
	for i := skipList.Level - 1; i >= 0; i-- {
		tmpNode = frontNode.NextNodes[i]
		for tmpNode != nil {
			if tmpNode.Data == _value {
				fmt.Printf("%v\n", _value)
				return true
			}
			fmt.Printf("%v-->", tmpNode.Data)
			if tmpNode.Data > _value {
				break
			}
			frontNode = tmpNode
			tmpNode = tmpNode.NextNodes[i]
		}
	}
	fmt.Printf("none\n")
	return false
}

func (skipList *SkipList) Display() {
	var (
		tmpNode = skipList.HeadNode
	)

	for i := skipList.Level - 1; i >= 0; i-- {
		fmt.Printf("Level %v   :", i)
		tmpNode = skipList.HeadNode.NextNodes[i]
		for tmpNode != nil {
			fmt.Printf(" %v", tmpNode.Data)
			tmpNode = tmpNode.NextNodes[i]
		}
		fmt.Println()
	}
}

func needToImprove() int { //  是否提升
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	var newLevel int
	for i := 0; i < MAXLEVEL; i++ {
		randValue := r.Intn(100)
		if randValue < LIFTINGPROBABILITY {
			newLevel++
		} else {
			break
		}
	}
	return newLevel
}

func main() {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	var skipListNode = NewSkipList()
	for i := 0; i < 1000; i++ {
		skipListNode.InsertNode(r.Intn(1000))
	}
	skipListNode.Display()
	// fmt.Println("---------------")
	// skipListNode.Delete(77)
	// skipListNode.Display()
	// fmt.Println(skipListNode.Find(77))
}
