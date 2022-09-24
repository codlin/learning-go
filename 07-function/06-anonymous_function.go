package main

import (
	"fmt"
	"os"
	"sort"
)

/*
匿名函数

拥有函数名的函数只能在包级语法块中被声明，通过函数字面量，我们可以绕过这一限制，在任何表达式中表示一个函数值。
函数字面量的语法和函数声明类似，区别在于func后面没有函数名。
函数字面值是一种表达式，它的值被称为匿名函数。
*/
func function_literal() func(a, b int) int {
	return func(a, b int) int { return a + b }
}

/*
函数字面量允许我们在使用函数时，再定义它。
更为重要的是，通过这种方式定义的函数可以访问完整的词法环境（lexical  environment），这意味着在函数中定义的内部函数可以引用该函数的变量。
下面的例子证明函数值不仅仅是一串代码，还记录了状态。匿名函数引用了外部函数的状态，这也是函数值属于引用类型和函数值不可比较的原因。
Go使用闭包技术技术实现函数值，Go程序员也把函数值叫闭包。
*/
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func prerequisite() {
	var prereqs = map[string][]string{
		"algorithms": {"data structures"},
		"calculus":   {"linear algebra"},
		"compilers": {
			"data structures",
			"formal languages",
			"computer organization",
		},
		"data structures":       {"discrete math"},
		"databases":             {"data structures"},
		"discrete math":         {"intro to programming"},
		"formal languages":      {"discrete math"},
		"networks":              {"operating systems"},
		"operating systems":     {"data structures", "computer organization"},
		"programming languages": {"data structures", "computer organization"},
	}

	for i, course := range topo_sort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

/* 如果对匿名函数进行递归调用，需要先声明一个变量，然后再将匿名函数赋给这个变量。如果使用简短赋值的方式产生变量未定义的编译错误 */
func topo_sort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string) // 先声明变量，然后再和匿名函数绑定
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}

/*
Go语言词法作用域的一个陷阱：捕获迭代变量。
这是一个极易犯错的地方，即使经验丰富的程序员也会在这个问题上犯错。
*/
func rm_dirs_misfunc() {
	tmpDirs := []string{"/tmp/a", "/tmp/b", "/tmp/c"}
	var rmdirs []func()
	for _, dir := range tmpDirs {
		os.MkdirAll(dir, 0755)
		rmdirs = append(rmdirs, func() { os.RemoveAll(dir) })
	}

	// do some work...
	for _, rmdir := range rmdirs {
		rmdir() // clean up
	}
}

/*
上面函数的问题在于循环变量的作用域。在上面的程序中，for循环语句引入了新的词法块，循环变量dir在这个词法块中被声明。
在该循环中生成的所有函数值都共享相同的循环变量。需要注意的是，函数值中记录的是循环变量的内存地址，而不是循环变量某一时刻的值。
以dir为例，后面的循环不断更新dir的值，当删除操作执行时，for循环已经完成。dir中存储的值是最后一次迭代的值。
这意味着，每次对os.RemoveAll调用删除的都是相同的目录。
这个问题不仅出现在for range循环中，for循环对下标变量的引用同样存在问题。
为了解决这个问题，我们会引入一个与循环变量同名的局部变量，作为循环变量的副本。
*/
func rm_dirs() {
	tmpDirs := []string{"/tmp/a", "/tmp/b", "/tmp/c"}
	var rmdirs []func()
	for _, dir := range tmpDirs {
		dir := dir // 创建局部变量，以下面避免匿名函数中引用循环的共享变量dir而出错
		os.MkdirAll(dir, 0755)
		rmdirs = append(rmdirs, func() { os.RemoveAll(dir) })
	}

	// do some work...
	for _, rmdir := range rmdirs {
		rmdir() // clean up
	}
}

func main() {
	f := squares()
	fmt.Println(f()) // 1
	fmt.Println(f()) // 4
	fmt.Println(f()) // 9
	fmt.Println(f()) // 16

	prerequisite()
}
