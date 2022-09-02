package main

import (
	"fmt"
	"unsafe"
)

type Empty struct {
}

//go:linkname zerobase runtime.zerobase
var zerobase uintptr // 使用go:linkname编译指令，将zerobase变量指向runtime.zerobase '//'和'go'中间没有空格

// go run -gcflags "-m -l" empty.go
// 编译器觉得变量应该分配在堆和栈上的原则是：
// 1. 变量是否被取地址；
// 2. 变量是否发生逃逸。
func main() {
	a := Empty{}
	b := struct{}{}
	fmt.Printf("sizeof(a): %v\n", unsafe.Sizeof(a))
	fmt.Printf("sizeof(b): %v\n", unsafe.Sizeof(b))

	fmt.Printf("ptr of a: %p\n", &a)
	fmt.Printf("ptr of b: %p\n", &b)
	fmt.Printf("ptr of zerobase: %p\n", &zerobase)

	c := new(Empty)
	d := new(Empty)
	fmt.Sprint(c, d) // 目的是让变量c和d发生逃逸。如果函数参数为 interface{}，编译期间很难确定其参数的具体类型，也会发生逃逸。fmt.Println() 的参数类型定义为 interface{}，因此也发生了逃逸。
	println(c)
	println(d)
	fmt.Println(c == d) // true

	e := new(Empty)
	f := new(Empty)
	println(e)          // 0xc00008ef47
	println(f)          // 0xc00008ef47
	fmt.Println(e == f) // flase
}
