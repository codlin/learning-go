package scope

import (
	"fmt"
	"log"
	"os"
)

/*
一个声明语句将程序中的实体和一个名字关联，比如一个函数和一个变量。
声明语句的作用域是指源代码中可以有效使用这个名字的范围。
不要将作用域和生命周期混为一谈。声明语句的作用域对应的是一个源代码的文本区域；它是一个编译时属性。
一个变量的生命周期是指程序运行时变量存在的有效时段，在此时段内它可以被程序的其它部分引用；是一个运行时概念。

语法块是由花括号包含的一系列语句，就像函数体或循环体花括号对应的语法块那样。语法块内声明的名字无法被外部语法块引用。
语法块定了内部声明的名字的作用域范围：
有一个语法块为整个源码，称为全局语法块；
然后是每个包的包语法块；
每个函数的语法块；
每个for，if和switch的语法块；
每个switch和select的分支也有单独的语法块；
显式用花括号包括的语法块；

声明语句对应的词法域决定了作用域范围的大小：
对于内置的类型、函数和常量，如int、len和true等是在全局作用域的，因此可以在整个程序中直接使用。
任何在函数外部声明的名字可以在同一个包的任何源文件中访问（包级语法域）。
对于导入的包，则是源文件级的作用域，其它文件无法访问当前源文件导入的包。
还有很多声明语句，只能在函数内部或内部的语法块中使用。

当编译器遇到一个名字引用时，它会查找声明，首先从内层的词法域向全局作用域查找。如果查找失败，则报告“未声明的名字”错误。
如果该名字在内部和外部的块分别声明过，则内部的声明首先被找到。在这种情况下，内部的声明屏蔽了外部的同名声明，让外部的声明的名字无法被访问。
*/
func f() int { return 1 }

var g = "g"

func shadow() {
	f := "f"
	fmt.Println(f) // "f"; local var f shadows package-level func f
	fmt.Println(g) // "g"; package-level var
	// fmt.Println(h)
}

func shadow2() {
	x := "hello"
	for _, x := range x { // 前面的 x 遮蔽了上面声明的 x；前面x的作用域范围是整个for循环语句
		x := x + 'A' - 'a'  // 前面的 x 遮蔽了 for 循环声明的 x
		fmt.Printf("%c", x) // "HELLO"
	}
}

func shadow3() {
	if x := f(); x == 0 { // x 作用域为整个if语句
		fmt.Println(x)
	} else if y := 2; x == y { // 可以访问x，因为嵌套在else语句里
		fmt.Println(x, y)
	} else {
		fmt.Println(x, y) // y 可以访问，因为它处在if语句的作用域范围
	}

	// fmt.Println(x, y) // 编译错误，x和y出了作用域
}

/* 在包级别，变量声明的顺序并不影响其作用域范围，先声明的可以引用后面的一个声明。但是循环递归引用会产生编译错误 */
var first = last
var last = 2

// var self = recur // 编译错误；递归引用
// var recur = self // 编译错误；递归引用

/* 对于短变量声明语句的作用范围，要特别注意，因为可能出现作用域遮蔽 */
var cwd string

func init() {
	cwd, err := os.Getwd() // package-level cwd 被遮蔽
	if err != nil {
		log.Fatalf("Getwd failed: %v", err)
	}
	fmt.Println(cwd) // 打印的是局部cwd
}

func CWD() string {
	return cwd // 返回的为空
}
