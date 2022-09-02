package main

import (
	"fmt"
	"unsafe"
)

func assignment() {
	a := []byte{'h', 'e', 'l', 'l', 'o'}
	b := a[2:3]
	c := a[2:3:3]
	fmt.Println(string(a), string(b), string(c)) // 输出 hello l l

	type sliceHeader struct {
		array unsafe.Pointer // 底层数据数组的指针
		len   int            // 切片长度
		cap   int            // 切片容量
	}
	ptrA := (*sliceHeader)(unsafe.Pointer(&a))
	ptrB := (*sliceHeader)(unsafe.Pointer(&b))
	ptrC := (*sliceHeader)(unsafe.Pointer(&c))
	fmt.Printf("切片%s: 底层数组地址=0x%x, 长度=%d, 容量=%d\n", "a", ptrA.array, ptrA.len, ptrA.cap)
	fmt.Printf("切片%s: 底层数组地址=0x%x, 长度=%d, 容量=%d\n", "b", ptrB.array, ptrB.len, ptrB.cap)
	fmt.Printf("切片%s: 底层数组地址=0x%x, 长度=%d, 容量=%d\n", "c", ptrC.array, ptrC.len, ptrC.cap)

	b = append(b, 'w')                           // 有副作用，改变的是切片 a 的值
	c = append(c, 'r')                           // 没有副作用，c生成了单独的切片。
	fmt.Println(string(a), string(b), string(c)) // 输出 helwo lw lr
	ptrA = (*sliceHeader)(unsafe.Pointer(&a))
	ptrB = (*sliceHeader)(unsafe.Pointer(&b))
	ptrC = (*sliceHeader)(unsafe.Pointer(&c))
	fmt.Printf("切片%s: 底层数组地址=0x%x, 长度=%d, 容量=%d\n", "a", ptrA.array, ptrA.len, ptrA.cap)
	fmt.Printf("切片%s: 底层数组地址=0x%x, 长度=%d, 容量=%d\n", "b", ptrB.array, ptrB.len, ptrB.cap)
	fmt.Printf("切片%s: 底层数组地址=0x%x, 长度=%d, 容量=%d\n", "c", ptrC.array, ptrC.len, ptrC.cap)
}

func main() {
	assignment()
}
