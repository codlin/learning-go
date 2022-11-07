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

func main() {
	concurrency()
}
