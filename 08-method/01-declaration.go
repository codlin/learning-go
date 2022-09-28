package main

import (
	"fmt"
	"math"
)

/*
方法的声明

在函数声明时，在其名字前面放上一个变量，即是一个方法。
这个附加的参数会将该函数附加到这种类型上，即相当于为这中类型定义了一个独占的方法。
*/
type Point struct {
	X, Y float64
}

// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

/*
上面的代码里附加的参数p，叫做方法的接收器（receiver）。
不同于其它语言用self或this作为接收器，在Go语言里，可以任意选择接收器的名字。
建议使用比较短的名字，通常使用类型的第一个字母。

在调用时，都是接受名字在方法之前，用点好分开。
*/
func test_method() {
	p := Point{1, 2}
	q := Point{4, 6}

	/* 这种p.Distance的表达式叫做选择器，因为它会选择合适的对应p这个对象的Distance方法来执行
	选择器也可以用来选择一个struct结构类型的字段，比如p.X。由于方法和字段都在同一个命名空间，所以如果有两个名字一样的方法和变量时会报错。
	*/
	fmt.Println(p.Distance(q))

	/* 调用包级别的Distance */
	Distance(p, q)
}

// A Path is a journey connecting the points with straight lines.
type Path []Point

/* 同样的方法名字，因为类型的不同，所以具有不同的命名空间，因此不会发生冲突 */
// Distance returns the distance traveled along the path.
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}
func test_distance() {
	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance()) // "12"
}

func main() {
	test_method()
	test_distance()
}
