package main

import "fmt"

/*
 # 类型声明

 一个类型声明语句创建一个新的类型名称，和现有类型具有相同的底层结构（如在内存的存储大小，内部是如何表达的，是否支持一些操作等）。
 新名命的类型提供了一种方法，用于隔离不同概念的类型，即使它们的底层类型相同也是不兼容的。

 类型声明一般出现在包一级，其创建格式为：
 type 类型名字 底层类型
 如果创建的类型名字的首字母大写，则在包外部也可以使用它。
*/

// 创建了两个类型，虽然Celsius和Fahrenheit具有相同的底层类型float64，但它们是不同的类型，因此不能相互比较或混在一起计算
// 这样可以避免使用不同温度单位混合计算导致的错误（譬如直接采用float64类型）。
type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
)

// Celsius(t)和Fahrenheit(t)是类型转换而不是函数调用，它不会改变值本身，但会使它们的语义发生变化。
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

/*
对应任意类型T，都有一个对应的类型转换操作T(x)，用于将x转换为类型T（如果T是指针类型，需要加括号，比如(* int)(p)，因为不加括号变成了*操作）。
只有当两个类型的底层类型相同时才允许这种转换操作，或者两者都是指向相同底层结构的指针类型，这些转换只改变类型而不会影响值本身。
*/
func tempconv() {
	fmt.Printf("%g\n", BoilingC-FreezingC)
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC))
	// fmt.Printf("%g\n", boilingF-FreezingC) // 编译错误，不是相同的类型
}

/* 比较运算符可以用来比较一个名命类型的变量和另一个相同类型的变量，或者有相同底层类型的未命名类型的值之间的比较。 */
func tempcmp() {
	var c Celsius
	var f Fahrenheit

	fmt.Println(c == 0) // "true"
	fmt.Println(f >= 0) // "true"
	// fmt.Println(c == f) // 编译错误，不是相同的名命类型
	fmt.Println(c == Celsius(f)) // "true"
}
