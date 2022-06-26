# 编译原理相关

## SSA中间代码生成 
生成不同架构下的SSA代码。代码从源码到SSA代码会经历许多步骤：
1. 词法分析。将源代码解析为编译器能“看懂”的结构。
2. 语法分析A。把词法分析的结果作为输入，根据LALR(1)文法生成抽象语法树AST。
3. 对语法树中的各节点进行类型检查。  
   遍历语法树中的各节点，检查每个接点的类型，找出可能存在的语法错误。  
   在此过程中也可能会对语法树进行改写，这不仅能够去除一些不被执行的代码，优化代码已提高执行效率，而且会修改make、new等关键字对应节点的**操作类型**（这些操作类型会对应到运行时包中的**具体函数类型**，然后在下一步生成代码环节会替换为实际的函数。如make对应的实际类型可能是 OMAKECHAN/OMAKESLICE/OMAKEMAP）；
4. 为语法树生成中间代码。  
   中间代码的生成过程是抽象语法树到SSA中间代码的转换过程，在此期间会再次改写语法树中的关键字（注：前一步改写的是关键字节点的操作类型，这一步会根据类型将关键字改写为运行时包中用到的真实函数，如OMAKECHAN类型的节点会将make函数改为makechan/makechan64，panic会被改写为gopanic等）。go语言中的很多关键字和内置函数都是在该阶段被转换成运行时包中的方法。  
   改写后的语法树会经过多轮处理（目前～50步优化等操作）生成最终SSA中间代码。

下面是生成对应指令架构集的SSA命令：
```shell
GOSSAFUNC=hello CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build hello.go
GOSSAFUNC=hello CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build hello.go
```

## 机器码生成

机器码的生成，实际上就是对SSA中间代码的**降级**（lower）过程。在SSA中间代码降级的过程中，编译器将一些值重写成目标CPU架构的特定值。
### 1. SSA降级
SSA降级是在中间代码生成过程中完成的。在近50轮的处理过程中，lower以及后面的阶段都属于SSA降级这一过程。
### 2. 汇编器
汇编器是将汇编语言翻译成机器语言的程序，Go语言的汇编器是基于 Plan 9 汇编器的输入类型设计的。

代码生成汇编的命令：
```shell
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go tool compile -N -l -S -o hello-amd64.o hello.go
go tool objdump hello-amd64.o 
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go tool compile -N -l -S -o hello-arm64.o hello.go
go tool objdump hello-arm64.o 
```
