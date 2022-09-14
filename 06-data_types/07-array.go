package main

import (
	"fmt"
	"strings"
	"unsafe"
)

/*
array

数组是一个由固定长度的相同类型的元素组成的序列，一个数组有0和或多个元素组成。
因为数组的长度是固定不能改变的，在Go语言中很少使用数组。
和数组对应的类型是切片（slice），它是可以增长和收缩的序列，slice的功能也更为灵活，但是要理解slice，需要先了解数组。
*/
func array_def() {
	/* 默认情况下每个数组都被初始化为元素类型对应的零值 */
	var a [3]int
	fmt.Printf("a: %v\n", a) // [0 0 0]

	/* 通过字面值显式初始化数组 */
	var b [3]int = [3]int{1, 2, 3}
	fmt.Printf("b: %v\n", b) // [1 2 3]

	/* 数组按照顺序进行初始化，如果没有指定默认值则用零值填充 */
	var c [3]int = [3]int{1, 2}
	fmt.Printf("c: %v\n", c) // [1 2 0]

	/* [...]是一个语法糖，可以根据后面指定的值自动计算数组的初始化大小 */
	d := [...]int{1, 2, 3}
	fmt.Printf("%T, d: %[1]v\n", d) // [3]int [1 2 3]

	/* 数组的长度是数组类型的一部分，即不同长度的数组不是同一个类型 */
	e := [3]int{1, 2, 3}
	// e = [4]int{1, 2, 3， 4} // complie error: cannot use ([4]int literal) (value of type [4]int) as [3]int value in assignment
	e = a // 同类型数组可以替换

	/* 数组可以通过下标索引，合法长度为0到数组的长度-1。 内存函数len可以返回数组中元素的个数。*/
	e[1] = 200
	fmt.Printf("e: %v\n", e) // [0 200 0]

	/* 数组可以通过 for或for range 遍历 */
	for idx, v := range a {
		fmt.Printf("index: %d\tvalue: %d\n", idx, v)
	}

	/* 除了通过提供顺序初始化值序列来初始化数组外，还可以通过指定一个索引和对应值列表的方法初始化数组 */
	f := [...]int{5: 25, 10: 100} // 定义了一个长度为11的整型数组，除了指定索引的值外，其余都为0
	fmt.Printf("f type: %T\t, value: %[1]v\n", f)

	/* 不同大小的数组是不能比较的，即使它们的元素类型相同。
	只有同类的数组才可以比较，并且在所有元素都相等时才认为数组是相等的 */
	g := [4]int{1, 2, 3, 4}
	h := [...]int{1, 2, 3, 4}
	// fmt.Println(a == g) // invalid operation: cannot compare a == g (mismatched types [3]int and [4]int)
	fmt.Println(b == c, g == h) // false, true
}

/*
根据数组元素数量的不同，Go语言编译器会做出两种不同的优化：
1. 对长度小于等于4的数组，会直接上数组中的元素放在栈上(在不考虑逃逸分析的情况下)；
2. 对应长度大于4的数组，会将数组的元素放置到静态区并在运行时取出
*/
func arry_len_le_4() {
	/* 声明一个长度为4的数组
	编译器会在编译时将其改写为：
	var a [4]int64
	&a[0] = 1
	&a[1] = 2
	&a[2] = 3
	&a[3] = 4
	即所有的元素都放置在栈上
	*/
	a := [4]int64{1, 2, 3, 4}
	p1 := &a
	p2 := &a[1]
	p3 := &a[2]
	p4 := &a[3]

	// 0xc0000a0000 0xc0000a0008 0xc0000a0010 0xc0000a0018
	// 0xc000098018 0xc000098020 0xc000098028 0xc000098030
	fmt.Printf("%p %p %p %p\n%p %p %p %p", p1, p2, p3, p4, &p1, &p2, &p3, &p4)
}

func arry_len_more_4() {
	/* 声明一个长度为5的数组 a
	编译器会在编译时将其改写，可以大致理解为（伪代码）：
	statictmp_0 := readonlystaticname(t) // statictmp_0 为编译器在静态存储区给这个数组分配的唯一的名字，并且数组的元素会在编译器就存储在静态存储区
	statictmp_0[0] = 1
	statictmp_0[1] = 2
	statictmp_0[2] = 3
	statictmp_0[3] = 4
	statictmp_0[4] = 5

	var a [5]int64
	a = statictmp_0
	即所有的元素都放置在栈上
	*/
	a := [...]int64{1, 2, 3, 4, 5}

	fmt.Printf("a: %v\n", a) // [1 2 3 4 5]
}

