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
	fmt.Println("create_chan")

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
	fmt.Println("cmp_chan")

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
	fmt.Println("chan_send_recv")

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
	}
}

/*
channel还支持close操作，用于关闭channel，随后对基于该channel的任何发送操作都将导致panic异常。
对应已经被close的channel执行接收数据操作，依然可以接收到之前已经发送的数据；如果channel中已经没有数据，则产生这个channel元素类型的零值数据。
使用内置函数close可以关闭channel。
*/
func chan_close() {
	fmt.Println("chan_close")

	a := make(chan int, 5)
	defer close(a)
}

func recv_on_closed_chan() {
	fmt.Println("receive data on a close chan")

	a := make(chan int, 5)
	go func(c chan int) {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}(a)

	time.Sleep(6 * time.Second)

	for i := 0; i < 7; i++ {
		b := <-a
		fmt.Printf("received data: %d\n", b)
	}
}

/* 下面例子中，函数运行完成后，发送数据的goroutine也会被中断发送 */
func chan_send_abort() {
	fmt.Println("chan_send_abort")

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
	}
}

/*
无缓存channel (Unbuffered Channel)

一个不带缓存的channel(unbuffered channel)的发送操作将导致发送者goroutine**阻塞**，直到另一个goroutine在相同的channel上执行接收操作。
当发送的值在channel上成功传输后，两个goroutine可以继续执行后面的语句。
反之，如果接收操作先发生，那么接收者goroutine也将阻塞，知道有另一个goroutine在相同的channel上执行发送操作。

基于对无缓存channel的发送和和接收操作将导致两个goroutine做一次同步操作。
因为这个原因，无缓存的channel有时候也被称为同步channel (synchronous channels)。

当通过一个无缓存channel发送数据时，接收者收到数据发生在唤醒发送goroutine之前（happens before，这是go语言并发内存模型的一个关键术语）。
在讨论并发性时，当我们说 X 发生在 Y 之前（happens before），我们不仅仅指 X 发生的时间早于 Y 这一种情况（也可能X比Y发生的时间晚），我们
要表达的意思是，保证在(执行)Y之前X的所有影响（例如对变量的更新）都已经完成了，我们可以放心的依赖它（已完成的X事件）。

当我们说x事件既不是发生在y之前，也不是发生在y之后，那么我们说x和y是并发的（x is concurrent with y）。这并不意味这x和y就是同一时间发生的，
我们只是不能确定这两个时间发生的先后顺序。
*/

/*
我们可以使用一个无缓存channel来同步两个goroutine。
基于channel发送消息有两个重要方面。首先每一个消息都有一个值，但有时候通信的事实和发生的时刻也同样重要。当我们更希望强调通信发生的时刻时，
我们称它为**消息事件**。有些消息事件并不额外携带信息，它仅仅是用作两个goroutines之间的同步，这个时候我们可以用struct{}空结构体来作为
channel的元素类型。虽然也可以使用布尔或整型实现同样的功能，`done <- 1` 语句也比 `done <- struct{}{}` 更短。
下面的例子中，我们使用空结构体struct{}作为channel的元素类型。对于传递的值，我们用struct{}{}，因为struct{}{}不占用内存空间。
*/
func sync_routines() {
	fmt.Println("sync_routines")

	done := make(chan struct{})
	go func(c chan struct{}, count int) {
		fmt.Println("cutting down...")
		for i := count; i >= 0; i-- {
			fmt.Printf("\r%d", i)
			time.Sleep(1 * time.Second)
		}
		c <- struct{}{}
	}(done, 5)

	<-done
	fmt.Printf("\rsub routine is done\n")
}

/*
Pipelines
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

	fmt.Printf("%s\n", strings.Repeat("-", 64))
	sync_routines()

	fmt.Printf("%s\n", strings.Repeat("=", 64))
}
