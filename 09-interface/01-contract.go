package main

import (
	"bufio"
	"bytes"
	"fmt"
)

/*
接口即约定

接口是对其它类型行为的抽象和概括；因为接口类型不会和特定的实现细节绑定在一起，通过这种抽象的方式可以让我们的函数更加的灵活和具有适应能力。
Go语言中接口类型的独特之处在于它是满足隐式实现的（satisfied implicitly）。
也就是说，我们没有必要对给定的具体类型定义所有满足的接口类型；接口简单的拥有一些必要的方法就足够了（即一个具体的类型可以转化为多个接口而不用显式声明它）。
这种设计可以让我们创建一个新的接口类型满足已经存在的具体类型却不会改变这些类型的定义。当我们使用的类型来自于不受我们控制的包时这种设计尤为有用。
一个类型可以自由的使用另一个满足相同接口的类型来进行替换被称作可替换性(LSP里氏替换)。这是一个面向对象的特征。
接口类型是一种抽象的类型，它由一系列方法组成，不包含其它内容。

*/

// 接口的定义：type + 标识符 + interface + { 方法集 }
type ReadWriter interface {
	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
}

type WordCount int

func (w *WordCount) Write(p []byte) (n int, err error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*w++
	}
	return int(*w), nil
}

func main() {
	input := `
	add a module:
	ng g ng-alain:module test
	`
	var w WordCount
	w.Write([]byte(input))
	fmt.Println(w)

	w = 0
	fmt.Fprintf(&w, "hello, world")
	fmt.Println(w)
}
