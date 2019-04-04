package main

import "fmt"

type BitMap struct {
	Bits  int
	Bytes []byte
}

func NewBitMap(bits int) *BitMap {
	return &BitMap{Bits: bits, Bytes: make([]byte, bits/8+1)}
}

func (bm *BitMap) Set(num int) {
	if num >= bm.Bits {
		return
	}
	index := num / 8
	bits := uint(num % 8)
	bm.Bytes[index] |= 1 << bits
}

func (bm *BitMap) IsExist(num int) bool {
	if num >= bm.Bits {
		return false
	}
	index := num / 8
	bits := uint(num % 8)
	return bm.Bytes[index]&(1<<bits) != 0
}

func main() {
	bm := NewBitMap(10)
	bm.Set(3)
	bm.Set(5)
	bm.Set(9)
	fmt.Println(bm.IsExist(3))
	fmt.Println(bm.IsExist(5))
	fmt.Println(bm.IsExist(9))
	fmt.Println(bm.IsExist(10))
}
