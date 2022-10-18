package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"syscall"
)

/*
类型断言

类型断言是一个使用在接口值上的操作。语法上它看起来像x.(T)，被称为断言类型，其中x表示一个接口的类型，T表示一个类型。
一个类型断言检查它操作对象的动态类型是否和断言的类型匹配。这里有两种可能：
 1. 如果断言的类型T是一个具体类型，然后类型断言检查x的动态类型是否和T相同。如果这个检查成功了，类型断言的结果是x的动态值，当然它的类型是T。
 2. 如果断言的类型T是一个接口类型，然后类型断言检查是否x的动态类型满足T。如果这个检查成功了，动态值没有获取到(extracted)，这个结果仍然是
    x和T是有相同类型和值成分（components)的接口值。换句话说，对一个接口类型的类型断言改变了类型的表达方式，改变了可以获取的方法集合。但是
    它保护了接口值内部的动态类型和值的部分。
*/
type stringReader struct {
}

func (s *stringReader) Read(p []byte) (n int, err error) {
	return 0, nil
}
func type_assert() {
	var w io.Writer
	w = os.Stdout
	rw := w.(io.ReadWriter) // success: *os.File has both Read and Write
	fmt.Printf("%T\n", rw)

	// w = new(stringReader)
	// rw = w.(io.ReadWriter) // panic: *stringReader has no Write method

	/* 为了防止检查失败时发生panic，断言可以返回第二个结果，它是一个布尔型的标识，用于指明是否成功 */
	t, ok := w.(io.Writer)
	fmt.Printf("%T\t%t\n", t, ok)
}

/*
如果断言操作的对象是一个nil接口值，那么不论断言的类型是什么类型，这个断言都会失败。
我们几乎不需要对一个更少限制性的接口（更少的方法集合）断言，因为它表现的就像赋值操作一样，除了对于nil接口值的情况。
*/
func nil_type_assert() {
	var w io.Writer
	t, ok := w.(io.Writer)
	fmt.Printf("%T\t%t\n", t, ok)
}

/*
我们可以基于类型断言区别错误类型。通常的做法是提供一个具体的错误类，实现内置包（buildin）中的error接口。
在错误处理时我们可以对错误error进行类型断言，判断具体的错误类型，从而得到更多的信息。
这个错误类型也可以是多层级的，我们可以定义一个错误接口作为基础错误，该接口内嵌error接口，这样可以强制具体的错误类实现这个接口。
我们可以在这个基础错误接口中里面加入额外的方法来限定具体错误类型必须实现这些接口。然后我在错误处理时可以用这个基础错误类型进行“兜底”。
但在Go语言中，一般实现单个的error接口应该就够了，不需要多层级的error机制。
*/
var ErrNotExist = errors.New("file does not exist")

type PathError struct {
	Op   string
	Path string
	Err  error
}

func (e *PathError) Error() string {
	return e.Op + " " + e.Path + " " + e.Err.Error()
}

func IsNotExist(err error) bool {
	if pe, ok := err.(*PathError); ok {
		err = pe.Err
	}
	return err == syscall.ENOENT || err == ErrNotExist
}

/*
就像两种类型的断言一样，接口也被以两种不同的方式使用：
 1. 一个接口的方法表达了实现这个接口的具体类型的间接相似性，但隐藏了代表的细节和这些具体类型本身的操作。
    所以这种情况重在方法上，而不是具体类型（类型断言类似于面向对象中的向上类型转换）
 2. 利用一个接口值可以持有各种具体类型的能力并将这个接口认为是这些类型的联合（union）。
    类型断言用来动态的区别具体类型并且对每一种情况都按需对待。
    这种情况重点在于具体的类型满足这个接口，而不是在于接口方法（如果它有的话），并且没有任何的信息隐藏。
    （这类似于面向对象的子类型多态，类型断言类似于向下类型转换）
*/
func sqlQuote(x interface{}) string {
	// 使用switch-case可以简化类型断言的if-else链
	switch x := x.(type) {
	case nil:
		return "NULL"
	case int, uint:
		return fmt.Sprintf("%d", x) // x has type interface{} here.
	case bool:
		if x {
			return "TRUE"
		}
		return "FALSE"
	case string:
		return sqlQuoteString(x) // (not shown)
	default:
		panic(fmt.Sprintf("unexpected type %T: %v", x, x))
	}
}

func sqlQuoteString(str string) string {
	return str
}

func main() {
	type_assert()
	nil_type_assert()
}
