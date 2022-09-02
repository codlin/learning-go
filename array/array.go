package main

import (
	"fmt"
	"strings"
)

func arry_len_4() {
	// 声明一个长度为4的数组
	a := [...]int64{1, 2, 3, 4}
	b := a
	c := &a

	fmt.Printf("a: %v\n", a) // [1 2 3 4]
	fmt.Printf("b: %v\n", b) // [1 2 3 4]

	b[1] = 100
	fmt.Printf("b修改后: %v\n", b)      // [1 100 3 4]
	fmt.Printf("b修改后， a的值: %v\n", a) // [1 2 3 4]
	fmt.Printf("a 的内存地址: %p\n", &a)
	fmt.Printf("b 的内存地址: %p\n", &b)

	/**
	 * 由此可见，将一个数组赋值给另一个变量，是值copy，即会生成一个新的数组
	 * 并且数组 b 的内容在编译期已经确定。
	**/
	fmt.Println(`由此可见，将一个数组赋值给另一个变量，是值copy，即会生成一个新的数组。
且数组 b 的内容在编译期已经确定，从内存地址看两个数组相邻地址差32个字节，都在栈上。`)

	fmt.Println(strings.Repeat("-", 32))
	c[2] = 200
	fmt.Printf("c: %v\n", *c)         // [1 2 200 4]
	fmt.Printf("c 修改后 a 的值: %v\n", a) // [1 2 200 4]
	fmt.Printf("a 的内存地址: %p\n", &a)
	fmt.Printf("c 的内存地址: %p，即 &c，表示c 在栈上的地址\n", &c)
	fmt.Printf("c 的相关值: %p， 以%%p打印 c 会打印出它的值：数组 a 的首地址\n", c)
	fmt.Printf("c 的相关值: %v，以%%v打印 c 会打印出它引用的值：数组 a 的内容并用 & 代表引用，不会显式显示地址值\n", c)
	fmt.Printf("c 的相关值: %v，以%%v打印 *c 会打出 c 指向的内容，即 a 的内容\n", *c)
	/**
	 * 从打印的 &a 和 &c 的值看，两者差别较大，按照正常来说应该差别
	**/
	fmt.Println(strings.Repeat("=", 32))
}

func arry_len_big_4() {
	// 声明一个长度为4的数组
	a := [...]int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	b := a
	c := &a

	fmt.Printf("a: %v\n", a) // [1 2 3 4]
	fmt.Printf("b: %v\n", b) // [1 2 3 4]

	b[1] = 100
	fmt.Printf("b修改后: %v\n", b)      // [1 100 3 4]
	fmt.Printf("b修改后， a的值: %v\n", a) // [1 2 3 4]
	fmt.Printf("a 的内存地址: %p\n", &a)
	fmt.Printf("b 的内存地址: %p\n", &b)

	/**
	 * 由此可见，将一个数组赋值给另一个变量，是值copy，即会生成一个新的数组
	 * 并且数组 b 的内容在编译期已经确定。
	**/
	fmt.Println(`由此可见，将一个数组赋值给另一个变量，是值copy，即会生成一个新的数组。
且数组 b 的内容在编译期已经确定，从内存地址看两个数组相邻地址差32个字节，都在栈上。`)

	fmt.Println(strings.Repeat("-", 32))
	c[2] = 200
	fmt.Printf("c: %v\n", *c)         // [1 2 200 4]
	fmt.Printf("c 修改后 a 的值: %v\n", a) // [1 2 200 4]
	fmt.Printf("a 的内存地址: %p\n", &a)
	fmt.Printf("c 的内存地址: %p，即 &c，表示c 在栈上的地址\n", &c)
	fmt.Printf("c 的相关值: %p， 以%%p打印 c 会打印出它的值：数组 a 的首地址\n", c)
	fmt.Printf("c 的相关值: %v，以%%v打印 c 会打印出它引用的值：数组 a 的内容并用 & 代表引用，不会显式显示地址值\n", c)
	fmt.Printf("c 的相关值: %v，以%%v打印 *c 会打出 c 指向的内容，即 a 的内容\n", *c)
	/**
	 * 从打印的 &a 和 &c 的值看，两者差别较大，按照正常来说应该差别
	**/
	fmt.Println(strings.Repeat("=", 32))
}

func main() {
	arry_len_4()

	arry_len_big_4()
}
