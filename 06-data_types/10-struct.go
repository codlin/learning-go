package main

import (
	"fmt"
	"time"
)

/*
struct

结构体是一种聚合的数据结构，是由0个或多个任意类型的值聚合成的实体。每个值称为结构体的成员。
*/

// 下面声明了一个叫 Employee 的命名结构体类型
type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
	Link      *Employee
}

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

func main() {
	struct_var()
	struct_as_func_param()
}
