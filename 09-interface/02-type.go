package main

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
