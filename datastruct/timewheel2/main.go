package main

// 层级时间轮

type OneNode struct {
	Task func()
	Next *OneNode
}

func AddNode(_oneNode *OneNode, _newTask func()) *OneNode { // 传的是指针的拷贝
	if _oneNode == nil {
		_oneNode = new(OneNode)
		_oneNode.Task = _newTask
		return _oneNode
	}

	_oneNode.Next = AddNode(_oneNode.Next, _newTask)
	return _oneNode
}

type TimeWheel struct {
	CurSlot   int        // 当前槽
	MaxSlot   int        // 总槽数
	Level     int        // 属于第几层
	SlotArray []*OneNode // 槽位
}

type LevelTimeWheel struct {
	TimeWheelArray []*TimeWheel // 所有的时间轮
}

func (ltw *LevelTimeWheel) Tick() {
	for _, v := range ltw.TimeWheelArray {
		if v.Level == 1 {
			v.CurSlot++

		}
	}
}

func main() {

}
