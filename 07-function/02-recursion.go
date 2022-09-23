package main

import "fmt"

/*
递归

函数可以直接或间接调用自身，称之为递归。

递归需要注意的事项：
大部分编程语言都是使用固定大小的函数调用栈，常见的大小从64KB到2MB不等。
固定大小栈会限制递归的深度，如果递归的深度超过了一定的值，会导致栈溢出。除此之外，还会导致安全性问题。

Go语言使用可变栈，栈的大小按需增补。这使得我们在使用递归时不需要考虑栈溢出和安全问题。
*/

func factorial(i int) int {
	if i < 2 {
		return 1
	}
	return i * factorial(i-1)
}

func main() {
	fmt.Println(factorial(4))
}
