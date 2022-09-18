package main

import (
	"fmt"
)

/*
struct

结构体是一种聚合的数据结构，是由0个或多个任意类型的值聚合成的实体。每个值称为结构体的成员。

结构体成员的输入顺序也很重要，如果我们交换其中的成员顺序或者合并相同类型的成员，那么我们相当于定义了不同类型的结构体。
通常，我们只是将相关的成员写在一起。

如果结构体的名字是以大写字母开头的，那么这个结构体是可以导出的，包的外部可以导入并引用。
如果结构体的成员名字是以大写字母开头的，那么这个成员就是可导出的，这是Go语言规则决定的。
一个结构体可能同时包含导出和未导出的成员。

一个命名为S的结构体将不能再包含S类型的成员：因为一个聚合的值不能再包含它自身。
但它可以包含*S指针类型的成员，这样可以让我们创建递归的数据结构，比如链表和树结构。

结构体类型的零值是每个成员都是零值。通常会将零值作为最合理的默认值。
如果结构体没有任何成员的话就是空结构体，写作struct{}。它的大小是0，也不包含任何信息，但有的时候依然是有价值的。
*/

// 下面声明了一个叫 Employee 的命名结构体类型
type Employee struct {
	ID            int
	Name, Address string // 相同类型的可以合并到一行，但不推荐这样做
	Position      string
	Salary        int
	ManagerID     int
	Link          *Employee
}

/*
空结构体

大小为0，也不包含任何信息，但有时候依然是有价值的。譬如chan时传递信号。
有些Go程序员，在使用map模拟set时，用它来代替布尔类型的value，只是强调key的重要性，
但是因为节省的空间有限，同时语法比较复杂，因为应尽量避免这种情况。
*/
type EmptyStruct struct{}

/* 结构体变量的声明和使用 */
func struct_var() {
	var dilbert Employee

	/* 结构体变量的成员可以通过点操作符访问 */
	dilbert.Salary = 10000
	fmt.Printf("salary: %d\n", dilbert.Salary)

	/* 或者对成员取地址，然后通过指针访问 */
	salary := &dilbert.Salary
	*salary -= 3000
	fmt.Printf("salary: %d\n", dilbert.Salary) // 7000

	/* 点操作符也可以和指向结构体的指针一起工作 */
	alias := &dilbert
	alias.Position = "Senior" // 相当于 （*alias).Position = "Senior"
	(*alias).Position += " Engineer"
}

/*
结构体的字面值

结构体值也可以用结构体字面值表示，有两种形式的结构体字面值语法：
1. 以结构体成员定义的顺序为每个结构体成员指定一个值。需要全部成员都指定。
2. 以成员名字和对应着的值初始化，可以包含部分或全部成员（推荐）
两种不同的写法不能混合使用
*/
func struct_literal() {
	a := Employee{0, "Bob", "US", "Senoir", 10000, 0, nil}
	b := Employee{Name: "dilbert", Salary: 7000}
	fmt.Println(a, b)
}

func struct_as_func_param() {
	dilbert := Employee{Name: "dilbert", Salary: 10000}

	// 结构体作为函数参数传递是值传递，浅复制所有成员
	dilbert.Link = &Employee{Name: "Jonus", ManagerID: 1}
	fmt.Println(dilbert, dilbert.Link)
	struct_func_param(dilbert) // dilbert.Name 没有被改变，但是 dilbert.Link.ManagerID 被修改了
	fmt.Println(dilbert, dilbert.Link)
}

/* 结构体作为函数参数传递是值传递，浅复制所有成员 */
func struct_func_param(e Employee) {
	e.Name = "Sean"
	e.Link.ManagerID = 200
	fmt.Println(e)
}

func employee_by_name(name string) Employee {
	return Employee{Name: name}
}

func employee_by_id(id int) *Employee {
	return &Employee{ID: id}
}

/*
结构体作为函数返回值
https://zhuanlan.zhihu.com/p/423516595
*/
func struct_as_return() {
	fmt.Println(employee_by_name("dilbert").Name)
	/* 如果函数返回的是一个结构体类型的值，那么无法直接对其中的成员赋值，会报编译错误
	也就是说，在赋值语句左边是一个值而不是变量，无法对一个值取地址。就好比你们无法对字面常量取址一样
	*/
	// employee_by_name("Bob").Salary = 1000 // compile error, UnassignableOperand

	/* 同样，对返回值的成员取地址也会报错 */
	// a := &employee_by_name("Bob").Salary // compile error, UnassignableOperand

	/* 如果 */
	e := employee_by_id(1)
	fmt.Println(e)
	employee_by_id(2).Name = "Bob"

}

func main() {
	struct_var()
	struct_as_func_param()
}
