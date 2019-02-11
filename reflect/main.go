package main

import (
	"fmt"
	"reflect"
)

type Test struct {
	I int
	F float32
	S string
}

func (test Test) Write(_val string) {
	fmt.Println("Write.......", _val)
}

func (test Test) Read() {
	fmt.Println("Read.......")
}

func main() {
	// var test Test = Test{I: 1, F: 2.0, S: "test"}
	// typ := reflect.TypeOf(test)
	// val := reflect.ValueOf(test)
	// fmt.Println(typ)

	// for i := 0; i < typ.NumField(); i++ {
	// 	fmt.Println(typ.Field(i).Name, val.Field(i))
	// }
	// fmt.Println("元素数量：", val.NumField())
	// for i := 0; i < typ.NumMethod(); i++ {
	// 	m := typ.Method(i)
	// 	fmt.Println(m.Name, m.Type)
	// }

	var test Test = Test{I: 1, F: 2.0, S: "test"}
	val := reflect.ValueOf(test)
	readFunc := val.MethodByName("Read")
	args := []reflect.Value{}
	readFunc.Call(args)

	writeFunc := val.MethodByName("Write")
	args = []reflect.Value{reflect.ValueOf("hello")}
	writeFunc.Call(args)
}
