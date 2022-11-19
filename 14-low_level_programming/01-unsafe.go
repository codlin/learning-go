package main

import (
	"fmt"
	"unsafe"
)

/*
unsafe包

包unsafe是很神奇的，虽然它像普通的包并且像普通包那样导
入，但是它事实上是由编译器实现的。它提供了对语言内置特性的
访问功能，而这些特性一般是不可见的，因为它们暴露了Go详细的
内存布局。把这些特性单独放在一个包中，就使得它们的本来就不
频繁的使用场合变得更加引入注目。另外，一些环境下，出于安全
原因，unsafe包的使用是受限制的。

包unsafe广泛使用在和操作系统交互的低级包（比如runtime、
os、syscall和net）中，但是普通程序从来不需要使用它。
*/

/*
unsafe.Sizeof

函数unsafe.Sizeof报告传递给它的参数在内存中占用的字节长度，
这个参数可以是任何类型的表达式，但表达式不会被计算。
Sizeof调用返回一个unitptr类型的常量表达式，所以这个结果可以
作为数组类型的维度或者用于计算其它常量。

Sizeof仅报告每个数据结构固定部分占用的内存占用字节长度，例如指针或者
字符串的长度，但不会报告诸如字符串内容这种简介内容。
*/
func sizeof() {
	fmt.Println(unsafe.Sizeof(float64(0))) // 8
	// 对于字符串，只报告了字符串数据结构头部的占用自己长度
	fmt.Println(unsafe.Sizeof("hello"))                           // 16
	fmt.Println(unsafe.Sizeof("hello world"))                     // 16
	fmt.Println(unsafe.Sizeof("hello, welcome to golang world!")) // 16
	// 对于slice，只报告了结构占有的长度，不报告实际长度
	a := make([]int, 100)
	fmt.Println(unsafe.Sizeof(a)) // 24
	for i := 0; i < 50; i++ {
		a = append(a, i)
	}
	fmt.Println(unsafe.Sizeof(a)) // 24

}

/*
非聚合类型的典型长度如下所示，当然准确的长度随工具
链的不同而不同。为了可移植性，我们将以字来表示引用类型（或
者包含引用的类型）的长度。在32位系统上，字的长度是4个字节；
而在64位系统上，字的长度是8个字节。

在类型的值在内存中对齐的情况下，计算机的加载或者写入会
很高效。例如，一个两字节值（如int16）的地址应该是一个偶数，
一个四字节值（如rune）的地址应该是4的整数倍，一个八字节值
（如float64、uint64或者64位指针）的地址应该是8的整数倍。更大
整数倍的对齐很少见，即使是像complex128这种大的数据类型。

聚合类型（结构体或数组）的值的长度至少是它的成员
或元素长度之和，并且由于“内存间隙”的存在，或许比这个更大一
些。内存空位是由编译器添加的未使用的内存地址，用来确保连续
的成员或者元素相对于结构体或数组的起始地址是对齐的。
Type 										Size
bool 										1 byte
intN, uintN, floatN, complexN 				N/8 bytes (for example, float64 is 8 bytes)
int, uint, uintptr 							1 word
*T 											1 word
string 										2 words (data, len)
[]T 										3 words (data, len, cap)
map 										1 word
func 										1 word
chan 										1 word
interface 									2 words (type, value)

语言规范并没有要求成员声明的顺序对应内存中的布局顺序，
所以在理论上，编译器可以自由安排。尽管这样说，但是实际上没
人这样做。如果结构体成员的类型是不同的，那么将相同类型的成
员定义在一起可以更节约内存空间。

关于内存对齐，可以参见 CSAPP 一书的相关章节。
*/

/*
unsafe.Alignof

函数unsafe.Alignof报告它参数类型所要求的对齐方式。
和Sizeof一样，它的参数可以是任意类型的表达式，并且返回一个常
量。典型地，布尔类型和数值类型对齐到它们的长度（最大8字
节），而其他的类型则按字对齐。
*/

/*
unsafe.Offsetof

函数unsafe.Offsetof，计算成员f相对于结构体x起始地址的偏移
值，如果有内存空位，也计算在内，该函数的操作数必须是一个成
员选择器x.f
*/
type x1 struct {
	a bool
	b int16
	c []int
}

type x2 struct {
	b int16
	a bool
	c []int
}

type x3 struct {
	b int16
	c []int
	a bool
}

