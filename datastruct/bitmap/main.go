package main

import "fmt"

func main() {
	bm := bitmap.NewBitMap(10)
	bm.Set(3)
	bm.Set(5)
	bm.Set(9)
	fmt.Println(bm.IsExist(3))
	fmt.Println(bm.IsExist(5))
	fmt.Println(bm.IsExist(9))
	fmt.Println(bm.IsExist(10))
}
