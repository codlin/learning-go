package main

import "fmt"

/* slice

slice代表变长的序列，slice中每个元素具有相同的类型。一个slice类型一般写作[]T，其中T代表slice元素的类型；
slice语法和数组比较像，但是没有固定的长度。
slice和数组之间有着紧密的联系。一个slice是一个轻量级的数据结构，提供了访问数组子序列（或全部）元素的功能。而
slice的底层的确引用了一个数组对象。一个slice由三个部分组成：指针、长度和容量。
指针指向slice第一个元素对应的底层数组元素的地址，要注意的是，不一定是数组的第一个元素。
长度对应slice的中元素的数目。容量一般是从slice的开始位置到底层数组的结尾位置（非元素结尾）。
内置函数len和cap可以返回slice的长度和容量。
*/

/* 多个slice之间可以共享底层数据，并且引用的底层数组部分区间可能重叠 */
func slice_share() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	b := a[4:7]
	c := a[6:9]
	fmt.Printf("%p, %v, len: %d, cap: %d\n", b, b, len(b), cap(b))
	fmt.Printf("%p, %v, len: %d, cap: %d\n", c, c, len(c), cap(c))
}
func main() {
	slice_share()
}
