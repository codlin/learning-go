package main

import (
	"fmt"
	"math/cmplx"
)

/*
GO语言提供了两种类型的复数类型：complex64和complex128，分别对应float32和float64两种浮点精度。
内置的complex用于构建复数，内建的real和imag函数分别返回复数的实部和虚部。
*/

func complex1() {
	var x complex128 = complex(1, 2) // 1+2i
	var y complex128 = complex(3, 4) // 3+4i
	fmt.Println(x * y)
	fmt.Println(real(x * y))
	fmt.Println(imag(x * y))

	a := 1 + 2i
	b := 3 + 4i
	fmt.Println(a*b, a+b)

	/* 复数也提供了比较功能，可以用==和!=进行比较，只有在两个复数的实部和虚部都相等时它们才是相等的
	浮点数的的相等比较是危险的，需要特别小心处理精度问题。
	*/
	fmt.Println(x == a, x == b)

	// math.cmplx包提供了复数处理的许多函数
	fmt.Println(cmplx.Sqrt(-1))
}

func main() {
	complex1()
}
