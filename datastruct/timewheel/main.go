package main

import (
	"fmt"
	"time"
)

// time wheel

type timeWheel struct {
	MaxSlot   int
	Nodes     []*OneNode
	CurSlot   int
	Round     int
	TaskCount int
}

// 双向链表
type OneNode struct {
	Id       int
	Task     func()
	Round    int
	Previous *OneNode
	Next     *OneNode
}

func AddNode(_oneNode *OneNode, _id int, _newTask func(), _round int) *OneNode {
	if _oneNode == nil {
		return &OneNode{
			Id:       _id,
			Task:     _newTask,
			Round:    _round,
			Previous: nil,
			Next:     nil,
		}
	}
	var tmpNode = _oneNode
	if tmpNode.Round > _round {
		newNode := &OneNode{
			Id:       _id,
			Task:     _newTask,
			Round:    _round,
			Previous: nil,
			Next:     nil,
		}
		newNode.Next = tmpNode
		newNode.Previous = tmpNode.Previous
		tmpNode.Previous = newNode
		return newNode
	}

	var nextNode = AddNode(tmpNode.Next, _id, _newTask, _round)
	tmpNode.Next = nextNode
	nextNode.Previous = tmpNode
	return tmpNode
}

func DeleteNode(_oneNode *OneNode, _id int) *OneNode {
	if _oneNode == nil {
		return nil
	}
	var tmpNode = _oneNode
	for tmpNode.Id == _id {
		tmpNode = tmpNode.Next
		if tmpNode == nil {
			return nil
		}
	}
	var nextNode = DeleteNode(tmpNode.Next, _id)
	tmpNode.Next = nextNode
	if nextNode != nil {
		nextNode.Previous = tmpNode
	}
	return tmpNode
}

func loopExcution(_curRound int, _oneNode *OneNode) *OneNode {
	if _oneNode == nil {
		return nil
	}
	if _oneNode.Round > _curRound {
		return _oneNode
	}
	if _oneNode.Next != nil {
		_oneNode.Next.Previous = _oneNode
	}
	if _oneNode.Round == _curRound && _oneNode.Task != nil {

		_oneNode.Task() // 执行了的任务 干掉它
		return _oneNode.Next
	}
	_oneNode.Next = loopExcution(_curRound, _oneNode.Next)
	return _oneNode
}

func DisplayNode(_oneNode *OneNode) {
	if _oneNode == nil {
		return
	}
	fmt.Printf("cur=%v", _oneNode.Task)
	if _oneNode.Previous != nil {
		fmt.Printf(", previous=%v", _oneNode.Previous.Task)
	}
	if _oneNode.Next != nil {
		fmt.Printf(", next=%v", _oneNode.Next.Task)
	}
	fmt.Println()
	DisplayNode(_oneNode.Next)
}

func NewTimeWheel(_maxSlot int) *timeWheel {
	timeWheel := &timeWheel{MaxSlot: _maxSlot}
	timeWheel.Nodes = make([]*OneNode, _maxSlot)
	for i, _ := range timeWheel.Nodes {
		timeWheel.Nodes[i] = new(OneNode)
	}
	return timeWheel
}

// 延时afterTime秒后执行_newTask
func (tw *timeWheel) AddTask(_newTask func(), afterTime int) {
	tw.TaskCount++
	willInSlot := (afterTime + tw.CurSlot) % tw.MaxSlot
	round := tw.Round + (afterTime+tw.CurSlot)/tw.MaxSlot
	tw.Nodes[willInSlot] = AddNode(tw.Nodes[willInSlot], tw.TaskCount, _newTask, round)
}

var timeCount int

func (tw *timeWheel) setTick() {
	fmt.Println(timeCount)
	timeCount++
	tw.CurSlot++
	if tw.CurSlot == tw.MaxSlot {
		tw.CurSlot = 0
		tw.Round++
	}
	node := tw.Nodes[tw.CurSlot]
	loopExcution(tw.Round, node)
}

func (tw *timeWheel) Run() {
	startTime := time.Now().Unix()
	for {
		if startTime < time.Now().Unix() {
			tw.setTick()
			startTime++
		}
	}
}

func main() {
	tw := NewTimeWheel(6)
	tw.AddTask(func() { fmt.Println("task1") }, 1)
	tw.AddTask(func() { fmt.Println("task2") }, 7)
	tw.AddTask(func() { fmt.Println("task3") }, 13)
	tw.AddTask(func() { fmt.Println("task4") }, 20)

	go func() {
		time.Sleep(time.Second * 5)
		tw.AddTask(func() { fmt.Println("task5") }, 10)
	}()

	tw.Run()

	for _, v := range tw.Nodes {
		if v != nil {
			DisplayNode(v)
		}
	}

	var ch = make(chan int)
	<-ch
}
