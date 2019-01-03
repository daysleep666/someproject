package main

import (
	"fmt"
)

// 层级时间轮

type SlotNode struct {
	Key     interface{}
	Task    func()
	RunTime int64
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

func (sn *SlotNode) AddTask(_key interface{}, _newTask func(), _runTime int64) {
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

const MaxSlot = 6 // 总槽数
type TimeWheel struct {
	RunTime   int         // 执行时间
	CurSlot   int         // 当前槽
	SlotArray []*SlotNode // 槽位
}

func NewTimeWheel() *TimeWheel {
	tw := &TimeWheel{SlotArray: make([]*SlotNode, MaxSlot)}
	for i, _ := range tw.SlotArray {
		tw.SlotArray[i] = NewSlotNode()
	}
	return tw
}

type LevelTimeWheel struct {
	TimeWheelArray []*TimeWheel // 所有的时间轮
}

func InitLevelTimeWheel() *LevelTimeWheel {
	return &LevelTimeWheel{[]*TimeWheel{}}
}

var count int64

func (ltw *LevelTimeWheel) Tick() {
	for {
		if count >= 40 {
			return
		}

		var hasRoundOne bool // 是否已经绕完了一圈
		for level, tw := range ltw.TimeWheelArray {
			if level == 0 { // 第一层的节点
				for slot, slotNode := range tw.SlotArray {
					fmt.Printf("%v:", count)
					if slotNode != nil {
						slotNode.Run()
						tw.SlotArray[slot] = NewSlotNode()
					}
					tw.CurSlot++

					if tw.CurSlot == MaxSlot { //绕完一圈了
						tw.CurSlot = 0
						hasRoundOne = true
					}

					{ // debug
						fmt.Println()
						count++
					}
				}
			} else {
				if hasRoundOne {
					tw.CurSlot++

					// 判断是否绕完一圈了
					if tw.CurSlot == MaxSlot { //绕完一圈了
						tw.CurSlot = 0
						hasRoundOne = true
					} else {
						hasRoundOne = false
					}
					// 看有没有任务，如果有,执行任务降级
					tmpNode := tw.SlotArray[tw.CurSlot].Next
					for tmpNode != nil {
						ltw.AddTask(tmpNode.Key, tmpNode.Task, computeFallRunTime(int64(level)-1, tmpNode.RunTime))
						// fmt.Printf("key=%v,runtime=%v,fallruntime=%v\n", tmpNode.Key, tmpNode.RunTime, computeFallRunTime(int64(level)-1, tmpNode.RunTime))
						tmpNode = tmpNode.Next
					}
					tw.SlotArray[tw.CurSlot] = NewSlotNode()
				}
			}
		}
	}
}

func (ltw *LevelTimeWheel) AddTask(_key interface{}, _newTask func(), _runTime int64) {
	if _runTime == 0 {
		return
	}

	twLevel := computeLevel(_runTime)

	for int64(len(ltw.TimeWheelArray)) <= twLevel {
		ltw.TimeWheelArray = append(ltw.TimeWheelArray, NewTimeWheel())
	}

	twSlot := computeSlot(twLevel, _runTime)
	ltw.TimeWheelArray[twLevel].SlotArray[twSlot].AddTask(_key, _newTask, _runTime)
}

// 假设一个时间轮有六个槽，一个槽1秒。
// 向时间轮里插入行为，当数据大于6秒时，仿照第一个时间轮动态建立第二个时间轮，
// 第二个时间轮有6*6秒，如果执行时间v在36秒内，则可以插入到(v%6)个槽中。
// 第一个时间轮每绕一圈，第二个时间轮走一下，当第二个时间轮走到有行为的槽时，
// 降级到第一个时间轮的第(v/6)槽中，第一个时间轮走到该槽后执行该行为。

func main() {
	ltw := InitLevelTimeWheel()
	// for i:=0;i<100;i++ {

	// }
	ltw.AddTask(1, func() { fmt.Printf("task1") }, 2)
	ltw.AddTask(2, func() { fmt.Printf("task2") }, 10)
	ltw.AddTask(3, func() { fmt.Printf("task3") }, 37)
	ltw.AddTask(4, func() { fmt.Printf("task4") }, 40)
	ltw.Tick()
}

func computeLevel(_val int64) int64 {
	var (
		count int64
		num   int64 = MaxSlot
	)
	for num < _val {
		count++
		num *= MaxSlot
	}
	return count
}

func computeSlot(_level, _val int64) int64 {
	if _level == 0 {
		return _val % MaxSlot
	}
	var tmp = int64(1)
	for i := int64(0); i < _level; i++ {
		tmp *= 6
	}
	tmp = _val - tmp

	var slot = int64(1)
	for tmp > MaxSlot {
		slot++
		tmp /= MaxSlot
	}
	return slot
}

func computeFallRunTime(_fallLevel, _runTime int64) int64 {
	var tmp = int64(1)
	for i := int64(0); i <= _fallLevel; i++ {
		tmp *= 6
	}
	tmp = _runTime - tmp
	return tmp
}
