package main

import (
	"fmt"
	"math"
)

/*
浮点数表示法

IEEE浮点标准用V=(-1)^s x M x 2^E的形式来表示一个数：
- 符号 s 决定是正数（s=0）还是负数（s=1），而对于数值0的符号位解释作为特殊的情况处理。
- 有效数 M 是一个二进制小数，它的范围在1~2-ε之间，或者0~1-ε之间。
- 指数 E 是2的幂（可能是负数），它的作用是对浮点数加权。
浮点数的位被划分为3个区域，以编码这些值：
- 一个单独的符号位s直接编码符号s
- k位的指数域exp=ek-1...e1e0（k-1为下标）编码指数E
- n位小数域frac=fn-1...f1f0（n-1为下标）编码有效数M，但是被编码的值也依赖于指数值是否等于0
在单精度浮点格式（float32）中，s、exp和frac域分别为1位，k=8和n=23，产生一个32位的值。
在双精度浮点格式（float64）中，s、exp和frac域分别为1位，k=11和n=52，产生一个64位的值。
给定*位*表示，根据exp的值，被编码的值分成3种不同情况：
- 规格化值
- 非规格化值
- 特殊值

1. 规格化值
这个是最普遍的情况，当exp的值既不是全0也不是全1的情况。在这种情况中，指数域解释为表示偏置（biased）形式的有符号整数。
指数的值`E=e-Bias`，其中e是无符号数，其位表示为ek-1...e1e0，而Bias是一个等于2^(k-1)-1（单精度是127，双精度是1023）的偏置值。
由此产生指数的范围，对于单精度是-126~+127（e除去0和255，和Bias相减），双精度是-1022~+1023。
小数域frac解释为描述小数值f，其中0<=f<1，其二进制表示为0.fn-1...f1f0（n-1为下标），也就是二进制小数点在最高有效位左边。
有效数M定义为`M=1+f`。有时，这种方式也叫做隐含的以1开头的表示，因为我们可以把M看成一个二进制表达式为1.fn-1...f1f0的数字。
我们总能够调整E，使有效数M在范围1<=M<2之中，那么这种方法是一种免费获得一个额外精度的技巧————既然第一位总等于1，那么就不需要显式的表示它了。

2. 非规格化
当指数为全0时，所表示的值就是非规格化形式的。在这种情况下，指数值是`E=1-Bias`，而有效数M的值是M=f，也就是小数域的值，不包含隐含的开头1。
为什么指数`E=1-Bias`，而不是-Bias?
因为这种方式提供了一种非规格化值平滑转到规格化值的方法（简单的说就是M=f，和规格化值相比1没了，具体详见CASPP第二章浮点数部分）。
非规格化数有两个目的：
  - 它们提供了一中表示0的方法，因为使用规格化值表示方法，必须总是M>=1，所以无法表示0.实际上，+0.0的浮点表示的位模式全为0，我们得到M=f=0。
    而-0.0除了符号位为1外，其余也都是0，根据IEEE浮点标准，值+0.0和-0.0在某些方面是认为是不同的（？？），而在其它方面是相同的。
  - 非规格化数的另外一个功能是用来表示非常接近于0的数。

3. 特殊数值
这类数值是指数全为1是出现的。
- 当小数域全为0时，得到的值表示无穷。当s=0时是+∞，当s=1时是-∞。
- 当小数域不为0时，结果值被称为'NaN'，意思是Not a number。

GO语言提供了两种精度的浮点数：float32和float64
常量math.MaxFloat32: ~3.4e38，math.SmallestNonzeroFloat32: ~1.4e-45
常量math.MaxFloat64: ~1.8e308，math.SmallestNonzeroFloat64: ~4.9e-324
*/
func float_maxmin() {
	fmt.Printf("float32 max: %e, min: %e\n", math.MaxFloat32, math.SmallestNonzeroFloat32)
	fmt.Printf("float32 max: %e, min: %e\n", math.MaxFloat64, math.SmallestNonzeroFloat64)
}

/*
浮点数的格式化

浮点数可以用下面几种格式化参数打印：
%g 将以更紧凑的方式打印，并提供足够的精度
%f 以浮点数打印
%e 以带指数的形式打印
三种方式都可以指定打印的宽度和控制打印精度。
*/
func format() {
	float32_f()
	float64_f()
}

/*
浮点数的精度和舍入
*/

// float32类型可以提供大约6个十进制数的精度（即小数点后6位是准确的）
func float32_f() {
	const e float32 = 3.1415926
	fmt.Printf("float32: %[1]e, %.7[1]f, %[1]g\n", e)
	const r float32 = 31415926.
	fmt.Printf("float32: %[1]e, %.7[1]f, %[1]g\n", r)
}

// float64类型可以提供大约15个十进制数的精度（即小数点后15位是准确的）
func float64_f() {
	const e float64 = 3.1415926141592614159261415926
	fmt.Printf("float64: %[1]e, %.17[1]f, %[1]g\n", e)
	const r float64 = 31415926141592614159261415926.
	fmt.Printf("float64: %[1]e, %.17[1]f, %[1]g\n", r)
}

