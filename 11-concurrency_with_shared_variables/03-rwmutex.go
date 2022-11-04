package main

import "sync"

/*
对于特殊的场景，如果一个共享变量函数只须读取变量的状态，所以多个请求可以安全的并发运行，但是写变量的请求没有同时运行即可。
在这种场景下，我们需要一种特殊类型的锁，它允许只读操作并发的执行，但写操作需要获取完全独享的访问权限。
这种锁称为多读单写锁。Go语言中的sync.RWMutex可以提供这种功能。RLock锁尽可用在临界区内共享变量无写操作的情景。
一般来讲，我们不应该假定那些逻辑上只读的函数和方法不会更新一些变量。如果有疑问，就应当使用独享版本的Lock。
仅在绝大部分goroutine都在获取读锁并且锁竞争比较激烈时，RWMutex才有优势。因为它需要更复杂的内部机制，
所以在竞争不激烈时，它比普通的互斥锁慢。
*/

// 还是银行的例子，使用加互斥锁的方式
var (
	mu      sync.RWMutex
	balance int
)

func Balance() int {
	mu.RLock()
	defer mu.RUnlock() // 使用defer延迟执行RUnlock

	return balance
}
