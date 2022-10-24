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
func close_chan() {
	fmt.Println("close_chan")

	a := make(chan int, 5)
	close(a)
	// a <- 1 // panic: send on closed channel
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

Channels也可以用于将多个goroutine串联在一起，一起channel的输出作为下一个channel的输入。这样串联起来的channel就是所谓的管道（pipeline）。
*/

// 这个程序是一个简单的管道，发送自然数给channel，下一个goroutine接收数据后计算乘方，然后发送给下一个channel，最后主goroutine打印出最终值。
func channels_pipeline_1() {
	fmt.Println("channels_pipeline_1")

	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	// Squares
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	// Printers (in main goroutine)
	for {
		fmt.Println(<-squares)
	}
}

/*
上面的程序 channels_pipeline_1 将生成0、1、2、4、9、...形式的无穷数列。
像这样的串联管道可以用在需要长时间运行的服务中，每个长时间运行的goroutine可能会包含一个死循环，在不同goroutine的死循环内部使用串联的channels来通信。

但是，如果我们希望发送有限的数量该如何处理？
如果发送者知道没有更多的值发送到channel的话，那么让接收者也能及时知道没有多余的值可接收将是有用的，因为接收者可以停止不必要的等待。
可以通过内置的close来关闭通道：close(naturals)。
在关闭的channel上发送数据将会导致panic。在关闭的channel上接收数据将会得到该channel元素类型的零值。
*/

// 这个程序在发送10个数后关闭channel naturals。因为无法检测到channel关闭，所以后面for循环一定次数后退出
func channels_pipeline_2() {
	fmt.Println("channels_pipeline_2")

	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; ; x++ {
			naturals <- x
			if x == 10 {
				close(naturals)
				break
			}
		}
	}()

	// Squares
	go func() {
		for x := 0; x < 15; x++ {
			x := <-naturals
			squares <- x * x
		}
		close(squares)
	}()

	// Printers (in main goroutine)
	for x := 0; x < 20; x++ {
		fmt.Println(<-squares)
	}
}

/*
没有办法直接测试一个channel是否关闭，但是接收操作有一个变体形式：
它过接收一个结果，多接收的第二个结果是一个布尔值，true表示成功从channels接收到值，flase表示channels已经被关闭并且里面没有值看接收。
基于这个特性，下面是修改后的代码：
*/
func channels_pipeline_3() {
	fmt.Println("channels_pipeline_3")

	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 10; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// Squares
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break
			}
			squares <- x * x
		}
		close(squares)
	}()

	// Printers (in main goroutine)
	for {
		x, ok := <-squares
		if !ok {
			break
		}
		fmt.Println(x)
	}
}

/*
因为上面 channels_pipeline_3 检查channel关闭的语法比较笨拙，而且这种模式很常用，所以Go语言的range循环支持可以直接在channel上迭代。
*/
func channels_pipeline_4() {
	fmt.Println("channels_pipeline_4")

	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 10; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// Squares
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// Printers (in main goroutine)
	for x := range squares {
		fmt.Println(x)
	}
}

/*
其实不需要关闭每一个channel。只有当需要告诉接收值goroutine，所有的数据已经全部发送完毕时才需要关闭channel。
不管一个channel是否被关闭，当它没有被引用时将会被Go语言的垃圾回自动回收器回收。
试图重复关闭一个channel将导致panic异常，试图关闭一个nil的channel也将导致panic异常。
关闭一个channel还会触发一个广播机制。
*/

