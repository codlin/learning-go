package main

/*
覆盖率

从本质上来看，测试从来不会结束。
著名计算机科学家Edsger Dijkstra说，“测试旨在发现bug，而不是证明其不存在”。

一个测试套件覆盖待测试包的比例称之为测试的覆盖率。
覆盖率无法通过数据来衡量，任何事情都是动态的，即使最微小的程序都无法精确的测量。
但还是有办法帮助我们将测试精力放在最有潜力的地方。
语句覆盖率是一种最简单且广泛使用的方法之一。一个测试套件的覆盖率指的是部分语句在
一次执行中至少执行一次。
Go的cover工具被集成到了go test中，用来衡量语句覆盖率并帮助识别测试间的明显差别。
`go tool cover`命令输出了覆盖工具的使用方法：
------------------------------------------------------------------------
$ go tool cover
Usage of 'go tool cover':
Given a coverage profile produced by 'go test':
        go test -coverprofile=c.out

Open a web browser displaying annotated source code:
        go tool cover -html=c.out

Write out an HTML file instead of launching a web browser:
        go tool cover -html=c.out -o coverage.html

Display coverage percentages to stdout for each function:
        go tool cover -func=c.out

Finally, to generate modified source code with coverage annotations
(what go test -cover does):
        go tool cover -mode=set -var=CoverageVariableName program.go
------------------------------------------------------------------------
以word目录为例执行覆盖率测试，进入word目录，带上-coverprofile标记来运行测试：
$ go test -coverprofile=c.out
这个标记通过检测产品代码，启用了覆盖数据收集。也就是说，它修改了源码的副本，这样在每个语句块
执行之前，设置一个布尔变量，每个语句块都对应一个变量。在修改的程序退出之前，它将每个变量的值
都写入到指定的c.out日志文件中并且输出被执行语句的汇总信息。如果只需要汇总信息，那么可以使用
`go test -cover`
如果go test命令指定了-covermode=count，每个语句块的检测会递增一个计数器而不是设置布尔变量。
关于每个语句块执行次数的日志使得量化成为可能，可由此识别出执行率比较高的“热块”或者相反的“冷块”。

在生成报告后，我们使用cover工具来处理生成的日志，生成一个html报告，并在一个新的浏览器窗口打开它。
$ go tool cover -html=c.out
是否使用`-covermode=count`，生成的html报告会有不同，使用这个标记的会一个low cover到high cover
的颜色等级，等级越高越绿，代表语句被执行的次数越多。

实现语句的100%覆盖听上去很宏伟，但是实际情况下这并不可行，也不会行之有效。
因为语句得以执行并不意味着这是没有bug的，拥有复杂表达式的语句块必须使用不同的输入执行多次来覆盖用例。
有一些语句（譬如panic）可能永远也不会被执行，其它的也很难检测并且实际上也很少被执行。
测试基本上是实用主义行为，在编写测试的代价和本可以本可以通过测试避免的错误之间平衡。
覆盖工具可以帮助识别最薄弱的点，但是和编程一样，设计好的测试用例通常需要一丝不苟的精神。
*/
