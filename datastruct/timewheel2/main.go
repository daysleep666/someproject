package main

import (
	"fmt"
)

// 层级时间轮

// 一个槽节点
type SlotNode struct {
	Key     interface{}
	Task    func()
	RunTime int
	Next    *SlotNode
}

func NewSlotNode() *SlotNode {
	return &SlotNode{Task: func() {}}
}

func (sn *SlotNode) Run() {
	sn.Task()
	if sn.Next != nil {
		sn.Next.Run()
	}
}

func (sn *SlotNode) AddTask(_key interface{}, _newTask func(), _runTime int) {
	var tmpNode = sn
	for tmpNode.Next != nil {
		if tmpNode.Key == _key {
			return
		}
		tmpNode = tmpNode.Next
	}

	tmpNode.Next = &SlotNode{Key: _key, Task: _newTask, RunTime: _runTime}
}

func (sn *SlotNode) RemoveTask(_key interface{}) {
	var (
		tmpNode   = sn
		frontNode = tmpNode
	)
	for tmpNode.Next != nil {
		if tmpNode.Key == _key {
			frontNode.Next = tmpNode.Next
		}
		frontNode = tmpNode
		tmpNode = tmpNode.Next
	}
}

// 一个时间轮
type TimeWheel struct {
	CurSlot   int         // 当前槽
	MaxSlot   int         // 总槽数
	OneStep   int         // 步长
	SlotArray []*SlotNode // 槽位
}

func NewTimeWheel(_maxSlot int, _oneStep int) *TimeWheel {
	tw := &TimeWheel{SlotArray: make([]*SlotNode, _maxSlot), MaxSlot: _maxSlot, OneStep: _oneStep}
	for i, _ := range tw.SlotArray {
		tw.SlotArray[i] = NewSlotNode()
	}
	return tw
}

func (tw *TimeWheel) HasRound() bool {
	return tw.CurSlot == tw.MaxSlot
}

// 一堆时间轮
type LevelTimeWheel struct {
	MaxSlot        int          // 总槽数
	TotalTick      int          // 总tick
	TimeWheelArray []*TimeWheel // 所有的时间轮
}

func InitLevelTimeWheel(_maxSlot int) *LevelTimeWheel {
	return &LevelTimeWheel{TimeWheelArray: []*TimeWheel{}, MaxSlot: _maxSlot}
}

func (ltw *LevelTimeWheel) Tick() {
	for {
		if ltw.TotalTick >= 15 {
			return
		}

		for level, tw := range ltw.TimeWheelArray {
			if level == 0 { // 第一层的节点
				for slot, slotNode := range tw.SlotArray {
					tw.CurSlot++
					fmt.Printf("%v:", ltw.TotalTick)
					if slotNode != nil {
						slotNode.Run()
						tw.SlotArray[slot] = NewSlotNode()
					}

					if tw.HasRound() { //绕完一圈了
						tw.CurSlot = 0
					}

					ltw.TotalTick++

					{ // debug
						fmt.Println()
					}
				}
			} else {
				// 判断是否该走一个tick
				if ltw.TotalTick%tw.OneStep == 0 {
					tw.CurSlot++

					// 判断是否绕完一圈了
					if tw.HasRound() {
						tw.CurSlot = 0
					}

					// 看有没有任务，如果有,执行任务降级
					tmpNode := tw.SlotArray[tw.CurSlot].Next
					for tmpNode != nil {
						// 降级后的时间等于当前执行时间-已经过了的时间
						ltw.AddTask(tmpNode.Key, tmpNode.Task, tmpNode.RunTime-tw.CurSlot*tw.OneStep)
						// fmt.Printf("key=%v,runtime=%v,fallruntime=%v\n", tmpNode.Key, tmpNode.RunTime, computeFallRunTime(int(level)-1, tmpNode.RunTime))
						tmpNode = tmpNode.Next
					}
					tw.SlotArray[tw.CurSlot] = NewSlotNode()
				}
			}
		}
	}
}

func (ltw *LevelTimeWheel) AddTask(_key interface{}, _newTask func(), _runTime int) {
	// if _runTime == 0 {
	// 	return
	// }
	twLevel := ltw.computeLevel(_runTime)

	var lastStep int = 1
	for i := len(ltw.TimeWheelArray); i <= twLevel; i++ {
		if i != 0 {
			lastStep = ltw.TimeWheelArray[i-1].OneStep * ltw.MaxSlot
		}
		newTW := NewTimeWheel(ltw.MaxSlot, lastStep)
		ltw.TimeWheelArray = append(ltw.TimeWheelArray, newTW)
	}

	twSlot := (_runTime / lastStep) % ltw.MaxSlot
	fmt.Printf("twLevel=%v,twslot=%v\n", twLevel, twSlot)
	ltw.TimeWheelArray[twLevel].SlotArray[twSlot].AddTask(_key, _newTask, _runTime)
}

func (ltw *LevelTimeWheel) computeLevel(_val int) int {
	var (
		count int
		num   int = ltw.MaxSlot
	)
	for num <= _val {
		count++
		num *= ltw.MaxSlot
	}
	return count
}

func main() {
	ltw := InitLevelTimeWheel(6)
	for i := 0; i < 9; i++ {
		ltw.AddTask(i, func() { fmt.Printf("task") }, i)
	}
	// ltw.AddTask(2, func() { fmt.Printf("task2") }, 1)
	// ltw.AddTask(3, func() { fmt.Printf("task3") }, 37)
	// ltw.AddTask(4, func() { fmt.Printf("task4") }, 40)
	ltw.Tick()
}