/*
单方向的channel (Unidirectional Channel Types)

随着程序的增长，人们习惯将大函数拆分成小函数。然后将channel作为参数传递进去。如func squarer(out, in chan int)
我们期望一个用于输出，一个用于输入，参数的名字也表达了这种意图。但这无法保证这些channel实际上是按照这种意图被使用的。
为了表明这种意图并防止被滥用，Go语言的类型系统提供了单项的channel类型，分别用于只发送或只接收的channel。
类型chan<- int表示一个只发送int的channel，只能发送不能接收。相反，类型<- chan int表示一个接收int的channel，只能接收不能发送。
这种限制将在编译器检测。
因为关闭操作只用于断言不在向channel发送新的数据，所以只有在发送者所在的goroutine才会调用close函数，
因此对一个只接收的channel调用close将是一个编译错误。
*/
func counter(out chan<- int) {
	for x := 0; x < 10; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func channels_pipeline_5() {
	fmt.Println("channels_pipeline_5")

	/* 注意，创建的channel还是双向的 */
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	/*
		调用counter(naturals)将导致将chan int 类型的naturals隐式地转换为chan<‐ int 类型的只发送型的
		channel。调用printer(squares)也会导致相似的隐式转换，这一次是转换为<‐chan int 类型只接收
		型的channel。任何双向channel向单向channel变量的赋值操作都将导致该隐式转换。这里并没有
		反向转换的语法：也就是不能一个将类似chan<‐ int 类型的单向型的channel转换为chan int 类型
		的双向型的channel。
	*/
	go counter(naturals)

	// Squares
	go squarer(squares, naturals)

	printer(squares)
}

/*
带缓存的channel （Buffered Channel）

带缓存的channel内部持有一个元素队列。队列的最大容量是在调用make函数创建时通过第二个参数指定的。
下面的语句创建一个可持有3个字符串元素的带缓存的channel。
ch = make(chan string, 3)
+----+         +----+----+----+
| *--|-------->|    |    |    |
+----+         +----+----+----+
向channel的发送操作就是向内部缓存队列的尾部插入元素，接收操作则是从队列的头部删除元素。
如果内部队列是满的，那么发送操作将阻塞一直到另一个goroutine执行接收操作而释放了新的队列空间。
相反，如果channel是空的，那么接收操作将阻塞直到另一个goroutine执行发送操作向队列插入元素。
可以使用内置函数cap()来获取channel的容量，使用len()来获取队列中有效元素的个数。
*/
func buffered_channel() {
	fmt.Println("buffered_channel")

	ch := make(chan string, 3)
	ch <- "A"
	ch <- "b"
	fmt.Println(len(ch))
	fmt.Println(cap(ch))
}

/*
Go语言新手有时候会将一个带缓存的channel当作同一个goroutine中的队列使用，虽然语法简单，但实际上是一个错误。
Channel和goroutine的调度机制是紧密相连的，如果没有goroutine从channel接收，发送方————也可能是整个程序————可能会永远阻塞。
如果你只是需要一个简单的队列，使用slice就可以了。

如果缓存因为没有goroutine接收而卡住，这种情况下称之为goroutines泄漏，这将是一个BUG。和垃圾变量不同，泄露的goroutines并不会被自动回收，
因此确保每个不再需要的goroutine能正常退出是重要的。

关于无缓存或带缓存channels之间的选择，或者是带缓存channels的容量大小选择，都可能影响程序的正确性。
无缓存channel更强地保证了每个发送操作与相应的同步接收操作；但是对于带缓存的channel，这些操作都是解耦的。
同时，即使我们知道将要发送到一个channel的信息数量的上限，创建一个对应容量大小的channel并在接收之前发送
所有的值，这种做法是不罕见的。如果未能分配足够的缓冲将导致程序锁死。
channel的缓存也可能影响程序的性能（类似于流水线上各个环节不同的处理速度）
*/

/*
Looping in Parallel
有点杂乱，待后补充
*/

func main() {
	fmt.Printf("%s\n", strings.Repeat("=", 64))
	create_chan()

	fmt.Printf("%s\n", strings.Repeat("-", 64))
	close_chan()

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

	fmt.Printf("%s\n", strings.Repeat("-", 64))
	channels_pipeline_2()

	fmt.Printf("%s\n", strings.Repeat("-", 64))
	channels_pipeline_3()

	fmt.Printf("%s\n", strings.Repeat("-", 64))
	channels_pipeline_4()

	fmt.Printf("%s\n", strings.Repeat("-", 64))
	channels_pipeline_5()

	fmt.Printf("%s\n", strings.Repeat("=", 64))
}
