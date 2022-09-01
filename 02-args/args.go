package main

import (
	"fmt"
	"os"
)

/*
程序的命令行参数可以从包的`os.Args`变量获取；os包外部使用`os.Args`访问该变量。
`os.Args`变量是字符串的一个切片（slice）：
`os.Args[0]`是命令本身的名字；其它的元素则是启动时程序传给它的参数。
`os.Args[m:n]`产生从第m个元素到第n-1个元素的切片。
*/

func main() {
	var s, step string

	// for 循环的方式1
	for i := 0; i < len(os.Args); i++ {
		fmt.Printf("arg %d: %s\n", i, os.Args[i])
		s += step + os.Args[i]
		step = " "
	}
	fmt.Println(s)

	// for 循环的方式2
	s = ""
	step = ""
	for i, arg := range os.Args[:] {
		fmt.Printf("arg %d: %s\n", i, arg)
		s += step + arg
		step = " "
	}
	fmt.Println(s)
}
