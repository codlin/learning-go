package main

import (
	"fmt"
	"strings"
)

/*
函数值

在Go中，函数是被看作第一等值（first-class values）：函数像其它值一样拥有类型，也可以被赋值给其它变量，传递给函数，从函数返回。

函数的零值是nil，调用值为nil的函数会引发panic。
函数可以和nil比较：
var f func(int) int
if f == nil {
	...
}

但函数的值是不可以比较的，也不可以用函数的值作为map的key。
*/

/*
函数值不仅仅是我们可以通过数据来参数化函数，亦可通过行为。
*/
func func1(r rune) rune { return r + 1 }

func func_behaviour() {
	fmt.Println(strings.Map(func1, "hello world"))
	fmt.Println(strings.Map(func1, "foo bar"))
}

func main() {
	func_behaviour()
}
