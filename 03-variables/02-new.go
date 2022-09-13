package main

import "fmt"

/*
另一个创建变量的方法是用内建的函数new，表达式new(T)创建一个T类型的匿名变量，初始化为类型T的零值，返回变量的地址，返回的指针类型为*T
*/
func new_var() {
	p := new(int)
	fmt.Println(*p)

	a := new([]int)
	fmt.Println(*a)
}

func main() {
	new_var()
}
