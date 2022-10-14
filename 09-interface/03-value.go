package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
接口类型值 a value of an interface type, or interface value

概念上讲，一个接口类型的值，即接口值，由两个构成要素(componets)：
一个具体的类型(concret type)，和这个具体类型的值(a value of that type)
       interface value
      +---------------+
 type |               |
      +---------------+
value |               |
      +---------------+
这两部分称之为接口的**动态类型**和**动态值**

对于像Go这样的静态语言，类型是编译时概念，所以一个类型不是一个值。
在我们的概念模型里，为每一种类型提供信息的一系列值称之为**类型描述符**（type descriptors），例如它的名字和方法集。
在接口值中，类型部件（type component）是由适当的类型描述符表示。
*/

// 名字stringReader和方法Read，都称之为类型描述符，它们构成了stringReader这个类型。
// **stringReader类型**是Reader接口的具体类型，即接口值中type的值是由**stringReader**类型表示。
// 而接口的值value部分则是由具体的变量值决定。
type Reader interface {
	Read(p []byte) (n int, err error)
}

type stringReader struct {
}

func (s *stringReader) Read(p []byte) (n int, err error) {
	return 0, nil
}

/*
在Go语言中，变量总被一个定义明确的值初始化，即使接口类型也不例外。
对于一个接口类型的零值，它的类型和值这两部分都设置为nil
      w
      +---------------+
 type |      nil      |
      +---------------+
value |      nil      |
      +---------------+

一个接口值被描述为nil还是非nil是基于它的动态类型来判断的。
这一点很关键，如果它的动态类型是nil，那么这个类型值就是nil；否则，
如果它的动态类型不为nil，那么即使接口值中value部分为nil，则这个接口值也不为nil

在下面4个语句中，变量w得到3个不同的值：（初始化值和最后的值相同）
*/

func interface_value() {
	var w io.Writer
	fmt.Println(w == nil) // true

	var f *os.File
	fmt.Println(f == nil) // true
	w = f                 // w的动态类型被设置为了*of.File，但值还是nil（因为f的值为nil）
	fmt.Println(w == nil) // false
	fmt.Println(w.Write([]byte("hello world")))

	w = new(bytes.Buffer)
	fmt.Println(w == nil) // false
	fmt.Println(w.Write([]byte("hello world")))

	w = nil
	fmt.Println(w == nil) // true
}

/*
通常在编译期间，我们不知道接口值的动态类型是什么，所以一个接口上的调用必须使用动态分配。因为不是直接调用，
编译器必须产生代码以获取具体类型的方法地址，然后间接的调用这个地址。
*/

func main() {
	interface_value()
	s := "[1][1.0]ABCDEFG"
	idx := strings.LastIndexByte(s, ']')
	s = s[idx+1:]
	fmt.Println(idx, s)
}