// TODO: 浮点数的舍入

/* 浮点数运算 */
func float_cal() {
	fmt.Println("+0.0 == -0.0", +0.0 == -0.0)

	// 浮点数的字面常量
	a, b := 1.0, +0.0
	fmt.Printf("%e, %e\n", a, b)

	// 1/+0 被定义为+∞
	infinity := a / b
	fmt.Println(infinity) // +∞

	// 1/-0 被定义为-∞
	infinity = a / (-b)
	fmt.Println(infinity) // -∞

	// 0/0 被定义为NaN
	fmt.Println(b / b) // NaN

	/* 函数math.IsNaN可以用来测试一个数是不是NaN */
	nan := b / b
	fmt.Printf("Is NaN: %v\n", math.IsNaN((nan)))

	/* math.NaN 可以用来表示一个非法的结果，但测试一个数是不是NaN是危险的
	因为NaN和任何数都是不相等的。在浮点数中，NaN、正无穷、负无穷都不是唯一的，每个都有非常多种的位模式表示
	*/
	fmt.Println(b/b == b/b, nan == nan, nan > nan, nan < nan, nan == math.NaN())

	// 处理浮点失败的做法
	v, ok := compute()
	if ok {
		fmt.Printf("got value: %e\n", v)
	}
}

/* 如果一个函数返回的浮点数结果可能失败，那么最好的做法是用单独的标志位报告失败 */
func compute() (value float64, ok bool) {
	failed := false
	result := 3.14
	if failed {
		return 0, false
	}
	return result, true
}

/* 浮点数加法满足交换律，但不满足结合律 */
func float_assoc() {
	fmt.Println("illustrate float_assoc")

	a, b := 3.14, 1e10
	// 浮点数满足交换律
	fmt.Println(a+b, a+b == b+a)

	/* 但不满足结合律 */
	// float64精度，在和大数相加减时，因为舍入，会导致值的丢失
	fmt.Println(a+b-b, a+(b-b), (a+b-b) == (a+(b-b))) // 3.1399993896484375 3.14 false

	// 在单精度下，直接导致 3.14 丢失
	var c, d float32 = 3.14, 1e10
	fmt.Println(c+d-d, c+(d-d), (c+d-d) == (c+(d-d))) // 0 3.14 false

	/* 不满足结合律的一个错误优化
	在大多数情况下没有问题，但是有时会产生不同于原始的值，可能差别很细小，
	在大多数应用种不太重要，但是在科学计算等对精度要求很高的场景，需要重视。
	可以参考CSAPP中关于火箭发射精度误差的例子。
	*/
	x := 3.14 + 3.15 + 2.71
	y := 3.15 + 2.71 + 1.41

	t := 3.15 + 2.71
	m := 3.14 + t
	n := t + 1.41
	fmt.Printf("x=%g, y=%g, m=%g, n=%g, x==m: %v, y==n: %v\n", x, y, m, n, x == m, y == n) // y==n: false

	/* 在大多数情况下浮点数也有逆元，-x+x=0，但是对于无穷例外 */
	e := 0.0
	fmt.Println(-a+a, -1/e+1/e, -1/e+1/e == 0, -1/e+1/e == math.NaN()) // 0 NaN false false

	/* 对于任何x，都有 NaN+x=NaN */
	fmt.Printf("3.14+3.14/0 == 3.14/0 is %v\n", a+a/e == a/e)
}

/*
浮点数加法满足单调性

如果a>=b，对于任何a和b的值，除了x!=NaN，都有 x+a >= x+b。
（当x=NaN时，对于任何x都有 NaN+x=NaN，任何NaN都不相等，所以不满足上面的条件）
*/
func float_mono() {
	fmt.Println("illustrate float_mono")

	a, b, c := 1.330, 1.331, -0.768
	fmt.Println(a+c < b+c) // true

	// 在不失最小精度的情况下可以得到正确的单调性结果
	var x, y, z float32 = 1.1234561, 1.1234562, -1.1234561
	fmt.Println(x+z, y+z, x+z < y+z) // 0 1.1920929e-07 true

	// 在失去最小精度的情况下无法满足单调性，但是仍然可以满足逆元
	x, y, z = 1.12345611, 1.12345612, -1.1234561
	fmt.Println(x+z, y+z, x+z < y+z) // 0 0 false
}

/* 浮点数乘法满足交换律，但不满足结合律和在加法上的分配律 */
func float_mul() {
	fmt.Println("illustrate float_mul")

	a, b, c := 1.1, 1e20, 3.3
	fmt.Println(a*b, a*b == b*a) //

	fmt.Println(a*b*c, a*b*c == b*a*c, a*b*c == b*c*a, a*b*c == a*(c*b)) //
}

func main() {
	float_maxmin()
	format()
	float_cal()
	float_assoc()
	float_mono()
	float_mul()
}
