package heap

import "fmt"

type Heap struct { // 大顶堆
	Arr       []int64
	CurLength int64
	MaxLength int64
}

func NewHeap(_length int64) *Heap {
	length := _length
	return &Heap{
		Arr:       make([]int64, length, length),
		MaxLength: length,
	}
}

func (h *Heap) Push(_data int64) {
	if h.CurLength == h.MaxLength {
		return
	}
	h.CurLength++

	curIndex := h.CurLength
	h.Arr[curIndex] = _data

	for (curIndex-1)/2 != 0 && h.Arr[curIndex/2] < h.Arr[curIndex] {
		// 判断当前节点和父节点谁大, 交换
		h.Arr[curIndex/2], h.Arr[curIndex] = h.Arr[curIndex], h.Arr[curIndex/2]
		curIndex = curIndex / 2
	}
}

func (h *Heap) Pop() {
	if h.CurLength == 0 {
		return
	}
	// 抹去堆顶元素
	h.Arr[1] = h.Arr[h.CurLength]
	h.Arr[h.CurLength] = 0

	var curIndex int64 = 1
	for curIndex < h.MaxLength && h.Arr[curIndex] < h.Arr[curIndex*2] {
		h.Arr[curIndex], h.Arr[curIndex*2] = h.Arr[curIndex*2], h.Arr[curIndex]
	}
}

func (h *Heap) Display() {
	for i, v := range h.Arr {
		if v == 0 {
			continue
		}
		fmt.Printf("[%v]%v ", i, v)
	}
	fmt.Println()
}

func main() {
	h := NewHeap(100)
	for i := int64(1); i < 10; i++ {
		h.Push(i)
	}

	h.Display()

	h.Pop()
	h.Display()
}
