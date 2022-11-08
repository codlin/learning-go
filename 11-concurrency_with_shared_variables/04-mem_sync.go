package main

import (
	"fmt"
	"sync"
)

/*
内存同步

对于`03-rwmutex.go`中的Balance函数加了读锁，这可能使人感到奇怪。
因为Balance只包含一个单个操作，它不存在另外一个goroutine插在中间执行的风险。
其实需要互斥锁的原因有两个：
1. 防止Banlance插到其它操作中间也是很重要的， 比如Withdraw；
2. 第二个原因更微妙，因为同步不仅涉及到多个goroutine的执行顺序问题，同步还会影响到内存。

现代的计算机一般都有多个处理器，每个处理器都有内存的本地缓存。
为了提高效率，对内存的写入是缓存在每个处理器中的，只有在必要时才刷回内存。
甚至刷回内存的顺序都有可能与goroutine当时的写入顺序不一致。
像channel通信或者mutex操作这样的同步原语会导致处理器刷新并提交所有累积的写入，以便保证
goroutine执行到该点的效果对运行在其它处理器上的goroutine可见。
*/

// 考虑如下代码可能的输出
func concurrency() {
	loop := 10000

	for i := 0; i < loop; i++ {
		var x, y int
		var wg sync.WaitGroup

		wg.Add(2)
		go func() {
			defer wg.Done()
			x = 1                   // A1
			fmt.Print("y:", y, " ") // A2
		}()

		go func() {
			defer wg.Done()
			y = 1                   // B1
			fmt.Print("x:", x, " ") // B2
		}()
		wg.Wait()

		fmt.Println()
	}
}

/*
上面的代码不但有我们认为的以下输出：
x: 0, y: 1
y: 0, x: 1
x: 1, y: 1
y: 1, x: 1

还会有我们意料之外的输出：
x: 0, y: 0
y: 0, x: 0
x: 1, y: 0
y: 1, x: 0

在单个goroutine内，每个语句的效果保证按照执行的顺序发生，就是说，goroutine是串行一致的（sequentially consistent）。
但在缺乏通道和互斥量来显式同步的情况下，并不能保证所有的goroutine看到的事件顺序是一致的。
尽管groutine A肯定能在读取y之前观察到x=1的效果，但它并不一定能观察到goroutine对y写入的效果，所以A可能会输出y的一个过期值。

尽管很容易把并发理解为多个goroutine中语句的某种交错执行方式，但正如上面的例子显示的，`这并不是一个现代编译器和CPU的工作方式`。
因为赋值和Print应对不同的变量，所以编译器就可能认为两个语句的执行顺序不会影响到结果，然后就交换了这两个语句的执行顺序。
CPU也有类似的问题，如果两个goroutine在不同的CPU上执行，每个CPU都有自己的缓存，那么一个goroutine的写入操作在同步到内存之前
对另外一个goroutine的Print的语句是不可见的。

这些并发问题可以通过简单的、成熟的模式来避免，即在可能的情况下，把变量限制在单个goroutine中，对于其它变量，使用互斥锁。
*/

func main() {
	concurrency()
}
