package main

/*
goroutine和线程（goroutines and threads）

每个OS线程都有一个固定大小的栈内存（通常为2MB），栈内存区域用于保存在其它函数调用期间那些正在执行或临时暂停的函数中局部变量。
这个固定的栈大小即太大又太小。对于一个小的goroutine，2MB的栈是一个巨大的浪费，比如有的goroutine仅仅等待一个WaitGroup
再关闭一个channel，对于这种情况，栈就太大了。另外，对于更复杂和深度递归的函数，固定大小的栈始终不够大。

最为对比，一个goroutine在生命周期开始时只有一个很小的栈，典型情况为2KB。与OS线程类似，goroutine的栈也是用来存放那些正在
执行或临时暂停的函数中的局部变量。但与OS线程不同的是，goroutine的栈大小是不固定的，它可以按需增大或缩小。
goroutine的栈大小限制可以达到1GB，比典型的固定大小的栈高几个数量级。当然，只有极少数的goroutine会使用如此大的栈。

*/

/*
OS对线程的调度代价较大，相比OS对线程的调用模型，Go语言运行时包含一个自己的调度器，这个调度器使用一个m:n调度的技术。
（关于线程模型，请参考：https://mp.weixin.qq.com/s/jhOSjVyRA6rNKqVT2pKMIQ）

Go调度只关心单个Go程序的goroutine调度问题。
Go调度只关心单个Go程序的goroutine调度问题。
Go调度只关心单个Go程序的goroutine调度问题。

与操作系统的线程调度模型不同的是，Go调度器不是由硬件时钟来定期触发的，而是由特定的Go语言结构来触发的。
比如当一个goroutine调用time.Sleep或被通道阻塞或对互斥量操作时，调度器会将这个goroutine设置为休眠模式，
并允许其它goroutine知道前一个可重新唤醒为止。因为它不需要切换到内核环境，所以调用一个goroutine比调用一个线程成本低很多。
*/

/*
Go调度器使用GOMAXPROCS参数来确定使用多少个操作系统线程同时来执行Go代码。
默认是机器上的CPU数量，所以在一个8核的机器上，调度器会把Go代码同时调度到8个OS线程上。（GOMAXPROCS是m:n调度中的n）
正在休眠或正被通信channel堵塞的goroutine不需要占用线程。
阻塞在I/O或其它系统调用中或调用非Go语言写的函数的goroutine需要一个独立的OS线程，但这个线程不计算在GOMAXPROCS内。
*/

/*
goroutine没有标识

在大部支持多线程的操作系统和编程语言里，当前线程都有一个独特的标识，它通常可以取一个整数或指针。这个特性可以让我们轻松的构建
一个线程的局部存储，它本质是就是一个全家的map，以线程的标识作为键，这样每个的线程就可以独立地使用这个map存储和获取值，而不受
其它线程干扰。

goroutine没有可供程序员访问的标识。这个是由设计来决定的，因为线程的局部存储又被滥用的倾向。
Go语言鼓励一种更简单的编程风格，其中，能影响一个函数行为的参数应当是显示指定的。这不仅让程序更易阅读，还让我们能自由的把一个函数
的子任务分发到多个不同的goroutine而不用担心这些goroutine的标识。
*/
