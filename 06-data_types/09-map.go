package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
map

hash表是一个无序的key/value对的集合。其中所有的key都是不同的，然后通过给定的key可以在常数时间复杂度内检索、更新或删除对应的value。
在Go语言中，一个map就是一个哈希表的引用，map类型可以写作map[K]V，其中K和V分别对应key和value。
map中所有key都具有相同的类型，所有value也都有着相同的类型，但是key和value可以是不同的类型。
其中K对应的key必须是支持==比较运算符的类型，所以map可以通过测试key是否相等来判断是否已经存在。
虽然浮点数也支持比较，但是把浮点类型作为K是一个很糟糕的主意，最坏的情况会出现NaN和任何数都不相等。
相对于V对应的value的数据类型没有任何的限制。
*/

/*
创建map的方法：
1. 内置的make函数可以创建一个map
2. 也可以用map字面值的语法创建一个map

map类型的零值是nil，也就是没有引用任何hash表。
创建nil map的方法：
1. 声明一个map变量，但是不初始化，如 var a map[int]int
2. 用nil进行类型转换
map上的大部分操作，包括查找、删除、len和range遍历都可以安全的在nil值的map上工作，它们的行为和空map类型。
但是向一个nil值的map存入元素将导致一个panic异常。
*/
func map_create() {
	a := make(map[string]int)
	fmt.Printf("a: %v\tlen: %d\tnil: %t\n", a, len(a), a == nil)

	b := map[string]int{"hello": 1, "world": 2}
	fmt.Printf("b: %v\tlen: %d\n", b, len(b))

	e := map[string]int{}
	fmt.Printf("e: %v\tlen: %d\tnil: %t\n", e, len(e), e == nil)
	e["hello"] = 1
	e["world"] = 2
	fmt.Printf("e: %v\tlen: %d\tnil: %t\n", e, len(e), e == nil)

	// nil map
	var d map[string]int
	fmt.Printf("d: %v\tlen: %d\tnil: %t\n", d, len(d), d == nil)
	// d["foo"] = 3 // runtime error

	/* 可以将nil转为对应的map类型，但是赋值时会出错，报nil map访问 */
	c := map[string]int(nil)
	fmt.Printf("c: %v\tlen: %d\tnil: %t\n", c, len(c), c == nil)
	// c["foo"] = 3 // runtime error
	delete(d, "hello")
	for k, v := range c {
		fmt.Println(k, v)
	}
}

/* map 元素添加、更新和删除 */
func map_operate() {
	a := make(map[string]int)

	// map 的赋值/更新方式和数组赋值方式类似
	a["hello"] = 0
	a["world"] = 1

	// 通过访问下表来取得对应的key值，如果key不存在，则返回V对应的value类型的零值
	fmt.Println(a["foo"]) //0

	/* 为了区分是value本身是零值还是key对应的值不存在，可以通过下面的方式判断 */
	_, ok := a["foo"]
	if !ok {
		fmt.Println("the key foo is not exist.")
	}

	/* 可以使用内置的delete函数来删除元素 */
	delete(a, "hello")
	fmt.Println(a)

	/* 删除一个不存在的key的行为是安全的，不会引发异常 */
	delete(a, "foo")

	/* x += y 和 x++ 等简短赋值语句也可以用在map上 */
	a["world"]++
	fmt.Println(a)

	/* 对不存在的key值进行增减也不会引发异常 */
	a["hello"] -= 1 // 相对于 a["hello"] = a["hello"] - 1 ， 右边的表达式不存在返回类型零值0，然后再减1
	fmt.Println(a)

	/* 但是map中的元素并不是变量，因此不能对元素进行取地址操作 */
	// &a["hello"] // compile error
	/* 禁止对map元素进行取址的原因是map可能随着元素个数的增长而重新分配更大的内存空间，从而可能导致之前的地址无效 */
}

/*
map 遍历

map的遍历的迭代顺序是不确定的，并且不同的哈希函数实现可能导致不同的遍历顺序。
在实践中，遍历顺序是随机的，每次遍历的顺序都不同。这是故意的，每次都使用随机的遍历顺序可以强制要求程序不依赖具体的哈希函数实现。
如果要按顺序遍历key/value对，必须显示的对key进行排序，然后按照排序的key取出value。
可以使用sort包的strings函数对字符串slice进行排序。
*/
func map_range() {
	a := map[string]int{"hello": 1, "world": 2, "foo": 3, "bar": 4}
	for k, v := range a {
		fmt.Println(k, v)
	}
	fmt.Println(strings.Repeat("-", 16))

	// var keys []string
	keys := make([]string, 0, len(a)) // 因为知道需要的长度，所以创建一个合适大小的slice会更有效率
	for k := range a {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k, a[k])
	}
}

/* map之间不能使用==进行相等比较 ，唯一例外的是和nil比较*/
func map_compare() {
	a := map[string]int{}
	var b map[string]int
	c := map[string]int(nil)

	// fmt.Println(a == b) // compile error
	fmt.Println(a == nil)
	fmt.Println(b == nil)
	fmt.Println(c == nil)
}

/* map的V可以是聚合类型，如一个map或者slice */
func map_composite() {
	graph := make(map[string]map[string]bool)
	fmt.Printf("graph: %p\t&graph: %p\n", graph, &graph) // 0xc00006a510 0xc000012030
	add_edge(graph, "a", "b")
	fmt.Printf("a --> b: %t\n", has_edge(graph, "a", "b"))
	fmt.Printf("a --> c: %t\n", has_edge(graph, "a", "c"))
	fmt.Printf("b --> c: %t\n", has_edge(graph, "b", "c"))
}

/* map作为函数参数传递时传递的是指针，即参数的局部变量保存的是map的地址 */
func add_edge(graph map[string]map[string]bool, from, to string) {
	fmt.Printf("add_edge, graph: %p\t&graph: %p\n", graph, &graph) // 0xc00006a510 0xc000012038
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}

	/* 取出的值是map类型，可以直接修改，不会担心局部变量问题 */
	edges[to] = true
}
func has_edge(graph map[string]map[string]bool, from, to string) bool {
	return graph[from][to]
}

func main() {
	map_create()
	map_operate()
	map_range()
	map_compare()
	map_composite()
}
