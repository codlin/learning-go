package main

import (
	"fmt"
	"os"
	"regexp"
	"runtime"
)

/*
Panic

运行时错误会引起panic。
一般而言，当panic发生时，程序会中断运行，并立即执行在该goroutine中被defer的函数。随后，程序崩溃并输出日志信息。
日志信息包括panic value和函数调用的堆栈跟踪信息。panic value通常是某种错误信息。
对于每个goroutine，日志信息中都会有与之相对应的发生panic时的函数调用堆栈跟踪信息。
不是所有的panic异常都来自运行时，直接调用内置的panic函数也会引发panic异常。panic函数接受任何值作为参数。

虽然Go的panic机制类似于其它语言的异常，但panic的使用场景有一些不同。由于panic会引起程序的崩溃，因此panic一般用于严重的程序错误。
对于大部分的代码中的漏洞，我们应该使用Go的错误机制，而不是用panic，尽量避免程序的崩溃。
在健壮的程序中，任何可以预料到的错误，如不正确的输入，错误的配置或失败的I/O操作都应该被优雅的处理，最好的处理方式就是使用Go的错误机制。

*/

/*
考虑regexp.Compile函数，该函数将正则表达式编译成有效的可匹配格式。当输入的正则表达式不合法时，该函数会返回一个错误。
当调用者明确的知道正确的输入不会引起函数错误时，要求调用者检查这个错误是不必要和累赘的。我们应该假设函数的输入一直合法，
就如前面的断言一样：当调用者输入了不应该出现的输入时，触发panic异常。

在程序源码中，大多数正则表达式是字符串字面值（string literals），因此regexp包提供了包装函
数regexp.MustCompile检查输入的合法性。
*/
func Compile(expr string) (*regexp.Regexp, error) { /* ... */ return nil, nil }
func MustCompile(expr string) *regexp.Regexp {
	re, err := Compile(expr)
	if err != nil {
		panic(err)
	}
	return re
}

// panic test
func f1(x int) {
	if x == 0 {
		return
	}
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f1(x - 1)
}

/*
当f(0)被调用时，发生panic异常，之前被延迟执行的的3个fmt.Printf被调用。
程序中断执行后，panic信息和堆栈信息才会被输出。
（以其它语言的经验来看，应该是调用f(0)时产生中断，后续的3个defer应该不会被打印，但事实上在Go语言里不是这样。）
在Go语言里，延迟defer函数的调用在释放堆栈信息之前。
*/
func f2(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f2(x - 1)
}

func print_stack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

func panic_test1() {
	f1(3)
}

func panic_test2() {
	f2(3)
}

func main() {
	defer print_stack()
	panic_test1()
	panic_test2()
}
