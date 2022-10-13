package main

import (
	"bytes"
	"fmt"
)

/*
接口类型

Go语言中有很多单方法接口的习惯，然后将多个单接口组合。
*/

type Reader interface {
	Read(p []byte) (n int, err error)
}
type Writer interface {
	Write(p []byte) (n int, err error)
}
type Closer interface {
	Close() error
}

type ReadWriter interface {
	Reader
	Writer
}

type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}

// 但也可以不使用嵌入组合，完全定义
type ReadWriter2 interface {
	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
}

// 或者混合使用
type ReadWriter3 interface {
	Read(p []byte) (n int, err error)
	Writer
}

/*
如果一个类型用于一个接口所需要的所有方法，那么这个类型就实现了这个接口。
接口指定的规则非常简单：表达一个类型属于某个接口只要这个类型实现了这个接口。
就像信封封装和隐藏信件一样，接口类型封装和隐藏具体类型和它的值。
一个有更多方法的接口类型，会告诉我们更多关于它的值持有的信息，并且对实现它的类型要求更加严格。
那么关于interface{}类型，它没有任何方法，请讲出哪些具体的类实现了它？
*/

/*
interface{}被称为空接口类型是不可或缺的。因为空接口类型对实现它的类型没有要求，所以我们可以将任意一个值赋给空接口类型。
对于interface{}类型持有的值，我们不能直接对它操作，因为interface{}没有任何方法。但我们可以通过类型断言的方式获取它的值。
*/
func interface_value() {
	var any interface{}
	any = true
	any = 1234

	i := any.(int)
	fmt.Printf("i:%d\n", i)

	any = "hello"
	any = map[string]int{"one": 1}
	any = new(bytes.Buffer)
	fmt.Println(any)
}

/*
每一个具体类型的组基于它们相同的行为可以表示成一个接口类型。
不像基于类的语言，它们一个类实现的接口集合需要显式的定义，在Go语言里我们可以在需要的时候
定义一个新的抽象或者特定特点的组，而不需要修改具体类型的定义。当具体的类型来源于不同的作者时这种方式特别有用。
*/
// stringReader 实现了Reader.Read接口，因此它可以看成是Reader类型
// 我们不需要要像其它语言一样必须显式定义它需要实现哪些接口，譬如这样：
// type stringReader struct implements Reader {}
type stringReader struct {
}

func (s *stringReader) Read(p []byte) (n int, err error) {
	return 0, nil
}

func interface_impl() {
	var r Reader = &stringReader{}
	r.Read([]byte{})
}

func main() {
	interface_value()
}
