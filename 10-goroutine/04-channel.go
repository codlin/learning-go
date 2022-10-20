package main

import (
	"fmt"
	"strings"
	"time"
)

/*
如果说goroutines是一个Go程序的并发体的话，channel就是它们之间的通信机制。
一个channel是一个通信机制，它可以让一个goroutine向另外一个goroutine发生值。
每一个channel都是一个特定类型值的管道，称为channel元素类型。一个可以发送int类型数据的channel一般写为chan int。
使用内置的make函数，可以创建一个channel。
以最简单的方式调用make创建是不带缓存的channel，但是我们也可以指定第二个整型参数，对应channel的容量。
*/
func create_chan() {
	a := make(chan int)
	fmt.Printf("%T\n", a)

	b := make(chan struct{}, 0)
	fmt.Printf("%T\n", b)

	c := make(chan map[string]int, 5)
	fmt.Printf("%T\n", c)
}

/*
和map类似，一个channel是对make创建的底层数据的引用。
当我们复制一个channel或者用于函数参数传递时，我们只是拷贝了一个channel引用，因为调用者和被调用着将引用同一个channel对象。
和其它引用类型一样，channel的零值也是nil。

两个相同类型的channel可以使用==运算符比较。如果两个channel引用的是同一个对象，那么比较的结果为真。一个channel也可以与nil比较。
*/
func cmp_chan() {
	a := make(chan int)
	fmt.Printf("%p\n", a)
	fmt.Println(a == a) // true

	b := make(chan int)
	fmt.Printf("%p\n", b)
	fmt.Println(a == b) // false

	fmt.Println(a == chan_ref(a)) // true

	var c chan int
	fmt.Printf("c == nil ? %t\n", c == nil)
}

func chan_ref(c chan int) chan int {
	fmt.Printf("chan_ref c: %p\n", c)
	return c
}

/*
一个channel有发送和接受两个主要操作，都是通信行为。
一个发送语句将一个值从一个goroutine发送到另一个执行接收操作的goroutine。
发送和接收两个操作都是使用<-运算符。
在发送语句中，<-分割channel和要发送的值。<-左边不是channel，右边是要发送的值。
在接收语句中，<-写在channel对象之前，即<-右边是channel，左边是用于接收值的变量。但左边也可以没有变量，一个不接收结果的接收操作也是合法的。
*/
func chan_send_recv() {
	a := make(chan int)
	go func(c chan int) {
		for i := 0; i < 5; i++ {
			c <- i
			fmt.Printf("sent data: %d\n", i)
			time.Sleep(1000 * time.Millisecond)
		}
	}(a)

	b := -1
	for i := 0; i < 5; i++ {
		b = <-a
		fmt.Printf("received data: %d\n", b)
		time.Sleep(1015 * time.Millisecond)
	}
}

/*
channel还支持close操作，用于关闭channel，随后对基于该channel的任何发送操作都将导致panic异常。
对应已经被close的channel执行接收数据操作，依然可以接收到之前已经发送的数据；如果channel中已经没有数据，则产生这个channel元素类型的零值数据。
使用内置函数close可以关闭channel。
*/
func chan_close() {
	a := make(chan int, 5)
	defer close(a)
}

func recv_on_closed_chan() {
	a := make(chan int, 5)
	go func(c chan int) {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}(a)

	time.Sleep(6 * time.Second)

	fmt.Println("receive data on a close chan")
	for i := 0; i < 7; i++ {
		b := <-a
		fmt.Printf("received data: %d\n", b)
	}
}

/* 下面例子中，函数运行完成后，发送数据的goroutine也会被中断发送 */
func chan_send_abort() {
	a := make(chan int)
	go func(c chan int) {
		for i := 0; i < 5; i++ {
			c <- i
			fmt.Printf("sent data: %d\n", i)
			time.Sleep(1000 * time.Millisecond)
		}
	}(a)

	b := -1
	for i := 0; i < 3; i++ {
		b = <-a
		fmt.Printf("received data: %d\n", b)
		time.Sleep(1015 * time.Millisecond)
	}
}

/*
不带缓存的channel
*/

func main() {
	fmt.Printf("%s\n", strings.Repeat("=", 64))
	create_chan()

	fmt.Printf("%s\n", strings.Repeat("-", 64))
	cmp_chan()

	fmt.Printf("%s\n", strings.Repeat("-", 64))
	chan_send_recv()

	fmt.Printf("%s\n", strings.Repeat("-", 64))
	recv_on_closed_chan()

	fmt.Printf("%s\n", strings.Repeat("-", 64))
	chan_send_abort()
	time.Sleep(3 * time.Second)

	fmt.Printf("%s\n", strings.Repeat("=", 64))
}
