package main

import (
	"fmt"
	"image/color"
	"math"
)

/*
Composing Types by Struct Embedding

通过在结构体中嵌入其它结构体来扩充当前结构体的方法

*/

type Point struct {
	X, Y float64
}

// method of the Point type
func (p *Point) Distance(q Point) float64 {
	fmt.Println("Point.Distance")
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// method of the Point type, use *Point as receiver
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

/*
我们完全可以将 ColoredPoint 定义为有3个字段的结构体，但我们将 Point 类型嵌入到 ColoredPoint 来提供X,Y字段。
在数据类型struct中我们知道，内嵌可以使我们在定义 ColoredPoint 时得到一种句法上的简写形式。其中包括 Point 中的一切字段和方法。
但是，如果是嵌入别的包中结构体，那么只能访问它导出的字段和方法。
*/
type ColoredPoint struct {
	Point
	Color color.RGBA
}

func test_embed_struct_field() {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X)

	cp.Y = 2
	fmt.Println(cp.Y)
}

/*
我们可以把ColoredPoint类型当作接收器，来调用嵌入的Point类型的方法，即使ColoredPoint中没有声明这些方法。
用这种方式，内嵌机制可以使我们定义字段特别多的复杂类型，我们可以将字段先按小类型分组，然后定义小类型的方法，之后再将它们组合起来。
*/
func test_embed_struct_method() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	p := ColoredPoint{Point{1, 1}, red}
	q := ColoredPoint{Point{5, 4}, blue}
	fmt.Println(p.Distance(q.Point))

	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point))
}

/*
虽然可以直接调用嵌入类型导出的方法，但嵌入类类型Point并非是ColoredPoint的基类，也不能将ColoredPoint看作是'is-a' Point类型。
Point的Distance方法，因为ColoredPoint并非是Point类型，所以不能直接传入q
p.Distance(q) // compile error: cannot use q (ColoredPoint) as Point

一个ColoredPoint并不是一个Point，但他 'has a' Point
并且它有从Point类里引入的Distance和ScaleBy方法。
如果你喜欢从实现的角度来考虑问题，内嵌字段会指导编译器去生成额外的包装方法来委托已经声明好的方法，
Distance和下面的形式是等价的：
*/
func (p *ColoredPoint) Distance(q Point) float64 {
	fmt.Println("ColoredPoint.Distance")
	return p.Point.Distance(q)
}

/*
当编译器解析一个选择器到方法时，比如p.Distance，
它会首先去找直接定义在ColoredPoint这个类型里的Distance方法，
然后找被ColoredPoint的内嵌字段们引入的方法，
然后去找Point和RGBA的内嵌字段引入的方法，
然后一直递归向下找。
如果选择器有二义性的话编译器会报错，比如你在同一级里有两个同名的方法。
*/

func main() {
	test_embed_struct_field()
	test_embed_struct_method()
}
