package main

import "fmt"

/*
可变参数函数

参数数量可能变的函数称为可变参数函数。典型的例子就是fmt.Printf函数。

在声明可变参数函数时，需要在参数列表的最后一个参数类型之前加上省略符号“...”，这表示该函数会接受任意数量的该类型参数。
*/
func sum(vals ...int) int {
	total := 0
	for _, v := range vals {
		total += v
	}
	return total
}

func test_sum() {
	/*下面的代码调用sum时：
	调用者隐式的创建一个数组，并将元素参数复制到数组中，再把数组的一个切片作为参数传递给被调用函数。
	*/
	fmt.Println(sum())
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3))

	/* 如果原始参数已经是切片时，需要在后面加上“...”，代表展开切片 */
	sum([]int{1, 2, 3}...)
}

/*
虽然在可变参数函数内部，...int型参数的行为看起来很像切片类型，但实际上，可变参数函数和以切片为参数的函数是不同的。
*/
func f(...int) {}
func g([]int)  {}
func variadic_slice() {
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", g)
}

func main() {
	test_sum()
	variadic_slice()
}
