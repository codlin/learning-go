package main

import "strconv"

/*
反射（Reflection）

Go语言提供了一种机制，在编译时不知道类型的情况下，可以来更新变量，并在运行时查看它们的值、
可以调用它们方法，以及直接对它们的布局进行操作，这种机制称之为反射。
反射也让我们把类型本身当作头等值来看待（first-class values）。

为什么需要反射？
有时候我们要写一个能够统一处理各种数据类型的函数，而这些类型可能无法共享同一个接口，也可能
它们的类型描述（representation）为止，也有可能这个类型在我们设计函数时还不存在，甚至这些
类型会同时存在上面三种问题。
一个熟悉的例子就是fmt.Printf中格式化的逻辑，它可以输出任意类型的任意值，甚至是用户自定义的
类型。让我们先通过学到的知识来实现一个类似的函数。
*/
func Sprint(x interface{}) string {
	type stringer interface {
		String() string
	}
	switch x := x.(type) {
	case stringer:
		return x.String()
	case string:
		return x
	case int:
		return strconv.Itoa(x)
	// 对其它整型或浮点型做类似处理
	case bool:
		if x {
			return "true"
		}
		return "false"
	default:
		// array, slice,chan, struct,map, pointer...
		return "???"
	}
}

/*
上面的例子虽然我们处理了一些基本类型，但是还有更多的类型我们无法处理，当我们无法透视一个未知类型的描述时，
这段代码就无法继续用了，这是我们就需要反射救场了。
*/
