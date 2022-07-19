package main

import "fmt"

func compare_nil() {
	// nil 和 nil 无法比较，会报编译错误
	// if nil == nil

	// nil 与 指针类型变量、通道、切片、函数、映射比较
	// 对于通道、切片、映射只有 var t T 或者手动赋值为 nil 时候(t = nil)，才能等于nil
	var a *int
	if a == nil {
		fmt.Println("a is nil")
	}

	var b chan *int
	if b == nil {
		fmt.Println("b is nil")
	}

	var c chan int
	if c == nil {
		fmt.Println("c is nil")
	}

	c1 := make(chan int, 0)
	if c1 == nil {
		fmt.Println("c1 is nil")
	}

	d := []*int{}
	if d == nil {
		fmt.Println("d is nil")
	}

	e := func() {
	}
	if e == nil {
		fmt.Println("e is nil")
	}

	var fn func()
	if fn == nil {
		fmt.Println("fn is nil")
	}

	var f map[int]int
	if f == nil {
		fmt.Println("f is nil")
	}

	g := make(map[int]int, 0)
	if g == nil {
		fmt.Println("g is nil")
	}
}

func nil_interface() {
	// nil 与 interface 比较
	var a interface{} // (T=nil, V=nil)
	var b *int        // (T=*int, V=nil)
	fmt.Printf("var a interface{} is nil: %v\n", a == nil)
	fmt.Printf("var b *int is nil: %v\n", b == nil)

	// a is nil, b is nil, but a is not equal to b
	fmt.Printf("a == b: %v\n", a == b) // false

	var pb interface{} = interface{}(b)
	fmt.Printf("pb == nil: %v\n", pb == nil) // false

	var c *int
	fmt.Printf("var c *int is nil: %v\n", c == nil)

	// b is nil, c is nil, b is equal c
	fmt.Printf("b == c: %v\n", b == c) // true

	// go 语言的接口变量有两个基础属性：Type 和 Value 。Type 指的是接口类型变量的底层类型，Value 是接口类型变量的底层值。
	// 两个变量只有在底层类型和值都相等时，在比较时才会认为相等。
	// 当 nil 与接口进行比较时，需要类型和值都为nil才相等
	d := 1
	e := 1
	c = &d
	fmt.Printf("c == e: %v\n", *c == e) // true. *c 和 e 都是 int 类型且值都为1

}

func main() {
	compare_nil()

	nil_interface()
}
