package main

import (
	"fmt"
	"time"
)

// time wheel

type timeWheel struct {
	MaxSlot   int
	Nodes     []*SlotNode
	CurSlot   int
	Round     int
	TaskCount int
}

// 双向链表
type SlotNode struct {
	Id       int
	Task     func()
	Round    int
	Previous *SlotNode
	Next     *SlotNode
}

func AddNode(_slotNode *SlotNode, _id int, _newTask func(), _round int) *SlotNode {
	if _slotNode == nil {
		return &SlotNode{
			Id:       _id,
			Task:     _newTask,
			Round:    _round,
			Previous: nil,
			Next:     nil,
		}
	}
	var tmpNode = _slotNode
	if tmpNode.Round > _round {
		newNode := &SlotNode{
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

func DeleteNode(_slotNode *SlotNode, _id int) *SlotNode {
	if _slotNode == nil {
		return nil
	}
	var tmpNode = _slotNode
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

func (tw *timeWheel) loopExcution(_curRound int, _slotNode *SlotNode) *SlotNode {
	if _slotNode == nil {
		return nil
	}
	if _slotNode.Round > _curRound {
		return _slotNode
	}
	if _slotNode.Next != nil {
		_slotNode.Next.Previous = _slotNode
	}
	if _slotNode.Round == _curRound && _slotNode.Task != nil {
		tw.TaskCount--
		_slotNode.Task() // 执行了的任务 干掉它
		return _slotNode.Next
	}
	_slotNode.Next = tw.loopExcution(_curRound, _slotNode.Next)
	return _slotNode
}

func DisplayNode(_slotNode *SlotNode) {
	if _slotNode == nil {
		return
	}
	fmt.Printf("cur=%v", _slotNode.Task)
	if _slotNode.Previous != nil {
		fmt.Printf(", previous=%v", _slotNode.Previous.Task)
	}
	if _slotNode.Next != nil {
		fmt.Printf(", next=%v", _slotNode.Next.Task)
	}
	fmt.Println()
	DisplayNode(_slotNode.Next)
}

func NewTimeWheel(_maxSlot int) *timeWheel {
	timeWheel := &timeWheel{MaxSlot: _maxSlot}
	timeWheel.Nodes = make([]*SlotNode, _maxSlot)
	for i, _ := range timeWheel.Nodes {
		timeWheel.Nodes[i] = new(SlotNode)
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

func (tw *timeWheel) setTick() {
	tw.CurSlot++
	if tw.CurSlot == tw.MaxSlot {
		tw.CurSlot = 0
		tw.Round++
	}
	node := tw.Nodes[tw.CurSlot]
	tw.loopExcution(tw.Round, node)
}

func (tw *timeWheel) Run() {
	startTime := time.Now().UnixNano()
	for {
		if startTime < time.Now().UnixNano() {
			tw.setTick()
			startTime++
		}
	}
}

func main() {
	tw := NewTimeWheel(6)
	tw.AddTask(func() { fmt.Printf("task1,curtaskcount=%v\n", tw.TaskCount) }, 1)
	tw.AddTask(func() { fmt.Printf("task2,curtaskcount=%v\n", tw.TaskCount) }, 7)
	tw.AddTask(func() { fmt.Printf("task3,curtaskcount=%v\n", tw.TaskCount) }, 13)
	tw.AddTask(func() { fmt.Printf("task4,curtaskcount=%v\n", tw.TaskCount) }, 20)

	tw.Run()

	for _, v := range tw.Nodes {
		if v != nil {
			DisplayNode(v)
		}
	}

	var ch = make(chan int)
	<-ch
}
