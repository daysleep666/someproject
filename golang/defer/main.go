package main

import "fmt"

func defer1() { // defer 是先进后出的
	defer fmt.Println("first")
	defer fmt.Println("two")
	defer fmt.Println("three")

	// output:
	// three
	// two
	// first
}

// 在defer执行时会将函数及其参数压入函数栈但是并不会被调用。
// 函数参数被压入到函数栈后，其值就固定了。
func defer2() { //
	var i = 0
	defer fmt.Printf("%v", i) // output:0
	i++
	fmt.Printf("cur1:i=%p.", &i)
}

func defer3() {
	var i = 0
	defer func() {
		fmt.Printf("%v", i) // output:1
	}()
	i++
}

func defer4() {
	var i = 0
	defer func(i int) {
		fmt.Printf("%v", i) // output:0
	}(i)
	i++
}

func defer5() {
	var i = 0
	defer c(b(a(i))) //  output:2
	i++
}

func a(i int) int {
	i++
	return i
}

func b(i int) int {
	i++
	return i
}

func c(i int) {
	fmt.Printf("%v", i)
}

// return不是原子操作
// return v分为
// 1. 返回值 = xxx
// 2. 空的return
// defer语句插在1.2之间
// *******
// 返回值 = xxx
// 调用defer函数
// 空的return
// *******
// 所以在defer中可能对返回值有修改，并影响到最终的结果

// 实际 1
func defer6() (result int) {
	defer func() {
		result++
	}()
	return 0
}

// 实际 5
func defer7() (result int) {
	t := 5
	defer func() {
		t += 5
	}()
	return t
}

// 实际 6
func defer8() (result int) {
	defer func(r int) {
		result += 5
	}(result)
	return 1
}

// return -> defer -> 返回

func main() {
	// defer1()
	// defer2()
	// defer3()
	// defer4()
	// defer5()
	fmt.Println(defer6())
	fmt.Println(defer7())
	fmt.Println(defer8())
}
