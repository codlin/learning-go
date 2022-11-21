package main

import (
	"fmt"
)

func assignment() {
	// 不能在初始化时将 nil 直接赋值给一个没有类型的变量，因为 nil 无法推断出变量的类型
	// b := nil

	// 可以这样
	var a chan int = nil
	fmt.Println(a)
}

func main() {
	assignment()
}
