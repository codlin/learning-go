package main

import (
	"fmt"
	"time"
)

/*
Goroutine

Go中的并发程序可以用两种方式实现：
1. goroutine和channel
2. 多线程共享内存

本文件是关于第一种类型的相关知识点。
goroutine和channel支持“顺序通信进程(communicationg sequential process)或被简称CSP。
CSP是一种现代的并发编程模型，在这种编程模型中值会在不同的编程实例（goroutine）中进行传递尽管大多数情况下仍然是被限制在单一实例中。

在Go语言中，每一个并发执行的单元叫做一个goroutine。当一个程序启动时，其主函数即在一个单独的goroutine里运行，被称作main goroutine。
新的goroutine都会用go关键字来创建。在语法上，go语句是一个普通的函数或方法调用前加上关键字go。
go语句会使一个函数在一个新创建的goroutine中运行。
*/
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func calc_fib() {
	// spinner和fib并发执行
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

/*
当值函数返回时，所有的goroutines都被直接打断，退出程序。
除了从主函数退出或直接终止程序外，没有其它的编程方法能让一个goroutine来打断另一个的执行。
但可以通过goroutine之间的通信来让一个goroutine来请求其它的goroutine，并让被请求的goroutine自行了断。
*/

func main() {
	calc_fib()
}
