package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

/*
reflect.Type和reflect.Value

反射功能由reflect包提供，它定义了两个重要的类型：Type和Value。
Type表示Go语言的一个类型，它是一个有很多方法的接口，这些方法可以用来识别类型以及透视类型的组成部分，
比如一个结构的组成部分或者一个函数的各个参数。
reflect.Type接口只要一个实现，即类型描述符。接口值中的动态类型也是类型描述符。
reflect.TypeOf函数接受任何类型的interface{}参数，并且把接口中的动态类型以reflect.Type的形式返回。
reflect.ValueOf函数接受任何类型的interface{}参数，并且把接口中的动态值以reflect.Value的形式返回。
*/

func reflect_demo() {
	/* 把一个具体值赋给一个接口类型时会发生一个隐式类型转换，转换会生成一个包含两部分内容的接口值：
	动态类型和动态值
	此处，动态类型是int，动态值是3
	*/
	t := reflect.TypeOf(3)  // 一个reflect.Type
	fmt.Println(t.String()) // "int"
	fmt.Println(t)          // "int"
}

/*
因为reflect.TypeOf返回一个接口的动态类型，所以它返回的总是具体类型（而不是接口类型）。
与reflect.TypeOf类似，reflect.ValueOf返回的也是一个具体值，但它也可以包含一个接口值。
*/
func reflect_stdout() {
	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w)) // "*os.File"

	v := reflect.ValueOf(3) // 3, 一个reflect.Value
	fmt.Println(v)          // "3"
	fmt.Println(v.String()) // 注意: "<int Value>"
	/*
		另一个与reflect.Type类似的是，reflect.Value也满足
		fmt.Stringer，但除非Value包含的是一个字符串，否则String方法的
		结果仅仅暴露了类型。通常，你需要使用fmt包的%v功能，它对
		reflect.Value会进行特殊处理。
	*/
	fmt.Printf("%v\n", v) // "3"
}

func main() {
	reflect_demo()
	reflect_stdout()
}
