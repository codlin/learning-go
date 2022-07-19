package main

import (
	"fmt"
	"time"
)

/**
# 作为通道的信号传输
使用通道时候，有时候我们只关心是否有数据从通道内传输出来，而不关心数据内容，这时候通道数据相当于一个信号，比如我们实现退出时候。下面例子是基于通道实现的信号量。
这个例子很废柴，此处只做演示
**/
// empty struct
var emtpy = struct{}{}

// Semaphore is empty type chan
type Semaphore chan struct{}

// P used to acquire n resources
func (s Semaphore) P(n int) {
	for i := 0; i < n; i++ {
		s <- emtpy
	}
}

// V used to release n resouces
func (s Semaphore) V(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}

// Lock used to lock resource
func (s Semaphore) Lock() {
	fmt.Println("created one ...")
	s.V(1)
}

// Unlock used to unlock resource
func (s Semaphore) Unlock() {
	s.P(1)
	fmt.Println("released one ...")
}

// NewSemaphore return semaphore
func NewSemaphore(N int) Semaphore {
	return make(Semaphore, N)
}

// go run -gcflags "-m -l" function-chan.go
func main() {
	s := NewSemaphore(1)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("Lock %v\n", i)
			s.Lock()
			fmt.Printf("Locked %v\n", i)
			time.Sleep(2000)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("Unlock %v\n", i)
			s.Unlock()
			fmt.Printf("Unlocked %v\n", i)
			time.Sleep(4000)
		}
	}()

	for i:=0; i<10; i++ {
		time.Sleep(5000)
	}
}