func array_pointer() {
	a := [...]int64{1, 2, 3, 4}
	b := a
	c := &a

	fmt.Println(strings.Repeat("=", 64))
	fmt.Printf("a: %v\n", a) // [1 2 3 4]
	fmt.Printf("b: %v\n", b) // [1 2 3 4]

	b[1] = 100
	fmt.Printf("b修改后: %v\n", b)      // [1 100 3 4]
	fmt.Printf("b修改后， a的值: %v\n", a) // [1 2 3 4]
	fmt.Printf("a 的内存地址: %p\n", &a)
	fmt.Printf("b 的内存地址: %p\n", &b)
	/* 由此可见，将一个数组赋值给另一个变量，是值copy，即会生成一个新的数组。
	且数组 b 的内容在编译期已经确定，从内存地址看两个数组相邻地址差32个字节（栈上和堆上都有可能）。
	在不考虑内存逃逸的情况下，对小于等于4的数组直接在栈上分配空间
	*/

	fmt.Println(strings.Repeat("-", 32))

	c[2] = 200
	fmt.Printf("c: %v\n", *c)  // [1 2 200 4]
	fmt.Printf("a: %v\n", a)   // [1 2 200 4]
	fmt.Printf("&a: %p\n", &a) // 数组的首地址
	fmt.Printf("c: %p，以%%p打印会打印出它的值：数组a的首地址\n", c)
	fmt.Printf("c: %v，以%%v打印会打印出它引用的值：数组a的内容。用&代表引用，不会显示指针\n", c)
	fmt.Printf("&c: %p，变量c在栈上的地址，它的内容时指向a的指针\n", &c)
	fmt.Printf("*c: %v，打印*c会打出c指向的内容，即a的内容\n", *c)

	fmt.Println(strings.Repeat("-", 64))

	d := [...]int64{1, 2, 3, 4, 5}
	f := &d
	fmt.Printf("d: %v\n", d)   // [1 2 3 4 5]
	fmt.Printf("f: %v\n", f)   // &[1 2 3 4 5]
	fmt.Printf("&d: %p\n", &d) // 数组的首地址
	fmt.Printf("f: %p，以%%p打印会打印出它的值：数组d的首地址\n", f)
	fmt.Printf("f: %v，以%%v打印会打印出它引用的值：数组d的内容。用&代表引用，不会显示指针\n", f)
	fmt.Printf("&f: %p，变量f在栈上的地址，它的内容是指向d的指针\n", &f)
	fmt.Printf("*f: %v，打印*f会打出f指向的内容，即d的内容\n", *f)
	fmt.Println(strings.Repeat("=", 64))

	// TODO: 疑问，如何打印出数组变量的地址？譬如上面&a打印的是数组的首地址，但我理解栈上应该有个存放局部变量a的8字节大小的地址
}

func struct_array() {
	a := [200]struct{}{}
	fmt.Printf("&a: %p, unsafe.Sizeof(a): %d, sizeof(a[0]): %d\n", &a, unsafe.Sizeof(a), unsafe.Sizeof(a[0]))
}

func main() {
	array_def()
	arry_len_le_4()
	arry_len_more_4()
	array_pointer()
	struct_array()
}

/* READ ME
下面是生成对应指令架构集的SSA命令：
```shell
GOSSAFUNC=arry_len_4 CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o arry_len_4.html ./array.go
```
*/

/*
```c
#include <stdio.h>
#include <stdlib.h>

int main()
{
	int a[]={1, 2, 3};
	int *b = a;
	int *c = a;
	int *d = (int *)malloc(sizeof(int)*3);
	int *e = d;

    printf("a: %p\t&a: %p\t&b: %p\t&c: %p\n", a, &a, &b, &c); // a: 0x7ffe85a686f4	&a: 0x7ffe85a686f4	&b: 0x7ffe85a686e8	&c: 0x7ffe85a686e0

	printf("d: %p\t&d: %p\t&e: %p\n", d, &d, &e); // d: 0x207b260	&d: 0x7fffb98a0b88	&e: 0x7fffb98a0b80

	free(d);
	d = NULL;
    return 0;
}
````
*/
