package main

import (
	"bytes"
	"fmt"
)

/* slice

slice代表变长的序列，slice中每个元素具有相同的类型。一个slice类型一般写作[]T，其中T代表slice元素的类型；
slice语法和数组比较像，但是没有固定的长度。
slice和数组之间有着紧密的联系。一个slice是一个轻量级的数据结构，提供了访问数组子序列（或全部）元素的功能。而
slice的底层的确引用了一个数组对象。一个slice由三个部分组成：指针、长度和容量。
指针指向slice第一个元素对应的底层数组元素的地址，要注意的是，不一定是数组的第一个元素。
长度对应slice的中元素的数目。容量一般是从slice的开始位置到底层数组的结尾位置（非元素结尾）。
内置函数len和cap可以返回slice的长度和容量。

make([]T, len)
make([]T, len, cap) // same as make([]T, cap)[:len]
*/

func slice_declare() {
	a := []int{1, 2, 3}
	b := []int{4: 5}
	c := make([]int, 6, 10)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

/* 多个slice之间可以共享底层数据，并且引用的底层数组部分区间可能重叠 */
func slice_share() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	b := a[4:7]
	c := a[6:9]
	fmt.Printf("%p, %v, len: %d, cap: %d\n", b, b, len(b), cap(b))
	fmt.Printf("%p, %v, len: %d, cap: %d\n", c, c, len(c), cap(c))
}

/* 如果切片操作超出cap(s)上限将导致一个panic，但超出len(s)则意味着扩展了slice，因为新的slice长度会变大 */
func slice_expand() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered in f", r)
		}
	}()

	a := []int{1, 2, 3, 4, 5}
	fmt.Printf("len: %d, cap: %d\n", len(a), cap(a))

	a = append(a, 6, 7, 8)
	fmt.Printf("len: %d, cap: %d\n", len(a), cap(a))

	a = a[:9]
	fmt.Printf("len: %d, cap: %d\n", len(a), cap(a))

	a = a[:20] // panic
	fmt.Printf("\n\n")
}

/* 函数内对slice进行追加不会影响到外部slice */
func slice_append(a []int) {
	a = append(a, 100, 200, 300)
}

func slice_update(a []int) {
	last := len(a) - 1
	a[last] = 100
}

// 新切片不增加长度
func slice_reslice_1(a []int) {
	last := len(a) - 1
	a = a[:last]
	a[last-1] = 200
}

// 新切片增加长度
func slice_reslice_2(a []int) {
	newLen := len(a) + 1
	a = a[:newLen]
	a[newLen-1] = 300
}

// reverse reverses a slice of ints in place.
func slice_reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

/*
slice在作为函数参数传递时是值传递，会在被调用函数内创建一个新的slice（即对底层数组创建一个新的slice别名）。
新的slice和原slice的指针都指向同样的底层数组。
如果新slice追加了元素或其它操作导致了它长度的增加，将不会影响到原slice（因为新slice会和原底层数组切割）。
如果新slice不扩充长度，只是对元素进行更新，则会影响到底层数组，进而影响到所有其它正在引用底层数组的slices。
*/
func slice_func_param() {
	a := make([]int, 0, 11)
	a = append(a, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	b := a[5:]
	fmt.Printf("a: %v\n", a)
	fmt.Printf("a: %v\n", b)

	slice_append(a)
	fmt.Println("after slice append in func slice_append:")
	fmt.Printf("a: %v\n", a)
	fmt.Printf("a: %v\n", b)
	fmt.Printf("不会影响原slice\n\n")

	slice_update(a)
	fmt.Println("after slice update in func slice_update:")
	fmt.Printf("a: %v\n", a)
	fmt.Printf("a: %v\n", b)
	fmt.Printf("不会影响原slice\n\n")

	slice_reslice_1(a)
	fmt.Println("after reslice in func slice_reslice_1 (not grow length):")
	fmt.Printf("a: %v\n", a)
	fmt.Printf("a: %v\n", b)
	fmt.Printf("不会影响原slice\n\n")

	slice_reslice_2(a)
	fmt.Println("after reslice in func slice_reslice_2 (grow length):")
	fmt.Printf("a: %v\n", a)
	fmt.Printf("a: %v\n", b)
	fmt.Printf("不会影响原slice\n\n")

	slice_reverse(a)
	fmt.Println("after slice reverse in func slice_reverse:")
	fmt.Printf("a: %v\n", a)
	fmt.Printf("a: %v\n", b)
	fmt.Printf("不会影响原slice\n\n")
}

/*
slice不能像数组一样，使用==操作符来判断两个slice是否含有全部相等的元素。
标准库中提供了高度优化的bytes.Equal函数来判断两个字节型slice是否相等。
对于其它类型的slice，我们必须自己展开所有元素进行比较。
*/
func slice_equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
func slice_compare() {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	c := []int{1, 2, 4}
	// fmt.Println(a == b) // compile error: invalid operation: cannot compare a == b (slice can only be compared to nil)
	fmt.Println(slice_equal(a, b))
	fmt.Println(slice_equal(a, c))

	d := []byte{1, 2, 3}
	e := []byte{1, 2, 3}
	f := []byte{1, 2, 4}
	fmt.Println(bytes.Equal(d, e))
	fmt.Println(bytes.Equal(d, f))

	/* slice唯一合法的比较操作是和nil比较 */
	fmt.Println(a == nil)
}

/*
slice的零值

一个零值的slice等于nil。一个nil值的slice并没有底层数组。
一个ni值的slice长度和容量都为0，但也有非nil值的slice的长度和容量都为0的，例如 []int{} 或 make([]int, 3)[3:]
与任意类型的nil值一样，我们可以用[]int(nil)类型转换表达式来生成一个对应类型的slice的nil值。

如果要测试一个slice是否为空，要用len(s)==0来判断，而不是用s==nil
*/
func slice_nil() {
	var a []int // len(a) == 0, a == nil
	fmt.Println(len(a), a == nil)

	a = nil // len(a) == 0, a == nil
	fmt.Println(len(a), a == nil)

	a = []int(nil) // len(a) == 0, a == nil
	fmt.Println(len(a), a == nil)

	a = []int{} // len(a) == 0, a != nil
	fmt.Println(len(a), a == nil)

	b := []byte(nil)
	fmt.Println(len(b), b == nil)

	c := []string(nil)
	fmt.Println(len(c), c == nil)

	if len(a) == 0 {
		fmt.Println("a is empty")
	}

	/* 一个nil值的slice行为和其它任意0长度的slice一样 */
	b = b[:]
}

func main() {
	slice_declare()
	slice_share()
	slice_expand()
	slice_func_param()
	slice_compare()
	slice_nil()
}
