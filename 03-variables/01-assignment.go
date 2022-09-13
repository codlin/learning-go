package main

import "fmt"

func ret_vars() (int, int) {
	return 1, 2
}

/*
变量的赋值
*/

func assign_1() {
	a := 1
	a++

	b := new(int)
	*b = 2

	c := []int{1, 2, 3}
	c[0] *= 10

	fmt.Println(a, b, c)
}

/* 元组的赋值 */
func assign_2() {
	a := 1
	b := 2

	// 元组赋值时先计算等号右边的值，然后再进行赋值
	a, b = b, a%b // 先计算a%b的值为1，然后再赋值（√）。而不是先把b的值赋给a，然后再计算a%b的值为0 (x)

	c := []int{1, 2, 3}
	c[0], c[1] = c[1], c[0]

	// 有些表达式会产生多值，比如调用一个有多个返回值的函数。当这样一个函数调用出现在元组赋值右边的表达式时（右边不能再有其它表达式），左边变量的数目必须和右边一致。
	d, e := ret_vars()
	fmt.Println(d, e)

	// 下面的代码编译错误，如果右侧为函数，不能再有其它表达式
	// d, e, f := ret_vars(), 2

}

// 求最大公约数
func gcd(x, y int) int {
	fmt.Printf("calculate gcd of %d and %d\n", x, y)

	for y != 0 {
		x, y = y, x%y
		fmt.Printf("x=%d, y=%d\n", x, y)
	}
	fmt.Printf("gcd is %d\n", x)
	return x
}

/* 可赋值性 Assignability */
func assignability() {
	/*
		赋值语句是显式的赋值形式，但程序中还有很多地方会发生隐式的赋值行为：
		函数调用会隐式将调用参数的值赋值给函数的参数变量，返回语句也会隐式地将返回操作的结果赋值给结果变量。
		一个复合类型的字面量也会产生隐式赋值行为。
	*/
	a := gcd(4344, 4390)

	b := []string{"foo", "bar"}
	c := map[string]string{"hello": "world", "foo": "bar"}

	/*
	 不管是隐式还是显式赋值，都要求函数左边和右边最终求到的值必须具有相同的数据类型。
	 可赋值性的规则对于不同的数据类型有着不同的要求：
	 1. 简单的规则是类型必须匹配
	 2. nil可以赋值给任何指针或引用类型的变量。
	 3. 常量有跟灵活的赋值规则，这样可以避免不必要的显式的类型转换
	*/
	var d *int = nil
	var e *string = nil
	f := 3.14
	g := 4

	/* 对于两个值是否可以用==或!=比较的能力也和赋值能力相关；对于任何类型的相等比较，第二个值必须和第一个值类型对应的变量是可赋值的，反之亦然 */
	// if d == e {} // 报类型不匹配错误
	// if f == g {} // 报类型不匹配错误

	fmt.Println(a, b, c, d, e, f, g)
}

func main() {
	gcd(231432, 589347)
}
