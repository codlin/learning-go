package main

import "fmt"

/*
boolean

布尔类型的值只有两种：false 和 true
布尔变量的初始默认值是 false
布尔值可以和&&和||操作符结合，并且有短路行为。
因为&&的优先级比|| 高，所以会优先结合&&，但是运算时还是从左往右
*/

func boolean() {
	c := 'd'
	if 'a' <= c && c <= 'z' ||
		'A' <= c && c <= 'Z' {
		fmt.Printf("character is %c\n", c)
	}

	// 布尔类型不能强制转为0或1
	a := true
	// int(a)
	b := 0
	if a {
		b = 1
	}
	fmt.Println(b)
}

func main() {
	boolean()
}
