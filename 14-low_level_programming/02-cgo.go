package main

/*
CGO

一个Go程序或许需要调用用C实现的硬件驱动程序，查询一个
用C++实现的嵌入式数据库，或者使用一些以Fortran实现的线性代
数协程。C作为一种编程混合语言已经很久了，所以无论那些广泛使
用的包是哪种语言实现的，它们都导出了和C兼容的API。

如果C库很小，我们可以使用纯Go语言来移植它，并且如果性
能对我们来说不是很关键，我们最好使用包os/exec以辅助子进程的
方式来调用C程序。
仅当你需要使用拥有有限C API并且复杂的、性能关键的库时，
使用cgo来把它们包装成Go语言的绑定才有意义。

*/
