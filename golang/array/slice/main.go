package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type SliceHeader struct {
	Data uintptr // 指向底层数组的指针
	Len  int     // 切片的长度， 大于切片的长度会报错
	Cap  int     // 底层数组的大小
}

// 切片 数组的封装
// 可动态扩容的数组

func main() {
	arr := make([]int64, 8, 10)
	for i := int64(0); i < 8; i++ { // a[0] - a[7]  = 0 - 7
		arr[i] = i
	}

	arrP := (*reflect.SliceHeader)(unsafe.Pointer(&arr))                                                       // 切片实际上是SliceHeader的形式
	fmt.Println((*int64)(unsafe.Pointer(arrP.Data)), *(*int64)(unsafe.Pointer(arrP.Data)), arrP.Len, arrP.Cap) // output:0xc4200140a0 0 8 10

	b := arr[:3]
	bP := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	fmt.Println((*int64)(unsafe.Pointer(bP.Data)), *(*int64)(unsafe.Pointer(bP.Data)), bP.Len, bP.Cap) // output:0xc4200140a0 0 3 10

	// 切片b和切片arr指向的是同一个地址空间，拥有同一个底层数组， 所以修改b也会修改arr
	b[0] = 999
	fmt.Printf("修改后的 arr[0]=%v, b[0]=%v\n", arr[0], b[0]) // output: 修改后的 arr[0]=999, b[0]=%999

	fmt.Println("----------")

	// 如果append没有超过cap时，还是会修改同一个底层数组
	c := arr[:]
	c = append(c, 1)
	cP := (*reflect.SliceHeader)(unsafe.Pointer(&c))
	fmt.Println((*int64)(unsafe.Pointer(cP.Data)), *(*int64)(unsafe.Pointer(cP.Data)), cP.Len, cP.Cap) // output:0xc4200140a0 999 9 10

	// 但是如果append超过cap时，就会动态申请一个新的内存，自动将原来的数据copy过去
	// cap自动扩容成原来的两倍
	d := arr[:]
	d = append(d, 1, 2, 3)
	dP := (*reflect.SliceHeader)(unsafe.Pointer(&d))
	fmt.Println((*int64)(unsafe.Pointer(dP.Data)), *(*int64)(unsafe.Pointer(dP.Data)), dP.Len, dP.Cap) // output:0xc42009a000 999 11 20

	// 单纯声明一个切片的话 是一个nil
	// var e []int64 =
	// eP := (*reflect.SliceHeader)(unsafe.Pointer(&e))
	// fmt.Println((*int64)(unsafe.Pointer(eP.Data)), *(*int64)(unsafe.Pointer(eP.Data)), eP.Len, eP.Cap) // panic

	// 如果想深度拷贝一个数组的话用copy
	f := make([]int64, 2, 2)
	copy(f, arr) // 将参数2的元素copy到参数1里
	fP := (*reflect.SliceHeader)(unsafe.Pointer(&f))
	fmt.Println((*int64)(unsafe.Pointer(fP.Data)), *(*int64)(unsafe.Pointer(fP.Data)), fP.Len, fP.Cap)

	fmt.Println("----------")

	g := arr[1:]
	gP := (*reflect.SliceHeader)(unsafe.Pointer(&g))
	fmt.Println((*int64)(unsafe.Pointer(gP.Data)), *(*int64)(unsafe.Pointer(gP.Data)), gP.Len, gP.Cap)
	fmt.Printf("&arr[0]=%p,&arr[1]=%p,&g=%v", &arr[0], &arr[1], (*int64)(unsafe.Pointer(gP.Data)))

	// 所以可以看出
	// 切片g:SliceHeader.Data = &arr[1]
	//		 SliceHeader.Len = arr.SliceHeader.Len-1
	//		 SliceHeader.Cap = arr.SliceHeader.Cap-1
}
