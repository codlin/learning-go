/*
一个GO语言编写的程序对应一个或多个以.go为文件后缀名的源文件。
每个源文件以包的声明语句开始，说明该文件属于哪个包。
包声明语句之后是import语句导入依赖的其它包，
然后是包一级的类型、变量、常量、函数的声明语句。包一级的各种`类型的声明语句`的顺序无关紧要。
*/
package main

import "fmt"

/*
GO语言主要有四种`类型的声明语句`：
1. var       变量声明
2. const     常量声明
3. type      类型实体对象声明
4. func      函数实体对象声明
*/

/* 常量 boilingF 是在包一级范围声明的，在包一级范围声明的名字可以在整个包对应的每个源文件中访问，而不是只在其所在的源文件访问。 */
const boilingF = 212.0

/* 一个函数的声明由函数名字，参数列表，可选的返回值以及包含函数定义的函数体构成 */
func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

func main() {
	var f = boilingF
	var c = fToC(f)
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
}
