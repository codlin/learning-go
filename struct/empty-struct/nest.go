package main

import "unsafe"

/**
空结构体本身不占用空间，但是作为某结构体内嵌字段时候，有可能是占用空间的：
1. 当空结构体是该结构体唯一的字段时，该结构体是不占用空间的，空结构体自然也不占用空间
2. 当空结构体作为第一个字段或者中间字段时候，是不占用空间的
3. 当空结构体作为最后一个字段时候，是占用空间的，大小跟其前一个字段保持一致
4. 当前空结构体作为数组元素和切片元素时不占空间
**/

type s1 struct {
}

type s2 struct {
	a struct{}
}

type s3 struct {
	_ struct{}
}

type s4 struct {
	a struct{}
	b byte
}

type s5 struct {
	b byte
	a struct{}
}

type s6 struct {
	b byte
	a struct{}
	c byte
}

type s7 struct {
	b byte
	_ struct{}
	c byte
}

type s8 struct {
	b int64
	a struct{}
}

type s9 struct {
	b int64
	a struct{}
	c int64
}

type s10 struct {
	a struct{}
	b struct{}
}

func main() {
	println(unsafe.Sizeof(s1{}))
	println(unsafe.Sizeof(s2{}))
	println(unsafe.Sizeof(s3{}))
	println(unsafe.Sizeof(s4{}))
	println(unsafe.Sizeof(s5{}))
	println(unsafe.Sizeof(s6{}))
	println(unsafe.Sizeof(s7{}))
	println(unsafe.Sizeof(s8{}))
	println(unsafe.Sizeof(s9{}))
	println(unsafe.Sizeof(s10{}))

	var a [10]int
	println(unsafe.Sizeof(a)) // 80

	var b [10]struct{}
	println(unsafe.Sizeof(b)) // 0

	var c = make([]struct{}, 10)
	println(unsafe.Sizeof(c)) // 24，即slice header的大小

	d := [10]struct{}{{}, {}}
	println(unsafe.Sizeof(d)) // 0，作为数组元素

	e := []struct{}{{}, {}}
	println(unsafe.Sizeof(e)) // 24，即slice header的大小

}