type x4 struct {
	c []int
	a bool
	b int16
}

type x5 struct {
	a bool
	c []int
	b int16
}

type x6 struct {
	c []int
	b int16
	a bool
}

func size_align_offset_of() {
	y1, y2, y3, y4, y5, y6 := x1{}, x2{}, x3{}, x4{}, x5{}, x6{}
	fmt.Printf("x1: %d\t%d\tx1.a: %d\t%d\t%d\tx1.b: %d\t%d\t%d\tx1.c: %d\t%d\t%d\n",
		unsafe.Sizeof(y1), unsafe.Alignof(y1),
		unsafe.Sizeof(y1.a), unsafe.Alignof(y1.a), unsafe.Offsetof(y1.a),
		unsafe.Sizeof(y1.b), unsafe.Alignof(y1.b), unsafe.Offsetof(y1.b),
		unsafe.Sizeof(y1.c), unsafe.Alignof(y1.c), unsafe.Offsetof(y1.c),
	)
	fmt.Printf("x2: %d\t%d\tx2.a: %d\t%d\t%d\tx2.b: %d\t%d\t%d\tx2.c: %d\t%d\t%d\n",
		unsafe.Sizeof(y2), unsafe.Alignof(y2),
		unsafe.Sizeof(y2.a), unsafe.Alignof(y2.a), unsafe.Offsetof(y2.a),
		unsafe.Sizeof(y2.b), unsafe.Alignof(y2.b), unsafe.Offsetof(y2.b),
		unsafe.Sizeof(y2.c), unsafe.Alignof(y2.c), unsafe.Offsetof(y2.c),
	)
	fmt.Printf("x3: %d\t%d\tx3.a: %d\t%d\t%d\tx3.b: %d\t%d\t%d\tx3.c: %d\t%d\t%d\n",
		unsafe.Sizeof(y3), unsafe.Alignof(y3),
		unsafe.Sizeof(y3.a), unsafe.Alignof(y3.a), unsafe.Offsetof(y3.a),
		unsafe.Sizeof(y3.b), unsafe.Alignof(y3.b), unsafe.Offsetof(y3.b),
		unsafe.Sizeof(y3.c), unsafe.Alignof(y3.c), unsafe.Offsetof(y3.c),
	)
	fmt.Printf("x4: %d\t%d\tx4.a: %d\t%d\t%d\tx4.b: %d\t%d\t%d\tx4.c: %d\t%d\t%d\n",
		unsafe.Sizeof(y4), unsafe.Alignof(y4),
		unsafe.Sizeof(y4.a), unsafe.Alignof(y4.a), unsafe.Offsetof(y4.a),
		unsafe.Sizeof(y4.b), unsafe.Alignof(y4.b), unsafe.Offsetof(y4.b),
		unsafe.Sizeof(y4.c), unsafe.Alignof(y4.c), unsafe.Offsetof(y4.c),
	)
	fmt.Printf("x5: %d\t%d\tx5.a: %d\t%d\t%d\tx5.b: %d\t%d\t%d\tx5.c: %d\t%d\t%d\n",
		unsafe.Sizeof(y5), unsafe.Alignof(y5),
		unsafe.Sizeof(y5.a), unsafe.Alignof(y5.a), unsafe.Offsetof(y5.a),
		unsafe.Sizeof(y5.b), unsafe.Alignof(y5.b), unsafe.Offsetof(y5.b),
		unsafe.Sizeof(y5.c), unsafe.Alignof(y5.c), unsafe.Offsetof(y5.c),
	)
	fmt.Printf("x6: %d\t%d\tx6.a: %d\t%d\t%d\tx6.b: %d\t%d\t%d\tx6.c: %d\t%d\t%d\n",
		unsafe.Sizeof(y6), unsafe.Alignof(y6),
		unsafe.Sizeof(y6.a), unsafe.Alignof(y6.a), unsafe.Offsetof(y6.a),
		unsafe.Sizeof(y6.b), unsafe.Alignof(y6.b), unsafe.Offsetof(y6.b),
		unsafe.Sizeof(y6.c), unsafe.Alignof(y6.c), unsafe.Offsetof(y6.c),
	)
	/* 虽然它们的名字叫作unsafe，但是这些函数本身是安全的，并
	且在做内存优化的时候，它们对理解程序底层内存布局很有帮助。 */
}

/*
unsafe.Pointer


*/
func main() {
	sizeof()
	size_align_offset_of()
}
