package main

import "fmt"

/*
空标识符，即_（就是下划线），用于任何语法需要变量名但程序不需要的时候。例如，在循环里，不需要下标index，或者不需要函数的某个返回值等等。
*/

func blank_id() {
	_, b := 1, 2

	c := map[string]interface{}{"foo": 1, "hello": 2}
	d := c["foo"]
	e := d.(interface{})
	_, ok := c["bar"]

	f := []int{10, 20, 30}
	for _, v := range f {
		fmt.Println(v)
	}

	fmt.Println(b, c, d, e, ok)
}

func main() {
	blank_id()
}
