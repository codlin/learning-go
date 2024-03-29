package main

import (
	"fmt"

	"mypackage/bar"
	"mypackage/foo"
)

/*
GO语言的包和其它语言的库和模块的概念类似，都是为了支持模块化、封装、单独编译和代码宠用。
一个包的源代码保存在一个或者多个以.go以文件后缀名的源文件中，通常一个包所在的目录路径的后缀是包导入的路径。
例如包gopl.io/ch1/helloworld对应的目录路径是$GOPATH/src/gopl.io/ch1/helloworld。
每一个包都通过唯一的字符串进行标识，它称为导入路径，它们用在import声明中。
对于准备共享或公开的包，导入路径需全局唯一，为了避免冲突，除了标准库中的包外，其它包的导入路径应该以互联网域名作为路径开始，
这样也方便查找包。

在每个go源文件开始都需要进行包声明。主要的目的是当包被其它包引入的时候作为默认的标识符（包名）。
通常，包名是导入路径的最后一段，于是，即使导入路径不同的两个包，二者也可以拥有相同的名字。
关于包名是导入路径的”最后一段，这个有三个例外：
1. 不管包的导入路径是什么，如果该包定义了一条命令（可执行的go程序），那么它总是使用名称main
2. 目录中可能有一些名字以_test.go结尾，包名中会出现以_test结尾。这样一个目录中有两个包：一个普通的，加上一个外部测试的。
   _test后缀告诉go test两个包都要构建，并且知名文件属于哪个包。
3. 有一些管理工具会在导入路径的尾部追加版本号，如“gopkg.in/yaml.v2”。包名不包含后缀，因此这个情况下包名为yaml。

一个go源文件可以在package声明后面和第一个非导入声明语句前面紧接着包含零个和多个import声明。
每一个导入可以单独执行一条导入路径，也可以通过圆括号括起来的列表一次导入多个包。

导入的包可以通过空行分组，就像本文件的导入方式似的。这类分组通常表示不同领域和方面的包。
导入顺序不重要，但按照惯例每一组都要按照字母排序（gofmt和goimports工具，或者IDE工具一般都会自动进行分组排序）。

如果需要把两个名字一样的包（如math/rand和crypto/rand）导入到第三个包中，导入声明就必须至少为其中的一个指定一个替代
名字来避免冲突。这叫作重命名导入。
import (
	"crypto/rand"
	mrand "math/rand"
)
替代名字仅影响当前文件。
重命名导入在没有冲突时也是非常有用的。如果有时用到自动生成的代码，导入的包名字非常冗长，使用一个代替的名字可能更方便。
*/

/*
空白导入

如果导入的名字没有在文件中引用，就会产生一个编译错误。但有时候我们必须导入一个包，这仅仅是为了利益其副作用：对包级别的
变量进行初始化表达式求值，并执行它的init函数。
为了防止“未使用的导入“错误，我们必须使用一个重命名导入，它使用一个替代的名字_，这表示导入的内容未空白标识符。
通常情况下，空白标识符不可能被引用。
import _ "image/png"
这称为空白导入。多数情况下，它用来实现一个编译时的机制，使用空白引用导入额外的包，来开启主程序中可选的特性。
上面的例子，png包会在自己的init函数中注册png类型以及编解码函数到更通用的image包中，这样编译器就会在编译时把png的包也
编译进去，从而支持png格式的编解码。否则，即使可以编译链接通过，那么在运行时也会因为没有png格式的编解码而报运行时错误。
数据库的驱动也是类似的机制。这两个例子都可以查看<The Go Programming Language>中10.5章节的内容。
*/

/*
包及其命名

当创建一个包时，使用简单的名字，但不要短到像加了密一样。尽可能保持可读性和无歧义。
例如，不要把一个辅助工具包命名未util，使用imageutil或ioutil等名称更加具体和清晰。
避免选择经常用于相关局部变量的包名，或者迫使使用者使用重名导入，例如使用以path命名的包。
*/

/*
包成员的可见性

每个包都对应一个独立的名字空间。包可以控制哪些名字是外部可见的，简单的规则是：
如果一个名字以大写字母开头，那么该名字是可导出的，外部可以引用。
*/

/*
包的初始化顺序

包的初始化首先是解决包级变量的依赖顺序，然后按照包级变量的声明出现的顺序依次初始化
*/
var a = b + c // a 第三个初始化
var b = f()   // b 第二个初始化
var c = 1     // c 第一个初始化

func f() int {
	fmt.Println("f was called")
	return c + 1
}

/*
如果包中包含多个源文件，那么将按照它们发给编译器的顺序进行初始化。GO语言的构建工具会首先将文件根据文件名排序，然后依次调用编译器编译。
对于在包级别声明的变量，如果有初始化表达式则用表达式初始化，还有一些没有初始化表达式的，可以用一个特殊的init()初始化函数来简化初始化工作。
每个文件可以包含多个init()函数。
*/
var d string

func init() {
	d = "hello"
	fmt.Println("init 1 was called")
}

/* 这样的init初始化函数除了不能被调用或引用外，其行为和普通函数类似。 在每个文件中的init初始化函数，在程序开始执行时按照它们的声明顺序被自动调用。 */
func init() {
	c = 2
	fmt.Println("init 2 was called")
}

/*
每个包在依赖解决的前提下，以导入声明的顺序初始化，每个包只能被初始化一次。
关于包的导入的更多例子请参见：https://www.liwenzhou.com/posts/Go/import_local_package_in_go_module/
*/

func main() {
	fmt.Println(a)
	fmt.Printf("foo.D %s\n", foo.D)
	fmt.Printf("bar.C %s\n", bar.C)
}
