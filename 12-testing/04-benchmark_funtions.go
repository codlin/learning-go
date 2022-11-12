package main

/*
Benchmark Functions

基准测试就是在一定的工作负载之下检测程序性能的一种方法。
在Go里面，基准测试函数看起来像一个测试函数，但是前缀是Benchmark并且拥有一个*testing.B参数来大多数提供和*testing.T相同的方法，
额外增加了一些和性能检测相关的方法。它还提供了一个整型成员，用来指定被检测操作的执行次数。
以word3目录中BenchmarkIsPalindrome中为例，和测试不同，默认情况下不会运行任何基准测试。
标记-bench的参数指定了要运行的基准测试。它是一个匹配Benchmark函数名称的正则表达式，它的默认值不匹配任何函数。
模式"."使它匹配包中所有的基准测试函数。因为这里只有一个基准测试函数，所以和指定-bench=BenchmarkIsPalindrome效果一样。
-----------------------------------------------------------------------
$ go test -bench ^BenchmarkIsPalindrome$
goos: darwin
goarch: amd64
pkg: word
cpu: Intel(R) Core(TM) i5-4278U CPU @ 2.60GHz
BenchmarkIsPalindrome-4          3301933               354.9 ns/op
PASS
ok      word    1.671s
-----------------------------------------------------------------------
基准测试名称的数字后缀表示GOMAXPROCS的值（此处是4），这个对并发基准测试很重要。
报告告诉我们每次IsPalindrome调用花费的时间是354.9ns，这个是3301933此调用的平均值。
因为基准测试运行器开始的时候并不清楚这个操作的耗时长短，所以开始的时候它使用了比较小的N值来做检
测，然后为了检测稳定的运行时间，推断出足够大的N值。

使用基准测试函数来实现循环而不是在测试驱动程序中调用代码的原因是，在基准测试函数中在循环外面可以执行一些必要的初
始化代码并且这段时间不加到每次迭代的时间中。如果初始化代码干扰了结果，参数testing.B提供了方法用来停止、恢复和重置计时
器，但是这些方法很少用到。
我们对性能优化，增加sPalindrome2和BenchmarkIsPalindrome2，下面是优化后的结果：
-----------------------------------------------------------------------
$ go test -bench=IsPalindrome
goos: darwin
goarch: amd64
pkg: word
cpu: Intel(R) Core(TM) i5-4278U CPU @ 2.60GHz
BenchmarkIsPalindrome-4          3232236               361.2 ns/op
BenchmarkIsPalindrome2-4         3265749               346.9 ns/op
PASS
ok      word    3.191s
-----------------------------------------------------------------------
性能有轻微的提升。
另外一个主意是为letters预分配一个容量足够大的数组，而不是
通过连续的append调用来追加。像这样声明一个合适长度的数组
letters：
func IsPalindrome3(s string) bool {
	letters := make([]rune, 0, len(s))
	//...
}
这次性能提升不少，输出如下：
-----------------------------------------------------------------------
$ go test -bench=IsPalindrome
goos: darwin
goarch: amd64
pkg: word
cpu: Intel(R) Core(TM) i5-4278U CPU @ 2.60GHz
BenchmarkIsPalindrome-4          3123691               472.1 ns/op
BenchmarkIsPalindrome2-4         3529777               342.9 ns/op
BenchmarkIsPalindrome3-4         3236635               322.1 ns/op
PASS
ok      word    6.090s
-----------------------------------------------------------------------
如上面的例子，最快的程序通常是那些进行内存分配次数最少的程序。
命令行标记-benchmem在报告中包含了内存分配统计数据。这里和优化之前的内存分配进行比较：
-----------------------------------------------------------------------
$ go test -bench=IsPalindrome -benchmem
goos: darwin
goarch: amd64
pkg: word
cpu: Intel(R) Core(TM) i5-4278U CPU @ 2.60GHz
BenchmarkIsPalindrome-4          2954821               773.3 ns/op           248 B/op          5 allocs/op
BenchmarkIsPalindrome2-4         3080832               349.5 ns/op           248 B/op          5 allocs/op
BenchmarkIsPalindrome3-4         6177043               177.0 ns/op           128 B/op          1 allocs/op
PASS
ok      word    9.702s
-----------------------------------------------------------------------
这种基准测试告诉我们给定操作的绝对耗时，但是在很多情况下，引起关注的性能问题是两个不同操作之间的相对耗时。
例如，如果一个函数需要1ms来处理1000个元素，那么它处理10000个或者100万个元素需要多久呢？
另外一个例子：I/O缓冲区的最佳大小是多少。对一个应用使用一系列的大小进行基准测试可以帮助我们选择最小的
缓冲区并带来最佳的性能表现。第三个例子：对于一个任务来讲，哪种算法表现最佳？对两个不同的算法使用相同的输入，
在重要的或者具有代表性的工作负载下，进行基准测试通常可以显示出每个算法的优点和缺点。

性能比较函数只是普通的代码。它们的表现形式通常是带有一个参数的函数，被多个不同的Benchmark函数传入不同的值来调
用，如下所示：
func benchmark(b *testing.B, size int) {  ...  }
func Benchmark10(b *testing.B) { benchmark(b, 10) }
func Benchmark100(b *testing.B) { benchmark(b, 100) }
func Benchmark1000(b *testing.B) { benchmark(b, 1000) }
参数size指定了输入的大小，每个Benchmark函数传入的值都不同但是在每个函数内部是一个常量。
不要使用b.N作为输入的大小。除非把它当作固定大小输入的循环次数，否则该基准测试的结果毫无意义。

基准测试比较揭示的模式在程序设计阶段很有用处，但是即使程序正常工作了，我们也不会丢掉基准测试。
随着的程序演变，或者它的输入增长了，或者它被部署在其他的操作系统上并拥有一些新特性，我们仍然
可以重用基准测试来回顾当初的设计决策。

*/
