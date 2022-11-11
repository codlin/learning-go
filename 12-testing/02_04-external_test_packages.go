package main

/*
外部测试包（External testing packages）

考虑包net/url，这个包提供了URL解析功能，还有net/http，这个包提供了Web服务器和HTTP客户端库。
如我们所知，高级的net/http包依赖于低级的net/url包。然而，net/url包中的一个测试是用来演示URL
和HTTP库之间进行交互的例子。
换句话说，低级别包的测试导入了高级别包。
在net/url包中声明这个测试函数会导致包循环引用，Go规范禁止循环引用。
     +-----------------+
     |    net/http     +
     +-----------------+
        /|\        |
circle!  |         |
	     |        \|/
     +-----------------+
     |    net/url      +
     +-----------------+

我们通常将测试函数定义在外部测试包中来解决这个问题。也就是说，在net/url目录中有一个_test.go文件，
它的包声明不是url，而是url_test。这个额外的后缀_test告诉go test工具，它应该被单独地编译一个包，
这个包仅包含这些文件，然后运行它的测试。
为了帮助理解，你可以认为这个外部测试包的导入路径是net/url_test，但事实上它无法通过该路径以及其他任何路径导入。
由于外部测试在一个单独的包里，因此它们可以引用一些依赖于被测试包的帮助包；这是包内测试无法做到的。
						+-----------------+
						|   net/url_test  +
						+-----------------+
                        /     /
     +-----------------+     /
     |    net/http     +    /
     +-----------------+   /
             /|\          /
              |          /
	          |         /
     +-----------------+
     |    net/url      +
     +-----------------+

为了避免包循环导入，外部测试包允许测试用例，尤其是集成测试用例，自由的导入其它的包，就像一个应用程序一样。

可以使用go list工具来汇总一个包目录中哪些是产品代码，哪些是包内测试，以及哪些是外部测试。
以fmt包为例：
1. GoFiles是包含产品代码的文件列表，这些文件是go build命令将编译进你程序的代码。
$ go list -f={{.GoFiles}} fmt
输出：[doc.go errors.go format.go print.go scan.go]
2. TestGoFiles是包含内部测试包的文件列表，这些文件以_test.go结尾，仅在编译测试的时候使用。
$ go list -f={{.TestGoFiles}} fmt
输出：[export_test.go]
3. XTestGoFiles是包外部测试列表，譬如url_test，这些文件必须引用url包才能使用它。同样，它也仅在测试时使用。
$ go list -f={{.XTestGoFiles}} fmt
输出：
[errors_test.go example_test.go fmt_test.go gostringer_example_test.go scan_test.go
stringer_example_test.go stringer_test.go]

有时候，外部包在测试时需要对内部的包有特殊的访问权限，例如为了避免循环引用，一个白盒测试必须存在于一个单独的包中。
在这种情况下，我们有一个小技巧：在包内测试文件_test.go中添加一些函数声明，将包内部的功能暴露给外部测试。
这些文件也因此测试提供了包的一个“后门”，如果一个源文件的存在的唯一目的就在于此，并且自己不包含任何的测试，它们
一般被称为export_test.go。更详细的例子可以参考<The Go Programming Language>第11.2.4章节。

*/
