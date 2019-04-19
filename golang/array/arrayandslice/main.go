package main

import (
	"fmt"
)

func ModifyArray(_arr [3]int64) {
	_arr[0] = 10
	fmt.Printf("ModifyArray的参数地址:%p,首元素指向的地址:%p\n", &_arr, &_arr[0])
}

func ModifySlice(_arr []int64) {
	_arr[0] = 10
	fmt.Printf("ModifySlice的参数地址:%p,首元素指向的地址:%p\n", &_arr, &_arr[0])
}

func main() {
	arr := [3]int64{1, 2, 3}
	fmt.Printf("修改前:%v,  地址:%p,首元素指向的地址:%p \n", arr, &arr, &arr[0])
	ModifyArray(arr)
	fmt.Printf("修改后:%v\n", arr)

	fmt.Println("-------------------------")

	slice := []int64{1, 2, 3}
	fmt.Printf("修改前:%v,  地址:%p,首元素指向的地址:%p  \n", slice, &slice, &slice[0])
	ModifySlice(slice)
	fmt.Printf("修改后:%v\n", slice)
}
