package main

/*
注意事项

反射是一个功能和表达能力都很强大的工具，但应该谨慎使用它，具体有三个原因：
1. 基于反射的代码时脆弱的。能导致编译器报告类型错误的每种写法，在反射中都有一个对应的误用方法。
编译器在编译时就能报告这个错误，而反射错误需要等到运行时才会以崩溃的方式告知。而这可能是代码写了
很久以后，甚至时代码开始执行很久以后才会发生的事情。

2. 类型其实也算某种形式的文档，而反射的相关操作则无法做到静态类型检查，所以大量使用反射的代码是很难理解的。
对于接受interface{}或者reflect.Value的函数，一定要写清楚期望的参数类型和限制条件（不变量）。

3. 基于反射的函数比特定类型优化的函数满一两个数量级。在一个典型的程序中，大部分函数与整体性能无关，所以为了
让程序更清晰可以使用反射。测试就很适合使用反射，因为大部分测试都使用小数据集。但对于关键路径上的函数，
则最好避免使用反射。
*/