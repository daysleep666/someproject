package main

import (
	"fmt"
	"math/rand"
)

type heap []int

func Heap(arr []int) {
	down(arr, len(arr)-1)
}

func up(arr []int, start, end int) {
	for start < end {

	}
}

func down(arr []int, start int) {
	for start > 0 {
		if (start-1)/2 >= 0 && arr[(start-1)/2] > arr[start] {
			arr[(start-1)/2], arr[start] = arr[start], arr[(start-1)/2]
		}
		start = (start - 1) / 2
	}
}

func IsHeap(arr []int, from int) bool {
	var a, b bool
	if from*2+1 < len(arr) {
		if arr[from*2+1] >= arr[from] {
			a = IsHeap(arr, from*2+1)
		} else {
			a = false
		}
	} else {
		a = true
	}
	if from*2+2 < len(arr) {
		if arr[from*2+2] >= arr[from] {
			b = IsHeap(arr, from*2+2)
		} else {
			b = false
		}
	} else {
		b = true
	}
	return a && b
}

func main() {
	arr := make([]int, 0)
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(10))
	}
	Heap(arr)
	fmt.Println(arr, IsHeap(arr, 0))
}
