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

/**
type Pointer
Pointer represents a pointer to an arbitrary type. 
There are four special operations available for type Pointer 
that are not available for other types:
Pointer代表了任意类型的一个指针。
指针类型有4种特别的操作不适用于其它类型:
- A pointer value of any type can be converted to a Pointer.
- A Pointer can be converted to a pointer value of any type.
- A uintptr can be converted to a Pointer.
- A Pointer can be converted to a uintptr.

Pointer therefore allows a program to defeat the type system and read and write arbitrary memory.
It should be used with extreme care.
指针允许程序击败类型系统并读写任意内存，因此使用时要格外小心。

The following patterns involving Pointer are valid.
Code not using these patterns is likely to be invalid today or to become invalid in the future.
Even the valid patterns below come with important caveats.
下面涉及到指针Pointer的模式是有效的。
不使用这些模式的代码很可能在今天无效，或者是将来变得无效。
即使是下面有效的模式也有重要的警告。

Running "go vet" can help find uses of Pointer that do not conform to these patterns,
but silence from "go vet" is not a guarantee that the code is valid.
运行"go vet"可以帮助找到不符合这些模式的Pointer用途。
但是"go vet"的沉默并不能保证代码有效。

模式1：
(1) Conversion of a *T1 to Pointer to *T2.

Provided that T2 is no larger than T1 and that the two share an equivalent memory layout,
this conversion allows reinterpreting data of one type as data of another type. 
假设 T2 不大于 T1 并且两者共享相同的内存布局，
这种转换允许将一种类型的数据重新解释为另一种类型的数据。

An example is the implementation of math.Float64bits:
func Float64bits(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}

模式2：
(2) Conversion of a Pointer to a uintptr (but not back to Pointer).

Converting a Pointer to a uintptr produces the memory address of the value pointed at, as an integer.
The usual use for such a uintptr is to print it.
将 Pointer 转换为 uintptr 会生成指向值的内存地址，作为整数。
这种 uintptr 的通常用途是打印它。

Conversion of a uintptr back to Pointer is not valid in general.
通常，将 uintptr 转换回 Pointer 是无效的。

A uintptr is an integer, not a reference. 
Converting a Pointer to a uintptr creates an integer value with no pointer semantics.
Even if a uintptr holds the address of some object, the garbage collector will not update 
that uintptr's value if the object moves, nor will that uintptr keep the object from being reclaimed.
uintptr 是一个整数，而不是一个引用。
将Pointer 转换为 uintptr 会创建一个没有指针语义的整数值。
即使 uintptr 保存了某个对象的地址，垃圾收集器也不会在对象移动时更新该 uintptr 的值，该 uintptr 也不会阻止该对象被回收。

The remaining patterns enumerate the only valid conversions from uintptr to Pointer.
其余模式枚举了从 uintptr 到 Pointer 的唯一有效转换。

模式3：
(3) Conversion of a Pointer to a uintptr and back, with arithmetic.
    使用算术将指针转换为 uintptr 并返回。

If p points into an allocated object, it can be advanced through the object by conversion to uintptr, 
addition of an offset, and conversion back to Pointer.
如果 p 指向分配的对象，则可以通过转换为 uintptr、添加偏移量并转换回 Pointer 来“穿过”对象。

p = unsafe.Pointer(uintptr(p) + offset)

The most common use of this pattern is to access fields in a struct or elements of an array:
// equivalent to f := unsafe.Pointer(&s.f)
f := unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Offsetof(s.f))

// equivalent to e := unsafe.Pointer(&x[i])
e := unsafe.Pointer(uintptr(unsafe.Pointer(&x[0])) + i*unsafe.Sizeof(x[0]))

It is valid both to add and to subtract offsets from a pointer in this way.
It is also valid to use &^ to round pointers, usually for alignment. 
In all cases, the result must continue to point into the original allocated object.
以这种方式从指针添加和减去偏移量都是有效的。
使用 &^ 对指针进行舍入也是有效的，通常用于对齐。
在所有情况下，结果必须继续指向原始分配的对象。

Unlike in C, it is not valid to advance a pointer just beyond the end of its original allocation:
与 C 不同，将指针前移刚好超出其原始分配的末尾是无效的：

// INVALID: end points outside allocated space.
var s thing
end = unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Sizeof(s))

// INVALID: end points outside allocated space.
b := make([]byte, n)
end = unsafe.Pointer(uintptr(unsafe.Pointer(&b[0])) + uintptr(n))
Note that both conversions must appear in the same expression, with only the intervening arithmetic between them:

// INVALID: uintptr cannot be stored in variable
// before conversion back to Pointer.
u := uintptr(p)
p = unsafe.Pointer(u + offset)
Note that the pointer must point into an allocated object, so it may not be nil.

// INVALID: conversion of nil pointer
u := unsafe.Pointer(nil)
p := unsafe.Pointer(uintptr(u) + offset)

模式4：
(4) Conversion of a Pointer to a uintptr when calling syscall.Syscall.
    调用syscall.Syscall 时将Pointer 转换为uintptr。

The Syscall functions in package syscall pass their uintptr arguments directly to the operating system, 
which then may, depending on the details of the call, reinterpret some of them as pointers. 
That is, the system call implementation is implicitly converting certain arguments back from uintptr to pointer.
syscall 包中的 Syscall 函数将它们的 uintptr 参数直接传递给操作系统，
然后，根据调用的细节，它可能会将其中一些重新解释为指针。
也就是说，系统调用实现隐式地将某些参数从 uintptr 转换回指针。

If a pointer argument must be converted to uintptr for use as an argument, 
that conversion must appear in the call expression itself:
如果必须将指针参数转换为 uintptr 才能用作参数，则该转换必须出现在调用表达式本身中：

syscall.Syscall(SYS_READ, uintptr(fd), uintptr(unsafe.Pointer(p)), uintptr(n))

The compiler handles a Pointer converted to a uintptr in the argument list of a call to a function 
implemented in assembly by arranging that the referenced allocated object, 
if any, is retained and not moved until the call completes, even though from the types alone it would 
appear that the object is no longer needed during the call.
编译器处理在对汇编中实现的函数调用的参数列表中转换为 uintptr 的指针，方法是保留引用的已分配对象（如果有的话），
并且在调用完成之前不移动该对象，即使仅从类型来看似乎在调用期间不再需要该对象。
(注：我们可以认为编译器针对每个syscall.Syscall函数调用中的每个被转换为uintptr类型的非类型安全指针实参添加了一些指令，
从而保证此非类型安全指针所引用着的内存块在此调用返回之前不会被垃圾回收和移动。)

For the compiler to recognize this pattern, the conversion must appear in the argument list:

// INVALID: uintptr cannot be stored in variable
// before implicit conversion back to Pointer during system call.
u := uintptr(unsafe.Pointer(p))
syscall.Syscall(SYS_READ, uintptr(fd), u, uintptr(n))

模式5：
(5) Conversion of the result of reflect.Value.Pointer or reflect.Value.UnsafeAddr from uintptr to Pointer.

Package reflect's Value methods named Pointer and UnsafeAddr return type uintptr instead of unsafe.
Pointer to keep callers from changing the result to an arbitrary type without first importing "unsafe".
However, this means that the result is fragile and must be converted to Pointer immediately after making the call,
in the same expression:
包 reflect 的名为 Pointer 和 UnsafeAddr 的 Value 方法返回类型 uintptr 而不是 unsafe。
这是为了防止调用者在不首先导入“unsafe”的情况下将结果更改为任意类型的指针。
但是，这意味着结果是脆弱的，必须在同一个表达式中将调用后立即转换为 Pointer，

// MUST DO
p := (*int)(unsafe.Pointer(reflect.ValueOf(new(int)).Pointer()))

As in the cases above, it is invalid to store the result before the conversion:
// INVALID: uintptr cannot be stored in variable
// before conversion back to Pointer.
u := reflect.ValueOf(new(int)).Pointer()
p := (*int)(unsafe.Pointer(u))

模式6：
(6) Conversion of a reflect.SliceHeader or reflect.StringHeader Data field to or from Pointer.

As in the previous case, the reflect data structures SliceHeader and StringHeader declare the field Data 
as a uintptr to keep callers from changing the result to an arbitrary type without first importing "unsafe".
However, this means that SliceHeader and StringHeader are only valid when interpreting the content of an actual slice 
or string value.
与前面的情况一样，反映数据结构 SliceHeader 和 StringHeader 声明字段数据
作为一个 uintptr 来防止调用者在不首先导入“unsafe”的情况下将结果更改为任意类型。
然而，这意味着 SliceHeader 和 StringHeader 仅在解释实际切片的内容或字符串值时有效。

var s string
hdr := (*reflect.StringHeader)(unsafe.Pointer(&s)) // case 1
hdr.Data = uintptr(unsafe.Pointer(p))              // case 6 (this case)
hdr.Len = n
In this usage hdr.Data is really an alternate way to refer to the underlying pointer in the string header,
not a uintptr variable itself.
在此用法中，hdr.Data 实际上是引用字符串标头中的底层指针的替代方法，而不是 uintptr 变量本身。

In general, reflect.SliceHeader and reflect.StringHeader should be used only as *reflect.SliceHeader 
and *reflect.StringHeader pointing at actual slices or strings, never as plain structs.
A program should not declare or allocate variables of these struct types.
一般来说，reflect.SliceHeader 和 reflect.StringHeader 应该只用作指向实际切片或字符串的 *reflect.SliceHeader 
和 *reflect.StringHeader，绝不能用作普通结构。程序不应声明或分配这些结构类型的变量。

// INVALID: a directly-declared header will not hold Data as a reference.
var hdr reflect.StringHeader
hdr.Data = uintptr(unsafe.Pointer(p))
hdr.Len = n
s := *(*string)(unsafe.Pointer(&hdr)) // p possibly already lost
**/

func main() {
	sizeof()
	size_align_offset_of()
}
