package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
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

**判断两个接口对象是否相等，不但要看它们的类型type是否一样，还要看值value是否一样**

一个接口值被描述为nil还是非nil是基于它的动态类型来判断的。
**这一点很关键，如果它的动态类型是nil，那么这个类型值就是nil；否则，
如果它的动态类型不为nil，那么即使接口值中value部分为nil，则这个接口值也不为nil**

在下面4个语句中，变量w得到3个不同的值：（初始化值和最后的值相同）
*/

func interface_value() {
	var w io.Writer
	fmt.Println(w == nil) // true
	fmt.Println(w == w)   // true

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
编译器必须产生代码以获取具体类型的方法地址，然后间接的调用这个地址。这个调用的接受者是接口动态值的拷贝，
上面w是*bytes.Buffer类型时，w.Write相当于是(*bytes.Buffer).Write
      w
      +---------------------+
 type |    *bytes.Buffer    |              bytes.Buffer
      +---------------------+             +--------------------+
value |         *-----------|------------>|  data []byte       |
      +---------------------+             +--------------------+
*/

/*
一个接口可以持有任意大的动态值。
接口可以使用==和!=进行比较。两个接口相等仅当它们都是nil值（类型和值都是nil）或者它们的动态类型相同并且动态值也是按照这个动态类型的==操作相等时。
因为接口值是可以比较的，所以它们可以用在map的键或者作为switch语句的操作数。
但是如果两个接口的动态类型相同，但是它们的值是不可比较的（比如切片），将它们进行比较就会失败并且产生panic：
*/
func interface_compare() {
	var x interface{} = []int{1, 2, 3}
	fmt.Println(x)
	// fmt.Println(x == x) // panic. interface value is not comparable.
}

/*
考虑到这点，接口类型是非常与众不同的。其它类型要么是安全的可比较（如基本类型和指针），要么是完全不可比较的（如切片slices，映射类型maps，和函数functions）
但在比较接口值或者包含了接口值的聚合（aggregate）类型时，我们必须要意识到潜在的panic。同样的风险也存在于使用接口作为map键或者switch操作数时。
只能比较你非常确定它们的动态值是可以比较类型的接口值。我们可以通过%T来打印接口值的动态类型，或者使用反射机制来获取接口动态值的类型。
*/

/*
警告（Cavate）：
一个包含nil指针的接口不是nil接口。
An Interface Containing a Nil Pointer Is Non-Nil
参见function：interface_value()
下面的代码，在调用f(buf)时如果buf是nil，那么在f中，out并不是nil，因为它的动态类型是*bytes.Buffer，但动态值是nil，所以out.Write会引发panic
*/
func non_nil(i int) {
	var buf *bytes.Buffer // 应该修改为：var buf io.Writer
	if i&1 == 0 {
		buf = new(bytes.Buffer)
	}
	f(buf) // NOTE: subtly incorrect!
}

// If out is non‐nil, output will be written to it.
func f(out io.Writer) {
	if out != nil {
		out.Write([]byte("done!\n")) // panic: nil pointer dereference
	}
}

/* Benchmark: 接口动态调用和类型直接调用的性能测试对比 */

func main() {
	interface_value()
	interface_compare()
}
