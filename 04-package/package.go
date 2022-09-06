package main

import (
	"fmt"

	"mypackage/bar"
	"mypackage/foo"
)

/*
GO语言的包和其它语言的库和模块的概念类似，都是为了支持模块化、封装、单独编译和代码宠用。
一个包的源代码保存在一个或者多个以.go以文件后缀名的源文件中，通常一个包所在的目录路径的后缀是包导入的路径。
例如包gopl.io/ch1/helloworld对应的目录路径是$GOPATH/src/gopl.io/ch1/helloworld。
每个包都对应一个独立的名字空间。包可以控制哪些名字是外部可见的，简单的规则是：
如果一个名字以大写字母开头，那么该名字是可导出的，外部可以引用。
*/

/* 包的初始化首先是解决包级变量的依赖顺序，然后按照包级变量的声明出现的顺序依次初始化 */
var a = b + c // a 第三个初始化
var b = f()   // b 第二个初始化
var c = 1     // c 第一个初始化

func f() int {
	fmt.Println("f was called")
	return c + 1
}

/*
如果包中包含多个源文件，那么将按照它们发给编译器的顺序进行初始化。GO语言的构建工具会首先将文件根据文件名排序，然后依次调用编译器编译。
对于在包级别声明的变量，如果有初始化表达式则用表达式初始化，还有一些没有初始化表达式的，可以用一个特殊的init()初始化函数来简化初始化工作。
每个文件可以包含多个init()函数。
*/
var d string

func init() {
	d = "hello"
	fmt.Println("init 1 was called")
}

/* 这样的init初始化函数除了不能被调用或引用外，其行为和普通函数类似。 在每个文件中的init初始化函数，在程序开始执行时按照它们的声明顺序被自动调用。 */
func init() {
	c = 2
	fmt.Println("init 2 was called")
}

/*
每个包在依赖解决的前提下，以导入声明的顺序初始化，每个包只能被初始化一次。
关于包的导入的更多例子请参见：https://www.liwenzhou.com/posts/Go/import_local_package_in_go_module/
*/

func main() {
	fmt.Println(a)
	fmt.Printf("foo.D %s\n", foo.D)
	fmt.Printf("bar.C %s\n", bar.C)
}
