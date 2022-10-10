package main

import (
	"fmt"
	"math"
)

/*
方法值和方法表达式

我们通常选择一个方法，并在同一个表达式里面执行，如p.Distance(q.Point)，实际上将其分成两步执行也是可以的。
*/

type Point struct {
	X, Y float64
}

// method of the Point type
func (p *Point) Distance(q Point) float64 {
	// fmt.Println("Point.Distance")
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// method of the Point type, use *Point as receiver
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

/*
p.Distance这种形式称之为选择器，选择器会返回（yield）一个方法值 ———— 一个将方法Point.Distance绑定到特定接收器变量的函数。
这个函数可以不通过指定特定接收器即可调用，只要传入参数即可。
*/
func method_value_bind_receiver() {
	p := Point{1, 2}
	q := Point{4, 6}

	distanceFromP := p.Distance // 方法值 method value
	fmt.Println("method value was bind to a specific receiver:", distanceFromP(q))
}

/*
和方法值相关的还有方法表达式。当调用一个方法时，与调用一个普通的函数相比，我们必须要用选择器语法来指定方法接收器。
当T是类型时，方法表达式可以时T.f或(*T).f，会返回一个函数值，这种函数会将第一个参数作为接收器，可以用普通函数的方式调用。
如果方法的接收器是值类型，那么使用方法表达式时需要使用T.f的形式，对返回的函数值（非函数的返回值）调用时，函数的第一个参数需要传入值类型的值；
如果方法的接收器是指针类型，那么转换为方法表达式时需要(*T).f的形式，在函数调用时，第一个参数需要传入指针类型的值。
*/
func method_expression() {
	p := Point{1, 2}
	q := Point{4, 6}

	distance := (*Point).Distance // 方法表达式 method expression
	fmt.Println("method expression first param as receiver", distance(&p, q))
}

/*
当你根据一个变量来决定调用同一类型的哪个函数时，方法表达式就显得很有用了。你可以根据选择来调用接收器各不相同的方法。
*/
func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }
func (p Point) Sub(q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }

type Path []Point

/* 此处要注意，虽然接收器是值类型，每次调用时会复制一个新的局部变量，但是因为Path本身是切片slice，所以底层数组是共享的 */
func (path Path) TranslateBy(offset Point, add bool) {
	fmt.Printf("in path: 0x%p\n", &path)

	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		// Call either path[i].Add(offset) or path[i].Sub(offset).
		(path)[i] = op((path)[i], offset)
	}
}

func select_method_expression() {
	path := Path{Point{1, 2}, Point{4, 6}}
	fmt.Printf("out path: 0x%p\n", &path)

	for i := 0; i < len(path); i++ {
		add := false
		if i&1 == 0 {
			add = true
		}
		path.TranslateBy(Point{float64(i + 1), float64(i + 1)}, add)
		fmt.Println(path)
	}
}

func main() {
	method_value_bind_receiver()
	method_expression()
	select_method_expression()
}
